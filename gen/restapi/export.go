package restapi

import (
	"net/http"

	"github.com/budougumi0617/sandbox_goswagger/gen/restapi/operations"
)

func ConfigureAPI(api *operations.SampleAPI) http.Handler {
	return configureAPI(api)
}
