package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // blank identifier required to avoid automatic delete on save
)

func TestMain(m *testing.M) {

	// Gin is configured to debug mode by default, so we set it to test mode for tests
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
