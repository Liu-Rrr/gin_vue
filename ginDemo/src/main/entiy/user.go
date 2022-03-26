package entiy

type User struct {
	ID       int
	Username string
	Password string
}

//防止gorm自动创建表加s

func (User) TableName() string {
	return "user"
}
