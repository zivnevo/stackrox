package component

import (
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/sensor/common/store/resolver"
)

// CompatibilityDetectionMessage should be used by old handlers
// it's here for compatibility reasons.
type CompatibilityDetectionMessage struct {
	// Object references the object that needs to be processed by the detector
	Object *storage.Deployment
	// Action of an event (CREATE_RESOURCE, REMOVE_RESOURCE, UPDATE_RESOURCE, and SYNC_RESOURCE)
	Action central.ResourceAction
}

// ResourceEvent message used by the event pipeline's components
type ResourceEvent struct {
	// ForwardMessages messages generated by the handlers that need to be forwarded to central
	ForwardMessages []*central.SensorEvent
	// CompatibilityDetectionDeployment should be used by old handlers
	// and it's here for compatibility reasons.
	// This property should be removed in the future and only the
	// deployment references should be sent
	CompatibilityDetectionDeployment []CompatibilityDetectionMessage
	// CompatibilityReprocessDeployments is also used for compatibility reasons with Network Policy handlers
	// in the future this will not be needed as the dependencies are taken care by the resolvers
	CompatibilityReprocessDeployments []string

	// DeploymentReference returns an implementation of a struct that can return a list of deployment ids
	// that require processing
	DeploymentReference resolver.DeploymentReference

	// DeploymentTiming has the timing object that needs to be added to any deployments resolved by this event.
	DeploymentTiming *central.Timing

	// ParentResourceAction is the resource action that will be sent to central on the deployment event.
	// If the ResourceEvent originated on a deployment event, this should be set to whatever action triggered
	// the event. For related resources updates, like RBACs and services, this should always be set to
	// UPDATE.
	ParentResourceAction central.ResourceAction

	// ForceDetection is a flag that will force a detection even if the deployment has no changes.
	// This is needed to trigger detection of deployments associated with a NetworkPolicy, since they are not part of the deployment spec
	// and therefore will not be triggered since the deduper won't recognize that a deployment was changed.
	ForceDetection bool
}

// NewResourceEvent wraps the SensorEvents, CompatibilityDetectionMessages, and the CompatibilityReprocessDeployments into a ResourceEvent message
func NewResourceEvent(sensorMessages []*central.SensorEvent, detectionDeployment []CompatibilityDetectionMessage, reprocessDeploymentsIds []string) *ResourceEvent {
	return &ResourceEvent{
		ForwardMessages:                   sensorMessages,
		CompatibilityDetectionDeployment:  detectionDeployment,
		CompatibilityReprocessDeployments: reprocessDeploymentsIds,
	}
}

// NewDeploymentRefEvent returns a resource event given a deployment reference and a resource action.
func NewDeploymentRefEvent(ref resolver.DeploymentReference, action central.ResourceAction, forceDetection bool) *ResourceEvent {
	return &ResourceEvent{
		DeploymentReference:  ref,
		ParentResourceAction: action,
		ForceDetection:       forceDetection,
	}
}

// MergeResourceEvents merges two ResourceEvents
func MergeResourceEvents(dest, src *ResourceEvent) *ResourceEvent {
	if dest == nil {
		dest = &ResourceEvent{}
	}

	if src != nil {
		dest.CompatibilityReprocessDeployments = append(dest.CompatibilityReprocessDeployments, src.CompatibilityReprocessDeployments...)
		dest.ForwardMessages = append(dest.ForwardMessages, src.ForwardMessages...)
		dest.CompatibilityDetectionDeployment = append(dest.CompatibilityDetectionDeployment, src.CompatibilityDetectionDeployment...)
		if src.ParentResourceAction != central.ResourceAction_UNSET_ACTION_RESOURCE {
			dest.ParentResourceAction = src.ParentResourceAction
		}
		if src.DeploymentReference != nil {
			dest.DeploymentReference = src.DeploymentReference
		}
		dest.ForceDetection = src.ForceDetection || dest.ForceDetection
	}
	return dest
}
