// Code generated by piper's step-generator. DO NOT EDIT.

package metadata

import (
	"os"
	"github.com/SAP/jenkins-library/pkg/config"
)

// retrieve step metadata
func AnsSendEventMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "ansSendEvent",
			Aliases: []config.Alias{},
			Description: "Send Event to the SAP Alert Notification Service",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "ansServiceKeyCredentialsId",Description: "Jenkins secret text credential ID containing the service key to access the SAP Alert Notification Service",Type: "jenkins",
					}, 
				},
				Parameters: []config.StepParameters{
					{
						Name:      "ansServiceKey",
						ResourceRef: []config.ResourceReference{
							{
								Name:"ansServiceKeyCredentialsId",
								Param: "ansServiceKey",
								Type: "secret",
							},
                        },
						Scope:     []string{"PARAMETERS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_ansServiceKey"),
					},
					{
						Name:      "eventType",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   `Piper`,
					},
					{
						Name:      "severity",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   `INFO`,
					},
					{
						Name:      "category",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   `NOTIFICATION`,
					},
					{
						Name:      "subject",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   `ansSendEvent`,
					},
					{
						Name:      "body",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   `Call from Piper step ansSendEvent`,
					},
					{
						Name:      "priority",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "int",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   0,
					},
					{
						Name:      "tags",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "map[string]interface{}",
						Mandatory: false,
						Aliases:   []config.Alias{},
						
					},
					{
						Name:      "resourceName",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   `Pipeline`,
					},
					{
						Name:      "resourceType",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   `Pipeline`,
					},
					{
						Name:      "resourceInstance",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_resourceInstance"),
					},
					{
						Name:      "resourceTags",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "map[string]interface{}",
						Mandatory: false,
						Aliases:   []config.Alias{},
						
					},
				},
			},
			
		},
	}
	return theMetaData
}
