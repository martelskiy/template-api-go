package healthcheck

import (
	"net/http"
)

// @Summary      Application healthcheck
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /status [get]
func GetStatus(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
	statusResponse := "healthy"
	responseWriter.Write([]byte(statusResponse))
}
