package security

import "github.com/swaggest/openapi-go/openapi3"

func AddSecurity(r *openapi3.Reflector) {
	spec := r.Spec
	var bearerFormat = "JWT"
	var mapSecurityScheme = make(map[string]openapi3.SecuritySchemeOrRef)

	securityScheme := openapi3.SecurityScheme{
		HTTPSecurityScheme: &openapi3.HTTPSecurityScheme{
			Scheme:       "bearer",
			BearerFormat: &bearerFormat,
		},
	}

	mapSecurityScheme["BearerAuth"] = openapi3.SecuritySchemeOrRef{
		SecurityScheme: &securityScheme,
	}

	spec.Components.WithSecuritySchemes(openapi3.ComponentsSecuritySchemes{
		MapOfSecuritySchemeOrRefValues: mapSecurityScheme,
	})
}
