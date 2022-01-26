package models

type User struct {
	Id       uint   `gorm:"column:id;primaryKey;comment:id;" json:"id"`
	Name     string `gorm:"column:name;type:varchar(120);default:'';not null;comment:用户名;" json:"name"`
	Email    string `gorm:"column:email;unique;type:varchar(120);default:'';not null;comment:邮箱;" json:"email"`
	Password []byte `gorm:"column:password;type:varchar(150);default:'';not null;comment:密码;" json:"-"`
}
