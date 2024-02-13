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
	const productionDNS = "go-hexagonal-skeletor.audibene.net"

	serverDNSVariable := openapi3.ServerVariable{
		Description: nil,
		Default:     productionDNS,
		Enum:        []string{productionDNS},
	}

	versionVariable := openapi3.ServerVariable{
		Description: nil,
		Default:     "v1",
		Enum:        []string{"v1"},
	}

	serverVariables := map[string]openapi3.ServerVariable{serverDNSVariableName: serverDNSVariable, versionVariableName: versionVariable}

	spec.Servers = append(spec.Servers, NewServer(localURI, "Local", serverVariables))
	spec.Servers = append(spec.Servers, NewServer(baseServerPath, "Production", serverVariables))
}

func NewServer(url string, description string, variables map[string]openapi3.ServerVariable) openapi3.Server {
	return openapi3.Server{
		URL:         url,
		Description: &description,
		Variables:   variables,
	}
}
