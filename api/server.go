package api

import (
	"fmt"

	db "github.com/florian-nguyen/tech-school_simple-bank/simple-bank/db/sqlc"
	"github.com/florian-nguyen/tech-school_simple-bank/simple-bank/token"
	"github.com/florian-nguyen/tech-school_simple-bank/simple-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking system
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// To use JWT token instead of Paseto
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	// tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		store:      store,
	}

	// Adding a validator engine to check on currencies
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency) // name of validation tag and validation function
	}

	// add routes to the router
	// When specifying new routes, the last function should be the handler, other functions should be the middlewares

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {

	return server.router.Run(address)

	// NB : The router field is private, so it cannot be accessed from outside this API package.
}

// gin.H is a shortcut for map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}

}
