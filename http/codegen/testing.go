package codegen

import (
	"testing"

	"goa.design/goa/codegen/service"
	httpdesign "goa.design/goa/http/design"
)

// RunHTTPDSL returns the HTTP DSL root resulting from running the given DSL.
func RunHTTPDSL(t *testing.T, dsl func()) *httpdesign.RootExpr {
	// reset all roots and codegen data structures
	service.Services = make(service.ServicesData)
	HTTPServices = make(ServicesData)
	return httpdesign.RunHTTPDSL(t, dsl)
}
