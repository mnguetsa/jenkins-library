// Code generated by piper's step-generator. DO NOT EDIT.

package metadata

import (
	"os"
	"github.com/SAP/jenkins-library/pkg/config"
)

// retrieve step metadata
func ApiKeyValueMapUploadMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "apiKeyValueMapUpload",
			Aliases: []config.Alias{},
			Description: "this steps creates an API key value map artifact in the API Portal",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "apimApiServiceKeyCredentialsId",Description: "Jenkins secret text credential ID containing the service key to the API Management Runtime service instance of plan 'api'",Type: "jenkins",
					}, 
				},
				Parameters: []config.StepParameters{
					{
						Name:      "apiServiceKey",
						ResourceRef: []config.ResourceReference{
							{
								Name:"apimApiServiceKeyCredentialsId",
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
						Name:      "key",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_key"),
					},
					{
						Name:      "value",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_value"),
					},
					{
						Name:      "keyValueMapName",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_keyValueMapName"),
					},
				},
			},
			
		},
	}
	return theMetaData
}
