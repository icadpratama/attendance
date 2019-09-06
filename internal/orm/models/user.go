package models

type User struct {
	BaseModelSoftDelete
	Email string  `gorm:"not null;unique_index:idx_email"`
	UserID *string
	Name *string
	NickName *string
	FirstName *string
	LastName *string
	Location *string
	Description *string
}