// Code generated by piper's step-generator. DO NOT EDIT.

package metadata

import (
	"os"
	"github.com/SAP/jenkins-library/pkg/config"
)

// retrieve step metadata
func AbapEnvironmentAssemblePackagesMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "abapEnvironmentAssemblePackages",
			Aliases: []config.Alias{},
			Description: "Assembly of installation, support package or patch in SAP BTP ABAP Environment system",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "abapCredentialsId",Description: "Jenkins credentials ID containing user and password to authenticate to the Cloud Platform ABAP Environment system or the Cloud Foundry API",Type: "jenkins",Aliases: []config.Alias{ {Name: "cfCredentialsId", Deprecated: false},  {Name: "credentialsId", Deprecated: false}, },
					}, 
				},
				Parameters: []config.StepParameters{
					{
						Name:      "cfApiEndpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "cloudFoundry/apiEndpoint"},},
						Default:   os.Getenv("PIPER_cfApiEndpoint"),
					},
					{
						Name:      "cfOrg",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "cloudFoundry/org"},},
						Default:   os.Getenv("PIPER_cfOrg"),
					},
					{
						Name:      "cfSpace",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "cloudFoundry/space"},},
						Default:   os.Getenv("PIPER_cfSpace"),
					},
					{
						Name:      "cfServiceInstance",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "cloudFoundry/serviceInstance"},},
						Default:   os.Getenv("PIPER_cfServiceInstance"),
					},
					{
						Name:      "cfServiceKeyName",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "cloudFoundry/serviceKey"},{Name: "cloudFoundry/serviceKeyName"},{Name: "cfServiceKey"},},
						Default:   os.Getenv("PIPER_cfServiceKeyName"),
					},
					{
						Name:      "host",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_host"),
					},
					{
						Name:      "username",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_username"),
					},
					{
						Name:      "password",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_password"),
					},
					{
						Name:      "addonDescriptor",
						ResourceRef: []config.ResourceReference{
							{
								Name:"commonPipelineEnvironment",
								Param: "abap/addonDescriptor",
							},
                        },
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_addonDescriptor"),
					},
					{
						Name:      "maxRuntimeInMinutes",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "int",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   360,
					},
					{
						Name:      "pollIntervalsInMilliseconds",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "int",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   60000,
					},
					{
						Name:      "certificateNames",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   []string{},
					},
				},
			},
			Containers: []config.Container{
				{Name: "cf",Image: "ppiper/cf-cli:v12",
				}, 
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "commonPipelineEnvironment",
						Type: "piperEnvironment",
						Parameters: []map[string]interface{}{
							{"name": "abap/addonDescriptor",},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
