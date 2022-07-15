package catalog

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the rancher2_catalog group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("rancher2_catalog", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.NameAsIdentifier

		// Currently not supported by Terrajet (v0.4)
		// r.TerraformResource.Schema["scope"].Default = "global"
	})
}
