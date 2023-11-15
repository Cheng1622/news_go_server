package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Userid       int64   `gorm:"type:varchar(255);not null;index:,unique" json:"userid,string"`
	Username     string  `gorm:"type:varchar(20);not null;unique" json:"username"`
	Password     string  `gorm:"size:1000;not null" json:"password"`
	Mobile       string  `gorm:"type:varchar(11);not null;unique" json:"mobile"`
	Email        string  `gorm:"type:varchar(255);not null;unique" json:"email"`
	Avatar       string  `gorm:"type:varchar(255)" json:"avatar"`
	Nickname     *string `gorm:"type:varchar(20)" json:"nickname"`
	Introduction *string `gorm:"type:varchar(255)" json:"introduction"`
	Status       uint    `gorm:"type:tinyint(1);default:1;comment:'1正常, 2禁用'" json:"status"`
	Creator      string  `gorm:"type:varchar(20);" json:"creator"`
	Roles        []*Role `gorm:"many2many:user_role;ForeignKey:Userid;AssociationForeignKey:Roleid" json:"roles"`
}
