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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Certificate
func (mg *Certificate) GetTerraformResourceType() string {
	return "rancher2_certificate"
}

// GetConnectionDetailsMapping for this Certificate
func (tr *Certificate) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"key": "spec.forProvider.keySecretRef"}
}

// GetObservation of this Certificate
func (tr *Certificate) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Certificate
func (tr *Certificate) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Certificate
func (tr *Certificate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Certificate
func (tr *Certificate) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Certificate
func (tr *Certificate) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Certificate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Certificate) LateInitialize(attrs []byte) (bool, error) {
	params := &CertificateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Certificate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this RancherNamespace
func (mg *RancherNamespace) GetTerraformResourceType() string {
	return "rancher2_namespace"
}

// GetConnectionDetailsMapping for this RancherNamespace
func (tr *RancherNamespace) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this RancherNamespace
func (tr *RancherNamespace) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RancherNamespace
func (tr *RancherNamespace) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RancherNamespace
func (tr *RancherNamespace) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RancherNamespace
func (tr *RancherNamespace) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RancherNamespace
func (tr *RancherNamespace) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RancherNamespace using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RancherNamespace) LateInitialize(attrs []byte) (bool, error) {
	params := &RancherNamespaceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RancherNamespace) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Project
func (mg *Project) GetTerraformResourceType() string {
	return "rancher2_project"
}

// GetConnectionDetailsMapping for this Project
func (tr *Project) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Project
func (tr *Project) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Project
func (tr *Project) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Project
func (tr *Project) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Project
func (tr *Project) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Project
func (tr *Project) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Project using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Project) LateInitialize(attrs []byte) (bool, error) {
	params := &ProjectParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Project) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AlertGroup
func (mg *AlertGroup) GetTerraformResourceType() string {
	return "rancher2_project_alert_group"
}

// GetConnectionDetailsMapping for this AlertGroup
func (tr *AlertGroup) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AlertGroup
func (tr *AlertGroup) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AlertGroup
func (tr *AlertGroup) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AlertGroup
func (tr *AlertGroup) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AlertGroup
func (tr *AlertGroup) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AlertGroup
func (tr *AlertGroup) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AlertGroup using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AlertGroup) LateInitialize(attrs []byte) (bool, error) {
	params := &AlertGroupParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AlertGroup) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AlertRule
func (mg *AlertRule) GetTerraformResourceType() string {
	return "rancher2_project_alert_rule"
}

// GetConnectionDetailsMapping for this AlertRule
func (tr *AlertRule) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AlertRule
func (tr *AlertRule) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AlertRule
func (tr *AlertRule) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AlertRule
func (tr *AlertRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AlertRule
func (tr *AlertRule) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AlertRule
func (tr *AlertRule) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AlertRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AlertRule) LateInitialize(attrs []byte) (bool, error) {
	params := &AlertRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AlertRule) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Logging
func (mg *Logging) GetTerraformResourceType() string {
	return "rancher2_project_logging"
}

// GetConnectionDetailsMapping for this Logging
func (tr *Logging) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"custom_target_config[*].certificate": "spec.forProvider.customTargetConfig[*].certificateSecretRef", "custom_target_config[*].client_cert": "spec.forProvider.customTargetConfig[*].clientCertSecretRef", "custom_target_config[*].client_key": "spec.forProvider.customTargetConfig[*].clientKeySecretRef", "elasticsearch_config[*].auth_password": "spec.forProvider.elasticsearchConfig[*].authPasswordSecretRef", "elasticsearch_config[*].auth_username": "spec.forProvider.elasticsearchConfig[*].authUsernameSecretRef", "elasticsearch_config[*].certificate": "spec.forProvider.elasticsearchConfig[*].certificateSecretRef", "elasticsearch_config[*].client_cert": "spec.forProvider.elasticsearchConfig[*].clientCertSecretRef", "elasticsearch_config[*].client_key": "spec.forProvider.elasticsearchConfig[*].clientKeySecretRef", "elasticsearch_config[*].client_key_pass": "spec.forProvider.elasticsearchConfig[*].clientKeyPassSecretRef", "fluentd_config[*].certificate": "spec.forProvider.fluentdConfig[*].certificateSecretRef", "fluentd_config[*].fluent_servers[*].password": "spec.forProvider.fluentdConfig[*].fluentServers[*].passwordSecretRef", "fluentd_config[*].fluent_servers[*].shared_key": "spec.forProvider.fluentdConfig[*].fluentServers[*].sharedKeySecretRef", "fluentd_config[*].fluent_servers[*].username": "spec.forProvider.fluentdConfig[*].fluentServers[*].usernameSecretRef", "kafka_config[*].certificate": "spec.forProvider.kafkaConfig[*].certificateSecretRef", "kafka_config[*].client_cert": "spec.forProvider.kafkaConfig[*].clientCertSecretRef", "kafka_config[*].client_key": "spec.forProvider.kafkaConfig[*].clientKeySecretRef", "splunk_config[*].certificate": "spec.forProvider.splunkConfig[*].certificateSecretRef", "splunk_config[*].client_cert": "spec.forProvider.splunkConfig[*].clientCertSecretRef", "splunk_config[*].client_key": "spec.forProvider.splunkConfig[*].clientKeySecretRef", "splunk_config[*].client_key_pass": "spec.forProvider.splunkConfig[*].clientKeyPassSecretRef", "splunk_config[*].token": "spec.forProvider.splunkConfig[*].tokenSecretRef", "syslog_config[*].certificate": "spec.forProvider.syslogConfig[*].certificateSecretRef", "syslog_config[*].client_cert": "spec.forProvider.syslogConfig[*].clientCertSecretRef", "syslog_config[*].client_key": "spec.forProvider.syslogConfig[*].clientKeySecretRef", "syslog_config[*].token": "spec.forProvider.syslogConfig[*].tokenSecretRef"}
}

// GetObservation of this Logging
func (tr *Logging) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Logging
func (tr *Logging) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Logging
func (tr *Logging) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Logging
func (tr *Logging) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Logging
func (tr *Logging) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Logging using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Logging) LateInitialize(attrs []byte) (bool, error) {
	params := &LoggingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Logging) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this RoleTemplateBinding
func (mg *RoleTemplateBinding) GetTerraformResourceType() string {
	return "rancher2_project_role_template_binding"
}

// GetConnectionDetailsMapping for this RoleTemplateBinding
func (tr *RoleTemplateBinding) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this RoleTemplateBinding
func (tr *RoleTemplateBinding) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RoleTemplateBinding
func (tr *RoleTemplateBinding) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RoleTemplateBinding
func (tr *RoleTemplateBinding) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RoleTemplateBinding
func (tr *RoleTemplateBinding) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RoleTemplateBinding
func (tr *RoleTemplateBinding) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RoleTemplateBinding using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RoleTemplateBinding) LateInitialize(attrs []byte) (bool, error) {
	params := &RoleTemplateBindingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RoleTemplateBinding) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Registry
func (mg *Registry) GetTerraformResourceType() string {
	return "rancher2_registry"
}

// GetConnectionDetailsMapping for this Registry
func (tr *Registry) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"registries[*].password": "spec.forProvider.registries[*].passwordSecretRef"}
}

// GetObservation of this Registry
func (tr *Registry) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Registry
func (tr *Registry) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Registry
func (tr *Registry) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Registry
func (tr *Registry) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Registry
func (tr *Registry) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Registry using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Registry) LateInitialize(attrs []byte) (bool, error) {
	params := &RegistryParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Registry) GetTerraformSchemaVersion() int {
	return 0
}
