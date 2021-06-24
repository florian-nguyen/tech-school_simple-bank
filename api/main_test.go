package api

import (
	"os"
	"testing"
	"time"

	db "github.com/florian-nguyen/tech-school_simple-bank/simple-bank/db/sqlc"
	util "github.com/florian-nguyen/tech-school_simple-bank/simple-bank/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // blank identifier required to avoid automatic delete on save
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {

	// Gin is configured to debug mode by default, so we set it to test mode for tests
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
