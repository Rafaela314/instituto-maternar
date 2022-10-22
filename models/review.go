package models

import (
	"time"
)

//Review represent review model
type Review struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Title      string    `json:"title" validate:"required"`
	Content    string    `json:"content" validate:"required"`
	Date       string    `json:"date"`
	City       int       `json:"city"`
	State      int       `json:"state"`
	Place      string    `json:"place"`
	PlaceRate  int       `json:"place_rate"`
	Doctor     string    `json:"doctor"`
	DoctorRate int       `json:"doctor_rate"`
	Doula      string    `json:"doula"`
	DoulaRate  int       `json:"doula_rate"`
	Team       string    `json:"team"`
	TeamRate   int       `json:"team_rate"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
	Image      string    `json:"image"`
}
