package server

import (

	//"database/sql"
	"net/http"
	db "resto-backend/db/sqlc"
	utill "resto-backend/util"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserResponse struct {
	Username          string           `json:"username"`
	FullName          string           `json:"fullname"`
	Email             string           `json:"email"`
	PasswordChangedAt pgtype.Timestamp `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
}
type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	SessionID            uuid.UUID    `json:"session_id"`
	AccessToken          string       `json:"access_token"`
	AccessTokenExpiresAt time.Time    `json:"Access_token_expires_At"`
	RefreshToken         string       `json:"refresh_token"`
	RefreshTokenAt       string       `json:"Refresh_Token_at"`
	User                 UserResponse `json:"user_response"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		//FullName:          user.FullName,
		Email: user.Email,
		//PasswordChangedAt:   user.PasswordChangedAt,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) createUser(ctx echo.Context) error {
	var request CreateUserRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	hashedPassword, err := utill.HashPassword(request.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := db.CreateUserParams{
		Username:     request.Username,
		PasswordHash: hashedPassword,
		Email:        request.Email,
	}

}
