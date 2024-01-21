package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

// Insert inserts a new snippet into the SnippetModel database table.
// It takes the title and content as string,
// and the expiration time of the snippet in days. It returns the ID of the newly inserted snippet as an int
// and an error if the insertion fails.
func (model *SnippetModel) Insert(title string, content string, daysToExpire int) (int, error) {
	return 0, nil
}

// Get returns a specific entry in the SnippetModel database table.
// It takes the id of the entry as int,
// Returns the snippet pointer if found one or error entry was not found
func (model *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Return the latest amount number of entries in the db
// Or less than that if there isn't enough entries
// Returns error if any error occurred
func (model *SnippetModel) Latest(amount int) ([]Snippet, error) {
	return nil, nil
}
