package utils

import (
	"github.com/foxkillerli/IELTS-assist/db"
	"os/user"
)

func Migrate() {
	database := db.GetInstance()
	database.AutoMigrate(&user.User{})
}
