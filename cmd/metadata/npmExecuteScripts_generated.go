// Code generated by piper's step-generator. DO NOT EDIT.

package metadata

import (
	"os"
	"github.com/SAP/jenkins-library/pkg/config"
)

// retrieve step metadata
func NpmExecuteScriptsMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "npmExecuteScripts",
			Aliases: []config.Alias{{Name: "executeNpm", Deprecated: false},},
			Description: "Execute npm run scripts on all npm packages in a project",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Resources: []config.StepResources{
					{Name: "source",Type: "stash",
					},
				},
				Parameters: []config.StepParameters{
					{
						Name:      "install",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   true,
					},
					{
						Name:      "runScripts",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   []string{},
					},
					{
						Name:      "defaultNpmRegistry",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","GENERAL","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "npm/defaultNpmRegistry"},},
						Default:   os.Getenv("PIPER_defaultNpmRegistry"),
					},
					{
						Name:      "virtualFrameBuffer",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
					{
						Name:      "scriptOptions",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   []string{},
					},
					{
						Name:      "buildDescriptorExcludeList",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   []string{`deployment/**`},
					},
					{
						Name:      "buildDescriptorList",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"PARAMETERS","STAGES","STEPS",},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   []string{},
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
						Name:      "publish",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
					{
						Name:      "repositoryUrl",
						ResourceRef: []config.ResourceReference{
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/npmRepositoryURL",
							},
                        
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/repositoryUrl",
							},
                        },
						Scope:     []string{"GENERAL","PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_repositoryUrl"),
					},
					{
						Name:      "repositoryPassword",
						ResourceRef: []config.ResourceReference{
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/npmRepositoryPassword",
							},
                        
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/repositoryPassword",
							},
                        },
						Scope:     []string{"GENERAL","PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_repositoryPassword"),
					},
					{
						Name:      "repositoryUsername",
						ResourceRef: []config.ResourceReference{
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/npmRepositoryUsername",
							},
                        
							{
								Name:"commonPipelineEnvironment",
								Param: "custom/repositoryUsername",
							},
                        },
						Scope:     []string{"GENERAL","PARAMETERS","STAGES","STEPS",},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_repositoryUsername"),
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
						Name:      "packBeforePublish",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
					{
						Name:      "production",
						ResourceRef: []config.ResourceReference{},
						Scope:     []string{"STEPS","STAGES","PARAMETERS",},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
				},
			},
			Containers: []config.Container{
				{Name: "node",Image: "node:lts-buster",
				}, 
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "commonPipelineEnvironment",
						Type: "piperEnvironment",
						Parameters: []map[string]interface{}{
							{"name": "custom/buildSettingsInfo",},
						},
					},
					{
						Name: "reports",
						Type: "reports",
						Parameters: []map[string]interface{}{
							{"filePattern": "**/bom-npm.xml","type": "sbom",},
							{"filePattern": "**/TEST-*.xml","type": "junit",},
							{"filePattern": "**/cobertura-coverage.xml","type": "cobertura-coverage",},
							{"filePattern": "**/e2e/*.json","type": "cucumber",},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
