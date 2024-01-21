package models

import (
	"context"
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
func (model *SnippetModel) Insert(title string, content string, daysToExpire uint) (int, error) {

	query := "INSERT INTO snippets(title, content, created, expires) VALUES ($1, $2, now(), now() + make_interval(days := $3)) RETURNING id;"

	row := model.DB.QueryRow(context.Background(), query, title, content, daysToExpire)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Get returns a specific entry in the SnippetModel database table.
// It takes the id of the entry as int,
// Returns the snippet pointer if found one or error entry was not found
func (model *SnippetModel) Get(id int) (*Snippet, error) {

	query := "SELECT id, title, content, created, expires FROM snippets WHERE id = $1;"

	row := model.DB.QueryRow(context.Background(), query, id)
	snippet := &Snippet{}

	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
	if err != nil {
		return nil, err
	}
	return snippet, nil
}

// Return the latest amount number of entries in the db
// Or less than that if there isn't enough entries
// Returns error if any error occurred
func (model *SnippetModel) Latest(amount int) ([]Snippet, error) {
	query := "SELECT id, title, content, created, expires FROM snippets ORDER BY created ASC LIMIT $1;"

	rows, err := model.DB.Query(context.Background(), query, amount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []Snippet{}
	for rows.Next() {
		snippet := Snippet{}
		err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, snippet)
	}

	return snippets, nil
}
