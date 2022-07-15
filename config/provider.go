/*
Copyright 2021 The Crossplane Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-rancher/config/catalog"
	"github.com/crossplane-contrib/provider-jet-rancher/config/cluster"
)

const (
	terraformResourcePrefix = "rancher2"
	modulePath              = "github.com/crossplane-contrib/provider-jet-rancher"
)

//go:embed schema.json
var providerSchema string

var skipList = []string{
	"rancher2_bootstrap",                    // requires support for the bootstrap flag in the provider config
	"rancher2_pod_security_policy_template", // deprecated as of Kubernetes v1.21, and will be removed in v1.25
}

var includeList = []string{
	"rancher2_app",
	"rancher2_auth_config_",
	"rancher2_catalog",
	"rancher2_certificate",
	"rancher2_cloud_credential",
	"rancher2_cluster",
	"rancher2_config_map_v2",
	"rancher2_etcd_backup",
	"rancher2_feature",
	"rancher2_global_",
	"rancher2_machine_config_v2",
	"rancher2_multi_cluster_app",
	"rancher2_namespace",
	"rancher2_node_",
	"rancher2_notifier",
	"rancher2_project",
	"rancher2_registry",
	"rancher2_role_template",
	"rancher2_secret",
	"rancher2_setting",
	"rancher2_storage_class_v2",
	"rancher2_token",
	"rancher2_user",
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), terraformResourcePrefix, modulePath,
		tjconfig.WithShortName("rancherjet"),
		tjconfig.WithRootGroup("rancher.jet.crossplane.io"),
		tjconfig.WithIncludeList(includeList),
		tjconfig.WithSkipList(skipList),
		tjconfig.WithDefaultResourceFn(DefaultResource(
			GroupKindOverrides(),
			KindOverrides(),
			IdentifierAssignedByRancher(),
			KnownReferencers(),
		)),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		catalog.Configure,
		cluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc

}

// DefaultResource returns a DefaultResoruceFn that makes sure the original
// DefaultResource call is made with given options here.
//nolint
func DefaultResource(opts ...tjconfig.ResourceOption) tjconfig.DefaultResourceFn {
	return func(name string, terraformResource *schema.Resource, orgOpts ...tjconfig.ResourceOption) *tjconfig.Resource {
		return tjconfig.DefaultResource(name, terraformResource, append(orgOpts, opts...)...)
	}
}
