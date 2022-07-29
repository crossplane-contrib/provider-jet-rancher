package cluster

import (
	"github.com/crossplane-contrib/provider-jet-rancher/config/common"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures the rancher2_cluster group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("rancher2_cluster", func(r *tjconfig.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = tjconfig.IdentifierFromProvider
		if block, ok := r.TerraformResource.Schema["gke_config_v2"]; ok {
			if attr, ok := block.Elem.(*schema.Resource).Schema["google_credential_secret"]; ok {
				// Terrajet (v0.4) has no option to rename the attribute name
				// attr.Name = "cloudCredentialId"
				attr.Description = "The GKE Cloud Credential ID to use"
				attr.Sensitive = false
			}
		}

		r.References["aks_config_v2.cloud_credential_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-rancher/apis/rancher/v1alpha1.CloudCredential",
		}
		r.References["eks_config_v2.cloud_credential_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-rancher/apis/rancher/v1alpha1.CloudCredential",
		}
		r.References["gke_config_v2.google_credential_secret"] = tjconfig.Reference{
			Type:              "github.com/crossplane-contrib/provider-jet-rancher/apis/rancher/v1alpha1.CloudCredential",
			RefFieldName:      "CloudCredentialIdRef",
			SelectorFieldName: "CloudCredentialIdSelector",
		}
	})
}
