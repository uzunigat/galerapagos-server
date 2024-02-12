package server

import (
	"fmt"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/config"
	"github.com/swaggest/openapi-go/openapi3"
)

const baseServerPath = "https://{ServerDNS}/api/{version}"
const serverDNSVariableName = "ServerDNS"
const versionVariableName = "version"

func AddServers(r *openapi3.Reflector) {

	config := config.MustLoad()

	spec := r.Spec

	localURI := fmt.Sprintf("%s:%s%s", "http://localhost", config.App.Port, "/api/{version}")
	const testingDNS = "go-hexagonal-skeletor.shared-testing.audibene.net"
	const stagingDNS = "go-hexagonal-skeletor.shared-staging.audibene.net"
	const productionDNS = "go-hexagonal-skeletor.audibene.net"

	serverDNSVariable := openapi3.ServerVariable{
		Description: nil,
		Default:     testingDNS,
		Enum:        []string{testingDNS, stagingDNS, productionDNS},
	}

	versionVariable := openapi3.ServerVariable{
		Description: nil,
		Default:     "v1",
		Enum:        []string{"v1"},
	}

	serverVariables := map[string]openapi3.ServerVariable{serverDNSVariableName: serverDNSVariable, versionVariableName: versionVariable}

	spec.Servers = append(spec.Servers, NewServer(localURI, "Local", serverVariables))
	spec.Servers = append(spec.Servers, NewServer(baseServerPath, "Testing", serverVariables))
	spec.Servers = append(spec.Servers, NewServer(baseServerPath, "Staging", serverVariables))
	spec.Servers = append(spec.Servers, NewServer(baseServerPath, "Production", serverVariables))
}

func NewServer(url string, description string, variables map[string]openapi3.ServerVariable) openapi3.Server {
	return openapi3.Server{
		URL:         url,
		Description: &description,
		Variables:   variables,
	}
}
