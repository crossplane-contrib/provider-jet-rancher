package secret

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane-contrib/provider-jet-rancher/config/common"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/pkg/errors"
)

// Configure configures the rancher2_secret_v2 groups
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("rancher2_secret_v2", func(r *tjconfig.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = tjconfig.NameAsIdentifier
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			cn, ok := parameters["cluster_id"].(string)
			if !ok || cn == "" {
				return "", errors.New("cannot get cluster_id from parameters")
			}
			ns, ok := parameters["namespace"].(string)
			if !ok || ns == "" {
				return "", errors.New("cannot get namespace from parameters")
			}
			return fmt.Sprintf("%s.%s/%s", cn, ns, externalName), nil
		}
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			id, ok := tfstate["id"]
			if !ok {
				return "", errors.Errorf("id not found")
			}
			idStr, ok := id.(string)
			if !ok {
				return "", errors.Errorf("id is not of type string")
			}
			words := strings.Split(idStr, "/")
			return words[len(words)-1], nil
		}
	})
}
