package model

type User struct {
	BaseModel
	FirstName   string  `json:"first_name" `
	LastName    string  `json:"last_name"`
	MiddleName  *string `json:"middle_name"`
	Username    string  `json:"username" gorm:"not null;unique"`
	Email       *string `json:"email"`
	Password    string  `json:"-"`
	Phone       *string `json:"phone"`
	LoginHint   *string `json:"login_hint"`
	DateOfBirth *string `json:"date_of_birth"`
	Address     *string `json:"address"`
	Image       *string `json:"image"`
	IsActive    bool    `json:"is_active"`
}
