package server

import (
	"errors"
	"github.com/hamza72x/go-gin-gorm/accounts"
	"github.com/hamza72x/go-gin-gorm/server/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) admin_createAccount(c *gin.Context) {
	req := &request.Admin_CreateUser{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc := &accounts.Account{
		Name:        req.Name,
		AccountType: accounts.USER,
	}

	if err := s.db.Create(acc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, acc)
}

func (s *Server) admin_listAccount(c *gin.Context) {
	accs := []*accounts.Account{}

	if err := s.db.Find(&accs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accs)
}

func (s *Server) admin_getAccount(c *gin.Context) {
	acc := &accounts.Account{}

	if err := s.db.First(acc, c.Param("id")).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, acc)
}
