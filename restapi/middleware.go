package restapi

import (
	"net/http"

	"bitbucket.org/iandkenzt/inventories/utils"
)

// AppSecretKeyMiddleware authenticate request using application secret key
func AppSecretKeyMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		apiKeyHeader := req.Header.Get("X-Api-Key")
		isAppKeyValid := ValidateAppSecretKey(apiKeyHeader)

		if apiKeyHeader == "" || isAppKeyValid == false {
			utils.Logger.Errorf("X-Api-Key is not valid")
			utils.SendJSONResponse(res, http.StatusForbidden, "forbiden", nil, http.StatusBadRequest)

			return
		}

		f(res, req)
	}
}
