package swarm

import (
	"bitbucket.org/stack-rox/apollo/docker-bench/utils"
	"bitbucket.org/stack-rox/apollo/pkg/api/generated/api/v1"
)

type autoLockRotation struct{}

func (c *autoLockRotation) Definition() utils.Definition {
	return utils.Definition{
		CheckDefinition: v1.CheckDefinition{
			Name:        "CIS 7.7",
			Description: "Ensure swarm manager auto-lock key is rotated periodically",
		},
		Dependencies: []utils.Dependency{utils.InitInfo},
	}
}

func (c *autoLockRotation) Run() (result v1.CheckResult) {
	if !utils.DockerInfo.Swarm.ControlAvailable {
		utils.NotApplicable(&result)
		utils.AddNotes(&result, "Autolock applies only to Swarm managers and this node is not a Swarm Manager")
		return
	}
	utils.Note(&result)
	utils.AddNotes(&result, "Rotate the auto-lock key periodically using 'docker swarm unlock-key --rotate'")
	return
}

// NewAutoLockRotationCheck implements CIS-7.7
func NewAutoLockRotationCheck() utils.Check {
	return &autoLockRotation{}
}
