package api

import (
	db "github.com/florian-nguyen/tech-school_simple-bank/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking system
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Adding a validator engine to check on currencies
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency) // name of validation tag and validation function
	}

	// add routes to the router
	// When specifying new routes, the last function should be the handler, other functions should be the middlewares

	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
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
