package model

import "time"

type ChatImages struct {
	Id         int       `gorm:"id"`
	SenderId   int       `gorm:"sender_id"`
	ReceiverId int       `gorm:"receiver_id"`
	ImageData  string    `gorm:"image_data"`
	SentAt     time.Time `gorm:"sent_at"`
}
type ChatMessages struct {
	Id          int       `gorm:"id"`
	SenderId    int       `gorm:"sender_id"`
	ReceiverId  int       `gorm:"receiver_id"`
	MessageText string    `gorm:"message_text"`
	SentAt      time.Time `gorm:"sent_at"`
}
type User struct {
	Id        int       `gorm:"id"`
	Username  string    `gorm:"username"`
	CreatedAt time.Time `gorm:"created_at"`
}
