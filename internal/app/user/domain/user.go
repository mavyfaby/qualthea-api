package domain

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;not null;type:int(11)"`
	FirstName string    `gorm:"column:first_name;not null;type:varchar(100)"`
	LastName  string    `gorm:"column:last_name;not null;type:varchar(100)"`
	BirthDate time.Time `gorm:"column:birth_date;not null;type:date"`
	Email     string    `gorm:"column:email;not null;type:varchar(100)"`
	Username  string    `gorm:"column:username;not null;type:varchar(50)"`
	Password  string    `gorm:"column:password;not null;type:varchar(100)"`
	CreatedAt time.Time `gorm:"column:created_at;not null;type:datetime"`
}
