package models

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	Name              string  `json:"name"`
	SenderName        string  `json:"sender"`
	RecipientName     string  `json:"receiver"`
	SenderLocation    string  `json:"sender_location"`
	RecipientLocation string  `json:"sender_receiver"`
	Cost              float64 `json:"fee"`
	Weight            float64 `json:"weight"`
}
