package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func HashPassword(password string) string {
	fmt.Println("got -- go func temp")
}

func VerifyPassword(userPasssword string, providerPassword string) (bool, string) {

}
