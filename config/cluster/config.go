package cluster

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures the rancher2_cluster group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("rancher2_cluster", func(r *tjconfig.Resource) {
		if block, ok := r.TerraformResource.Schema["gke_config_v2"]; ok {
			if attr, ok := block.Elem.(*schema.Resource).Schema["google_credential_secret"]; ok {
				attr.Description = "The GKE Cloud Credential ID to use"
				attr.Sensitive = false
			}
		}
	})
}
