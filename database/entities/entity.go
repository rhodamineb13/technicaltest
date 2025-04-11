package entities

import "time"

type MyClient struct {
	Id           int       `gorm:"type:GENERATED ALWAYS AS IDENTITY;primaryKey;not null"`
	Name         string    `gorm:"type:char(250); not null"`
	Slug         string    `gorm:"type:char(100); not null"`
	IsProject    string    `gorm:"column:is_project;type:varchar(30);check:is_project in ('0', '1');not null;default:'0'"`
	SelfCapture  string    `gorm:"column:self_capture;type:char(1);default:'1'"`
	ClientPrefix string    `gorm:"column:client_prefix;type:char(4);not null"`
	ClientLogo   string    `gorm:"column:client_logo;type:char(255);not null;default:'no-image.jpg'"`
	Address      string    `gorm:"column:address;type:text;default:NULL"`
	PhoneNumber  string    `gorm:"column:phone_number;type:char(50);default:NULL"`
	City         string    `gorm:"column:city;type:char(50);default:NULL"`
	CreatedAt    time.Time `gorm:"type:timestamp(0);default:NULL"`
	UpdatedAt    time.Time `gorm:"type:timestamp(0);default:NULL"`
	DeletedAt    time.Time `gorm:"type:timestamp(0);default:NULL"`
}
