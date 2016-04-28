package models

import "time"

type Pet struct {
	Id      int       `json:"_id"`
	Name    string    `json:"name"`
	Added   time.Time `json:"added"`
	Removed time.Time `json:"removed"`
}

type Pets []Pet
