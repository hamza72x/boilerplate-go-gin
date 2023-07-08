package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) authAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO:- Check if the user is authenticated and is an admin
	}
}

func (s *Server) auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO:- Check if the user is authenticated
	}
}
