package info

import "github.com/swaggest/openapi-go/openapi3"

func AddInfo(r *openapi3.Reflector) {
	r.Spec.WithSecurity(map[string][]string{"BearerAuth": []string{}})
	r.Spec.Info.
		WithTitle("galerapagos-server").
		WithVersion("1.0.0").
		WithDescription("Galerapagos Server REST API spec")
}
