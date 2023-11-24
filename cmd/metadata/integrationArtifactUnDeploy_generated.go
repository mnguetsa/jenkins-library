// Code generated by piper's step-generator. DO NOT EDIT.

package metadata

import (
	"os"
	"github.com/SAP/jenkins-library/pkg/config"
)

// retrieve step metadata
func IntegrationArtifactUnDeployMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "integrationArtifactUnDeploy",
			Aliases: []config.Alias{},
			Description: "Undeploy a integration flow",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "cpiApiServiceKeyCredentialsId",Description: "Jenkins secret text credential ID containing the service key to the Process Integration Runtime service instance of plan 'api'",Type: "jenkins",
					}, 
				},
				Parameters: []config.StepParameters{
					{
						Name:      "apiServiceKey",
						ResourceRef: []config.ResourceReference{
							{
								Name:"cpiApiServiceKeyCredentialsId",
								Param: "apiServiceKey",
								Type: "secret",
							},
                        },
						Scope:     []string{"PARAMETERS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_apiServiceKey"),
					},
					{
						Name:      "integrationFlowId",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","GENERAL","STAGES","STEPS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_integrationFlowId"),
					},
				},
			},
			
		},
	}
	return theMetaData
}
