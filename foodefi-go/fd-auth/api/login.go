package api

import (
	"fmt"
	"foodefi-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Username string `json:"username"`
	UserRole string `json:"user-role"`
	Token    string `json:"token"`
}

func (server *Server) login(ctx *gin.Context) {
	var request loginRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("validation error")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.Queries.GetUser(ctx, request.Username)
	if err != nil {
		fmt.Println("username not found")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(&util.InvalidLoginRequest{}))
		return
	}

	if true == util.CheckHash(request.Password, user.Password) {
		response := loginResponse{
			Username: user.Username,
			UserRole: user.Role,
			Token:    util.RandomStringAlphabet(10), //TODO: JWT Token
		}
		fmt.Println("username password matched")
		ctx.JSON(http.StatusOK, response)
		return
	} else {
		fmt.Println("wrong password")
		ctx.JSON(http.StatusBadRequest, errorResponse(&util.InvalidLoginRequest{}))
		return
	}
}
