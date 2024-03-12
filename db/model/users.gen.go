// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "Users"

// User mapped from table <Users>
type User struct {
	UserID        string  `gorm:"column:user_id;type:varchar(255);primaryKey" json:"user_id"`
	FirebaseUID   string  `gorm:"column:firebase_uid;type:varchar(255);not null;uniqueIndex:firebase_uid,priority:1" json:"firebase_uid"`
	Username      string  `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Firstname     string  `gorm:"column:firstname;type:varchar(255);not null" json:"firstname"`
	Lastname      string  `gorm:"column:lastname;type:varchar(255);not null" json:"lastname"`
	FirstnameKana string  `gorm:"column:firstname_kana;type:varchar(255);not null" json:"firstname_kana"`
	LastnameKana  string  `gorm:"column:lastname_kana;type:varchar(255);not null" json:"lastname_kana"`
	StatusMessage string  `gorm:"column:status_message;type:text;not null" json:"status_message"`
	IconImageHash *string `gorm:"column:icon_image_hash;type:varchar(255)" json:"icon_image_hash"`
	Tags          []Tag   `gorm:"many2many:user_tags" json:"tags"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
