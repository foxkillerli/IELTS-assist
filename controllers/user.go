package controllers

import (
	"encoding/json"
	"github.com/foxkillerli/IELTS-assist/db"
	"github.com/foxkillerli/IELTS-assist/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *gin.Context) {
	var user models.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	database := db.GetInstance()
	database.Create(&user)
	c.JSON(200, gin.H{"message": "success"})
}

func UserLogin(c *gin.Context) {

}
