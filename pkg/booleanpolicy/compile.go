package booleanpolicy

import (
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/booleanpolicy/query"
	"github.com/stackrox/rox/pkg/set"
)

func sectionToQuery(section *storage.PolicySection, stage storage.LifecycleStage) (*query.Query, error) {
	if len(section.GetPolicyGroups()) == 0 {
		return nil, errors.New("no groups")
	}
	fieldQueries, err := sectionToFieldQueries(section, nil)
	if err != nil {
		return nil, err
	}
	contextQueries := constructRemainingContextQueries(stage, section, fieldQueries)
	fieldQueries = append(fieldQueries, contextQueries...)

	return &query.Query{FieldQueries: fieldQueries}, nil
}

func sectionToFieldQueries(section *storage.PolicySection, allowedGroups *set.FrozenStringSet) ([]*query.FieldQuery, error) {
	fieldQueries := make([]*query.FieldQuery, 0, len(section.GetPolicyGroups()))
	for _, group := range section.GetPolicyGroups() {
		if allowedGroups != nil && !allowedGroups.Contains(group.GetFieldName()) {
			continue
		}
		fqs, err := policyGroupToFieldQueries(group)
		if err != nil {
			return nil, errors.Wrapf(err, "constructing query for group %s", group.GetFieldName())
		}
		fieldQueries = append(fieldQueries, fqs...)
	}
	return fieldQueries, nil
}

func policyGroupToFieldQueries(group *storage.PolicyGroup) ([]*query.FieldQuery, error) {
	if len(group.GetValues()) == 0 {
		return nil, errors.New("no values")
	}

	f, ok := fieldsToQB[group.GetFieldName()]
	if !ok {
		return nil, errors.Errorf("no QB known for group %q", group.GetFieldName())
	}
	metadata := f(nil)
	if metadata == nil || metadata.qb == nil {
		return nil, errors.Errorf("no QB known for group %q", group.GetFieldName())
	}

	if metadata.negationForbidden && group.GetNegate() {
		return nil, errors.Errorf("invalid group: negation not allowed for field %s", group.GetFieldName())
	}
	if metadata.operatorsForbidden && len(group.GetValues()) != 1 {
		return nil, errors.Errorf("invalid group: operators not allowed for field %s", group.GetFieldName())
	}

	fqs := metadata.qb.FieldQueriesForGroup(group)
	if len(fqs) == 0 {
		return nil, errors.New("invalid group: no queries formed")
	}

	return fqs, nil
}

func matchAllQueryForField(fieldName string) *query.FieldQuery {
	return &query.FieldQuery{
		Field:    fieldName,
		MatchAll: true,
	}
}

// Add matchAll field queries for context fields that are not already included fieldQueries
func constructRemainingContextQueries(stage storage.LifecycleStage, section *storage.PolicySection, fieldQueries []*query.FieldQuery) []*query.FieldQuery {
	fieldSet := set.NewStringSet()
	for _, query := range fieldQueries {
		fieldSet.Add(query.Field)
	}
	contextFieldSet := set.NewStringSet()
	for _, group := range section.GetPolicyGroups() {
		field := group.GetFieldName()
		if f, ok := fieldsToQB[field]; ok {
			metadata := f(nil)
			if contextFieldsToAdd, ok := metadata.contextFields[stage]; ok {
				for _, contextField := range contextFieldsToAdd.AsSlice() {
					contextFieldSet.Add(contextField)
				}
			}
		}
	}
	var contextQueries []*query.FieldQuery
	for contextField := range contextFieldSet {
		if !fieldSet.Contains(contextField) {
			contextQueries = append(contextQueries, matchAllQueryForField(contextField))
		}
	}
	return contextQueries
}
