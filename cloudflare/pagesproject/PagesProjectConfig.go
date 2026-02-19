// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pagesproject

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PagesProjectConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#account_id PagesProject#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#name PagesProject#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Production branch of the project. Used to identify production deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#production_branch PagesProject#production_branch}
	ProductionBranch *string `field:"required" json:"productionBranch" yaml:"productionBranch"`
	// Configs for the project build process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#build_config PagesProject#build_config}
	BuildConfig *PagesProjectBuildConfig `field:"optional" json:"buildConfig" yaml:"buildConfig"`
	// Configs for deployments in a project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#deployment_configs PagesProject#deployment_configs}
	DeploymentConfigs *PagesProjectDeploymentConfigs `field:"optional" json:"deploymentConfigs" yaml:"deploymentConfigs"`
	// Configs for the project source control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#source PagesProject#source}
	Source *PagesProjectSource `field:"optional" json:"source" yaml:"source"`
}

