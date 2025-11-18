package database

import "database/sql"

type EventModel struct {
	DB *sql.DB
}

type Events struct {
	Id          int    `json:"id"`
	OwnerId     string `json:"OwnerId" binding:"required"`
	Name        string `json:"name" binding:"required, min=3"`
	Description string `json:"Description" binding:"required, min=10"`
	Date        string `json:"Date" binding:"required, datetime=2006-01-02"`
	Location    string `json:"Location" binding:"required, min=3"`
}
