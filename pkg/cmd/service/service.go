package service

import (
	"fmt"
	"os/exec"
	"strings"

	"git.ecd.axway.int/apigov/apic_agents_sdk/pkg/cmd/service/daemon"

	"git.ecd.axway.int/apigov/apic_agents_sdk/pkg/util/log"
)

var (
	// Name -
	Name string
	// Description -
	Description string

	dependencies = []string{"network"}

	globalAgentService AgentService

	execCommand = exec.Command
)

//AgentService -
type AgentService struct {
	service     daemon.Daemon
	Name        string
	Description string
	Path        string
	PathArg     string
	User        string
	Group       string
}

func init() {
	service, err := daemon.New(Name, Description, dependencies...)
	if err != nil {
		log.Errorf("error hit creating the service definition: %s", err.Error())
	}

	globalAgentService = AgentService{
		service:     service,
		Name:        Name,
		Description: Description,
	}
}

// HandleServiceFlag - handles the action needed based ont eh service flag value
func (a *AgentService) HandleServiceFlag(command string) error {
	var err error
	var status string
	// complete teh appropriate action for the service
	switch strings.ToLower(command) {
	case "install":
		log.Debug("installing the agent service")
		log.Infof("service will look for config file at %s", a.Path)
		_, err = a.service.Install(a.PathArg, a.Path)
	case "remove":
		log.Debug("removing the agent service")
		_, err = a.service.Remove()
	case "start":
		log.Debug("starting the agent service")
		_, err = a.service.Start()
	case "stop":
		log.Debug("stoping the agent service")
		_, err = a.service.Stop()
	case "status":
		log.Debug("getting the agent service status")
		status, err = a.service.Status()
	case "enable":
		log.Debug("setting the agent to start on reboot")
		_, err = a.service.Enable()
	default:
		err = fmt.Errorf("unknown value of '%s' given", command)
	}

	// error hit
	if err != nil {
		log.Errorf("service %s command failed: %s", strings.ToLower(command), err.Error())
	} else {
		log.Debugf("service %s command succeeded", strings.ToLower(command))
		if status != "" {
			log.Info(status)
		}
	}
	return err
}
