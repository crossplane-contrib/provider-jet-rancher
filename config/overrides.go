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
	"regexp"
	"strings"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/types/name"
	"github.com/pkg/errors"
)

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

// GroupMap contains all overrides we'd like to make to the default group search.
var GroupMap = map[string]GroupKindCalculator{
	"rancher2_app":                         ReplaceGroupWords("app", 0),
	"rancher2_app_v2":                      ReplaceGroupWords("appv2", 0),
	"rancher2_auth_config_activedirectory": ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_adfs":            ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_azuread":         ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_freeipa":         ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_github":          ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_keycloak":        ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_okta":            ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_openldap":        ReplaceGroupWords("authconfig", 2),
	"rancher2_auth_config_ping":            ReplaceGroupWords("authconfig", 2),
	"rancher2_catalog":                     ReplaceGroupWords("catalog", 0),
	"rancher2_catalog_v2":                  ReplaceGroupWords("catalogv2", 0),
	"rancher2_certificate":                 ReplaceGroupWords("project", 0),
	"rancher2_cloud_credential":            ReplaceGroupWords("", 0),
	"rancher2_cluster":                     ReplaceGroupWords("cluster", 0),
	"rancher2_cluster_v2":                  ReplaceGroupWords("clusterv2", 0),
	"rancher2_config_map_v2":               ReplaceGroupWords("cluster", 0),
	"rancher2_etcd_backup":                 ReplaceGroupWords("cluster", 0),
	"rancher2_feature":                     ReplaceGroupWords("", 0),
	"rancher2_machine_config_v2":           ReplaceGroupWords("", 0),
	"rancher2_multi_cluster_app":           ReplaceGroupWords("appv2", 0),
	"rancher2_namespace":                   ReplaceGroupWords("project", 0),
	"rancher2_node_driver":                 ReplaceGroupWords("node", 0),
	"rancher2_node_pool":                   ReplaceGroupWords("node", 0),
	"rancher2_node_template":               ReplaceGroupWords("node", 0),
	"rancher2_notifier":                    ReplaceGroupWords("cluster", 0),
	"rancher2_project":                     ReplaceGroupWords("project", 0),
	"rancher2_registry":                    ReplaceGroupWords("project", 0),
	"rancher2_role_template":               ReplaceGroupWords("", 0),
	"rancher2_secret":                      ReplaceGroupWords("secret", 0),
	"rancher2_secret_v2":                   ReplaceGroupWords("secretv2", 0),
	"rancher2_setting":                     ReplaceGroupWords("", 0),
	"rancher2_storage_class_v2":            ReplaceGroupWords("cluster", 0),
	"rancher2_token":                       ReplaceGroupWords("admin", 0),
	"rancher2_user":                        ReplaceGroupWords("", 0),
}

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		words := strings.Split(strings.TrimPrefix(resource, "rancher2_"), "_")
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{
	"rancher2_app_v2":                      "App",
	"rancher2_auth_config_activedirectory": "ActiveDirectory",
	"rancher2_auth_config_adfs":            "ADFS",
	"rancher2_auth_config_azuread":         "AzureAD",
	"rancher2_auth_config_freeipa":         "FreeIPA",
	"rancher2_auth_config_github":          "GitHub",
	"rancher2_auth_config_openldap":        "OpenLDAP",
	"rancher2_catalog_v2":                  "Catalog",
	"rancher2_cloud_credential":            "CloudCredential",
	"rancher2_cluster_v2":                  "Cluster",
	"rancher2_config_map_v2":               "ConfigMap",
	"rancher2_machine_config_v2":           "MachineConfig",
	"rancher2_namespace":                   "RancherNamespace",
	"rancher2_secret_v2":                   "Secret",
	"rancher2_storage_class_v2":            "StorageClass",
}

// RancherAssignedIdentifier is a list of regex for the resources whose identifier
// are assigned by Rancher
var RancherAssignedIdentifier = []string{
	"rancher2_auth_config_.+$",
	"rancher2_cloud_credential$",
	"rancher2_cluster$",
	"rancher2_cluster_v2$",
	"rancher2_project$",
}

// GroupKindOverrides overrides the group and kind of the resource if it matches
// any entry in the GroupMap.
func GroupKindOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		}
	}
}

// KindOverrides overrides the kind of the resources given in KindMap.
func KindOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if k, ok := KindMap[r.Name]; ok {
			r.Kind = k
		}
	}
}

// IdentifierAssignedByRancher assigns the Rancher generated ID
// to the resource.
func IdentifierAssignedByRancher() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if matches(r.Name, RancherAssignedIdentifier) {
			r.ExternalName = tjconfig.IdentifierFromProvider
		}
	}
}

// KnownReferencers adds referencers for fields that are known and common among
// more than a few resources.
func KnownReferencers() tjconfig.ResourceOption { //nolint:gocyclo
	return func(r *tjconfig.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}
			switch k {
			case "cluster_id":
				r.References["cluster_id"] = tjconfig.Reference{
					Type: "github.com/crossplane-contrib/provider-jet-rancher/apis/cluster/v1alpha2.Cluster",
				}
			case "project_id":
				r.References["project_id"] = tjconfig.Reference{
					Type: "github.com/crossplane-contrib/provider-jet-rancher/apis/project/v1alpha1.Project",
				}
			}
		}
	}
}

func matches(name string, regexList []string) bool {
	for _, r := range regexList {
		ok, err := regexp.MatchString(r, name)
		if err != nil {
			panic(errors.Wrap(err, "cannot match regular expression"))
		}
		if ok {
			return true
		}
	}
	return false
}
