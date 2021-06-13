package api

import (
	"net/http"
	"time"

	db "github.com/florian-nguyen/tech-school_simple-bank/simple-bank/db/sqlc"
	"github.com/florian-nguyen/tech-school_simple-bank/simple-bank/db/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"` //alphanum : username can only contain ASCII characters
	Password string `json:"password" binding:"required,min=6"`    // minimum password length of six characters
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	// see Gin documentation for binding keyword and possible arguments
}

type createUserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

// createUser
// gin.Context is used to store useful information. This context is used all the time in frameworks such as Gin (or Fiber).
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	// Gin data validation - returns an error if data specified by user is invalid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Response without hashed password is sent back
	resp := createUserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}
