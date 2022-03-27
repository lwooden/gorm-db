package Models

import (
	"gorm.io/gorm"
)

type ExternalVerse struct {
	Reference        string   `json:"reference"`
	Verses           []Verses `json:"verses"`
	Text             string   `json:"text"`
	Translation_ID   string   `json:"translation_id"`
	Translation_Name string   `json:"translation_name"`
	Translation_Note string   `json:"translation_note"`
}

type Verses struct {
	Book_ID   string `json:"book_id"`
	Book_Name string `json:"book_name"`
	Chapter   int    `json:"chapter"`
	Verse     int    `json:"verse"`
	Text      string `json:"text"`
}

type Verse struct {
	gorm.Model
	ID               int    `json:"id" gorm:"primaryKey"`
	Book             string `json:"book"`
	Chapter          int    `json:"chapter"`
	Verse            int    `json:"verse"`
	Text             string `json:"text"`
	Translation_Name string `json:"translation_name"`
	CategoryID       int    `json:"categoryId"`
	// Category         Category
}

// Struct used to capture a verse from a HTTP POST request body (JSON)
type VerseDAO struct {
	Book             string `json:"book" binding:"required"`
	Chapter          int    `json:"chapter" binding:"required"`
	Verse            int    `json:"verse" binding:"required"`
	Text             string `json:"text" binding:"required"`
	Translation_Name string `json:"translation_name" binding:"required"`
	CategoryID       int    `json:"categoryId" binding:"required"`
}
