package catalog

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("rancher_catalog", func(r *config.Resource) {

		r.ShortGroup = "catalog"
	})
}
