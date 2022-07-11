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

	v2 "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/app/v2"
	configactivedirectory "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/auth/configactivedirectory"
	configadfs "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/auth/configadfs"
	configazuread "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/auth/configazuread"
	configfreeipa "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/auth/configfreeipa"
	configgithub "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/auth/configgithub"
	configkeycloak "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/auth/configkeycloak"
	configokta "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/auth/configokta"
	configopenldap "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/auth/configopenldap"
	configping "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/auth/configping"
	v2catalog "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/catalog/v2"
	credential "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cloud/credential"
	alertgroup "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/alertgroup"
	alertrule "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/alertrule"
	driver "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/driver"
	logging "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/logging"
	roletemplatebinding "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/roletemplatebinding"
	sync "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/sync"
	template "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/template"
	v2cluster "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/cluster/v2"
	mapv2 "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/config/mapv2"
	backup "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/etcd/backup"
	dns "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/dns"
	dnsprovider "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/dnsprovider"
	role "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/role"
	rolebinding "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/global/rolebinding"
	configv2 "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/machine/configv2"
	clusterapp "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/multi/clusterapp"
	drivernode "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/node/driver"
	pool "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/node/pool"
	templatenode "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/node/template"
	securitypolicytemplate "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/pod/securitypolicytemplate"
	alertgroupproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/alertgroup"
	alertruleproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/alertrule"
	loggingproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/logging"
	roletemplatebindingproject "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/project/roletemplatebinding"
	providerconfig "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/providerconfig"
	app "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/app"
	bootstrap "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/bootstrap"
	catalog "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/catalog"
	certificate "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/certificate"
	cluster "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/cluster"
	feature "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/feature"
	namespace "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/namespace"
	notifier "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/notifier"
	project "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/project"
	registry "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/registry"
	setting "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/setting"
	token "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/token"
	user "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/rancher2/user"
	templaterole "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/role/template"
	classv2 "github.com/crossplane-contrib/provider-jet-rancher/internal/controller/storage/classv2"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		v2.Setup,
		configactivedirectory.Setup,
		configadfs.Setup,
		configazuread.Setup,
		configfreeipa.Setup,
		configgithub.Setup,
		configkeycloak.Setup,
		configokta.Setup,
		configopenldap.Setup,
		configping.Setup,
		v2catalog.Setup,
		credential.Setup,
		alertgroup.Setup,
		alertrule.Setup,
		driver.Setup,
		logging.Setup,
		roletemplatebinding.Setup,
		sync.Setup,
		template.Setup,
		v2cluster.Setup,
		mapv2.Setup,
		backup.Setup,
		dns.Setup,
		dnsprovider.Setup,
		role.Setup,
		rolebinding.Setup,
		configv2.Setup,
		clusterapp.Setup,
		drivernode.Setup,
		pool.Setup,
		templatenode.Setup,
		securitypolicytemplate.Setup,
		alertgroupproject.Setup,
		alertruleproject.Setup,
		loggingproject.Setup,
		roletemplatebindingproject.Setup,
		providerconfig.Setup,
		app.Setup,
		bootstrap.Setup,
		catalog.Setup,
		certificate.Setup,
		cluster.Setup,
		feature.Setup,
		namespace.Setup,
		notifier.Setup,
		project.Setup,
		registry.Setup,
		setting.Setup,
		token.Setup,
		user.Setup,
		templaterole.Setup,
		classv2.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
