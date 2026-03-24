package models

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CategoryID   *int      `json:"category_id" db:"category_id"`
	Title        string    `json:"title" db:"title"`
	Username     string    `json:"username" db:"username"`
	URL          string    `json:"url" db:"url"`
	PasswordHash []byte    `json:"-" db:"password_hash"`
	Notes        []byte    `json:"-" db:"notes"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdateAt     time.Time `json:"updated_at" db:"updated_at"`
}

type EntryDisplay struct {
	ID         uuid.UUID `json:"id"`
	CategoryID *int      `json:"category_id"`
	Title      string    `json:"title"`
	Username   string    `json:"username"`
	URL        string    `json:"url"`
	Password   string    `json:"-"`
	Notes      string    `json: "-"`
}

type Category struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	IconName  string    `json:"icon_name" db:"icon_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewEntry(categoryID *int, title, username, url string) *Entry {
	return &Entry{
		ID:         uuid.New(),
		CategoryID: categoryID,
		Title:      title,
		Username:   username,
		URL:        url,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
	}
}

func (e *Entry) ToDisplay(key []byte) (*EntryDisplay, error) {
	display := &EntryDisplay{
		ID:         e.ID,
		CategoryID: e.CategoryID,
		Title:      e.Title,
		Username:   e.Username,
		URL:        e.URL,
	}

	// написать дешифрование пароля и заметок

	return display, nil
}
