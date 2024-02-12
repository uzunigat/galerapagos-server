//go:build unit

package routes_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/config"
	routes "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/sys/routes"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSysHealthRoute(t *testing.T) {
	router := gin.New()

	routes.AttachSysRoutes(router, &config.AppConfig{Name: "test-application"}, &utils.AppStateManager{})

	req := httptest.NewRequest("GET", "/sys/health", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	var response map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &response)

	assert.Equal(t, resp.Code, 200)
	assert.Nil(t, err)
	assert.Equal(t, response["status"], "UP")
}
