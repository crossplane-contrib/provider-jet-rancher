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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	token "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/admin/token"
	app "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/app/app"
	appappv2 "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/appv2/app"
	multiclusterapp "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/appv2/multiclusterapp"
	activedirectory "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/activedirectory"
	adfs "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/adfs"
	azuread "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/azuread"
	freeipa "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/freeipa"
	github "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/github"
	keycloak "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/keycloak"
	okta "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/okta"
	openldap "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/openldap"
	ping "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/authconfig/ping"
	catalog "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/catalog/catalog"
	catalogcatalogv2 "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/catalogv2/catalog"
	alertgroup "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/alertgroup"
	alertrule "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/alertrule"
	cluster "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/cluster"
	configmap "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/configmap"
	driver "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/driver"
	etcdbackup "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/etcdbackup"
	logging "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/logging"
	notifier "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/notifier"
	roletemplatebinding "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/roletemplatebinding"
	storageclass "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/storageclass"
	sync "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/sync"
	template "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/template"
	clusterclusterv2 "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/clusterv2/cluster"
	dns "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/dns"
	dnsprovider "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/dnsprovider"
	role "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/role"
	rolebinding "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/rolebinding"
	nodedriver "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/node/nodedriver"
	nodepool "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/node/nodepool"
	nodetemplate "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/node/nodetemplate"
	alertgroupproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/alertgroup"
	alertruleproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/alertrule"
	certificate "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/certificate"
	loggingproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/logging"
	project "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/project"
	ranchernamespace "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/ranchernamespace"
	registry "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/registry"
	roletemplatebindingproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/roletemplatebinding"
	providerconfig "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/providerconfig"
	cloudcredential "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher/cloudcredential"
	feature "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher/feature"
	machineconfig "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher/machineconfig"
	roletemplate "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher/roletemplate"
	setting "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher/setting"
	user "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher/user"
	secret "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/secret/secret"
	secretsecretv2 "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/secretv2/secret"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		token.Setup,
		app.Setup,
		appappv2.Setup,
		multiclusterapp.Setup,
		activedirectory.Setup,
		adfs.Setup,
		azuread.Setup,
		freeipa.Setup,
		github.Setup,
		keycloak.Setup,
		okta.Setup,
		openldap.Setup,
		ping.Setup,
		catalog.Setup,
		catalogcatalogv2.Setup,
		alertgroup.Setup,
		alertrule.Setup,
		cluster.Setup,
		configmap.Setup,
		driver.Setup,
		etcdbackup.Setup,
		logging.Setup,
		notifier.Setup,
		roletemplatebinding.Setup,
		storageclass.Setup,
		sync.Setup,
		template.Setup,
		clusterclusterv2.Setup,
		dns.Setup,
		dnsprovider.Setup,
		role.Setup,
		rolebinding.Setup,
		nodedriver.Setup,
		nodepool.Setup,
		nodetemplate.Setup,
		alertgroupproject.Setup,
		alertruleproject.Setup,
		certificate.Setup,
		loggingproject.Setup,
		project.Setup,
		ranchernamespace.Setup,
		registry.Setup,
		roletemplatebindingproject.Setup,
		providerconfig.Setup,
		cloudcredential.Setup,
		feature.Setup,
		machineconfig.Setup,
		roletemplate.Setup,
		setting.Setup,
		user.Setup,
		secret.Setup,
		secretsecretv2.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
