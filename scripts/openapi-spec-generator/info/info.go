package info

import "github.com/swaggest/openapi-go/openapi3"

func AddInfo(r *openapi3.Reflector) {
	r.Spec.WithSecurity(map[string][]string{"BearerAuth": []string{}})
	r.Spec.Info.
		WithTitle("ta.go-hexagonal-skeletor").
		WithVersion("1.0.0").
		WithDescription("Go Layered Skeletor Service REST API spec")
}
