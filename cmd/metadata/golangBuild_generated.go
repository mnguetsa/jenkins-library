// Code generated by piper's step-generator. DO NOT EDIT.

package metadata

import (
	"os"
	"github.com/SAP/jenkins-library/pkg/config"
)

// retrieve step metadata
func GolangBuildMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "golangBuild",
			Aliases: []config.Alias{},
			Description: "This step will execute a golang build.",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "golangPrivateModulesGitTokenCredentialsId",Description: "Jenkins 'Username with password' credentials ID containing username/password for http access to your git repos where your go private modules are stored.",Type: "jenkins",
					}, 
				},
				Parameters: []config.StepParameters{
					{
						Name:      "buildFlags",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   []string{},
					},
					{
						Name:      "buildSettingsInfo",
						ResourceRef: []config.ResourceReference{
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/buildSettingsInfo",
							},
                        },
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_buildSettingsInfo"),
					},
					{
						Name:      "cgoEnabled",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
					{
						Name:      "coverageFormat",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   `html`,
					},
					{
						Name:      "createBOM",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"GENERAL","STEPS","STAGES","PARAMETERS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
					{
						Name:      "customTlsCertificateLinks",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"GENERAL","PARAMETERS","STAGES","STEPS",},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   []string{},
					},
					{
						Name:      "excludeGeneratedFromCoverage",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   true,
					},
					{
						Name:      "ldflagsTemplate",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_ldflagsTemplate"),
					},
					{
						Name:      "output",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_output"),
					},
					{
						Name:      "packages",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   []string{},
					},
					{
						Name:      "publish",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
					{
						Name:      "targetRepositoryPassword",
						ResourceRef: []config.ResourceReference{
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/rawRepositoryPassword",
							},
                        
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/repositoryPassword",
							},
                        },
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_targetRepositoryPassword"),
					},
					{
						Name:      "targetRepositoryUser",
						ResourceRef: []config.ResourceReference{
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/rawRepositoryUsername",
							},
                        
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/repositoryUsername",
							},
                        },
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_targetRepositoryUser"),
					},
					{
						Name:      "targetRepositoryURL",
						ResourceRef: []config.ResourceReference{
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/rawRepositoryURL",
							},
                        
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/repositoryUrl",
							},
                        },
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_targetRepositoryURL"),
					},
					{
						Name:      "reportCoverage",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   true,
					},
					{
						Name:      "runLint",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
					{
						Name:      "runTests",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   true,
					},
					{
						Name:      "runIntegrationTests",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
					{
						Name:      "targetArchitectures",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"GENERAL","STEPS","STAGES","PARAMETERS",},
						Type:      "[]string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   []string{`linux,amd64`},
					},
					{
						Name:      "testOptions",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   []string{},
					},
					{
						Name:      "testResultFormat",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   `junit`,
					},
					{
						Name:      "privateModules",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"GENERAL","STEPS","STAGES","PARAMETERS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_privateModules"),
					},
					{
						Name:      "privateModulesGitToken",
						ResourceRef: []config.ResourceReference{
							{
								Name:"golangPrivateModulesGitTokenCredentialsId",
								Param: "password",
								Type: "secret",
							},
                        
							{
								Name:"golangPrivateModulesGitTokenVaultSecret",
								Type: "vaultSecret",
								Default: "golang",
							},
                        },
						Scope:     []string{"GENERAL","PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_privateModulesGitToken"),
					},
					{
						Name:      "artifactVersion",
						ResourceRef: []config.ResourceReference{
							{
								Name:"commonPipelineEnvironment",
								Param: "artifactVersion",
							},
                        },
						Scope:     []string{"GENERAL","PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_artifactVersion"),
					},
					{
						Name:      "golangciLintUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   `https://github.com/golangci/golangci-lint/releases/download/v1.51.2/golangci-lint-1.51.2-linux-amd64.tar.gz`,
					},
				},
			},
			Containers: []config.Container{
				{Name: "golang",Image: "golang:1",Options: []config.Option{ {Name: "-u", Value: "0"}, },
				}, 
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "commonPipelineEnvironment",
						Type: "piperEnvironment",
						Parameters: []map[string]interface{}{
							{"name": "custom/buildSettingsInfo",},
							{"name": "custom/artifacts","type": "piperenv.Artifacts",},
						},
					},
					{
						Name: "reports",
						Type: "reports",
						Parameters: []map[string]interface{}{
							{"filePattern": "**/bom-golang.xml","type": "sbom",},
							{"filePattern": "**/TEST-*.xml","type": "junit",},
							{"filePattern": "**/cobertura-coverage.xml","type": "cobertura-coverage",},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
