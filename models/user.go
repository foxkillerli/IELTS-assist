package models

import "time"

type User struct {
	BaseModel
	UserName     string    `gorm:"default ''" json:"username"`
	Password     string    `gorm:"default ''" json:"-"`
	Email        string    `gorm:"default ''" json:"email"`
	Mobile       string    `gorm:"not null" json:"mobile"`
	LastLogin    time.Time `json:"-" json:"last_login"`
	WeChatOpenId string    `gorm:"default ''" json:"-"`
}
