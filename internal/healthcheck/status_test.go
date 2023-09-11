package healthcheck

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenHTTPCall_WhenGetStatus_ThenReturnsOKStatusCode(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	w := httptest.NewRecorder()

	GetStatus(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_GivenHTTPCall_WhenGetStatus_ThenReturnsHealthyResponseBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	w := httptest.NewRecorder()

	GetStatus(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "healthy", string(data))
}
