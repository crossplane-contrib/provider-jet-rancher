package namespace

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the rancher2_namespace group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("rancher2_namespace", func(r *tjconfig.Resource) {
		r.Kind = "RancherNamespace"

		r.References["project_id"] = tjconfig.Reference{
			Type: "Project",
		}
	})
}
