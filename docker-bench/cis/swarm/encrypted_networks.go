package swarm

import (
	"bitbucket.org/stack-rox/apollo/docker-bench/utils"
	"bitbucket.org/stack-rox/apollo/pkg/api/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

type encryptedNetworks struct{}

func (c *encryptedNetworks) Definition() utils.Definition {
	return utils.Definition{
		CheckDefinition: v1.CheckDefinition{
			Name:        "CIS 7.4",
			Description: "Ensure data exchanged between containers are encrypted on different nodes on the overlay network",
		},
		Dependencies: []utils.Dependency{utils.InitInfo},
	}
}

func (c *encryptedNetworks) Run() (result v1.CheckResult) {
	if !utils.DockerInfo.Swarm.ControlAvailable {
		utils.NotApplicable(&result)
		utils.AddNotes(&result, "Checking encrypted networks applies only to Swarm managers and this node is not a Swarm Manager")
		return
	}
	ctx, cancel := docker.TimeoutContext()
	defer cancel()
	filters := filters.NewArgs()
	filters.Add("driver", "overlay")
	networkList, err := utils.DockerClient.NetworkList(ctx, types.NetworkListOptions{
		Filters: filters,
	})
	if err != nil {
		utils.Warn(&result)
		utils.AddNotef(&result, "Couldn't list Docker networks: %+v", err)
		return
	}
	utils.Pass(&result)
	for _, network := range networkList {
		if val := network.Options["encrypted"]; val != "true" {
			utils.Warn(&result)
			utils.AddNotef(&result, "Overlay Network '%v' (%v) is not encrypted", network.Name, network.ID)
		}
	}
	return
}

// NewEncryptedNetworks implements CIS-7.4
func NewEncryptedNetworks() utils.Check {
	return &encryptedNetworks{}
}
