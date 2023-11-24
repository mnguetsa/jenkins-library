// Code generated by piper's step-generator. DO NOT EDIT.

package metadata

import (
	"os"
	"github.com/SAP/jenkins-library/pkg/config"
)

// retrieve step metadata
func CloudFoundryCreateServiceKeyMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "cloudFoundryCreateServiceKey",
			Aliases: []config.Alias{},
			Description: "cloudFoundryCreateServiceKey",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "cfCredentialsId",Description: "Jenkins 'Username with password' credentials ID containing user and password to authenticate to the Cloud Foundry API.",Type: "jenkins",
					}, 
				},
				Parameters: []config.StepParameters{
					{
						Name:      "cfApiEndpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{{Name: "cloudFoundry/apiEndpoint"},},
						Default:   os.Getenv("PIPER_cfApiEndpoint"),
					},
					{
						Name:      "username",
						ResourceRef: []config.ResourceReference{
							{
								Name:"cfCredentialsId",
								Param: "username",
								Type: "secret",
							},
                        
							{
								Name:"cloudfoundryVaultSecretName",
								Type: "vaultSecret",
								Default: "cloudfoundry-$(org)-$(space)",
							},
                        },
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_username"),
					},
					{
						Name:      "password",
						ResourceRef: []config.ResourceReference{
							{
								Name:"cfCredentialsId",
								Param: "password",
								Type: "secret",
							},
                        
							{
								Name:"cloudfoundryVaultSecretName",
								Type: "vaultSecret",
								Default: "cloudfoundry-$(org)-$(space)",
							},
                        },
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_password"),
					},
					{
						Name:      "cfOrg",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{{Name: "cloudFoundry/org"},},
						Default:   os.Getenv("PIPER_cfOrg"),
					},
					{
						Name:      "cfSpace",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{{Name: "cloudFoundry/space"},},
						Default:   os.Getenv("PIPER_cfSpace"),
					},
					{
						Name:      "cfServiceInstance",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{{Name: "cloudFoundry/serviceInstance"},},
						Default:   os.Getenv("PIPER_cfServiceInstance"),
					},
					{
						Name:      "cfServiceKeyName",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS","GENERAL",},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{{Name: "cloudFoundry/serviceKey"},{Name: "cloudFoundry/serviceKeyName"},{Name: "cfServiceKey"},},
						Default:   os.Getenv("PIPER_cfServiceKeyName"),
					},
					{
						Name:      "cfServiceKeyConfig",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "cloudFoundry/serviceKeyConfig"},},
						Default:   os.Getenv("PIPER_cfServiceKeyConfig"),
					},
					{
						Name:      "cfAsync",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"GENERAL","PARAMETERS","STAGES","STEPS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   true,
					},
				},
			},
			Containers: []config.Container{
				{Name: "cf",Image: "ppiper/cf-cli:latest",
				}, 
			},
			
		},
	}
	return theMetaData
}
