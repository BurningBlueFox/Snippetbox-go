package models

import (
	"context"
	"fmt"
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

	created := time.Now()
	expires := created.Add(time.Duration(daysToExpire) * 24 * time.Hour)

	query := fmt.Sprintf("INSERT INTO snippets(title, content, created, expires) VALUES (%s, %s, %s, %s) RETURNING id;",
		title,
		content,
		created,
		expires)

	row := model.DB.QueryRow(context.Background(), query)

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

	query := fmt.Sprintf("SELECT id, title, content, created, expires FROM snippets WHERE id = %d;",
		id)

	row := model.DB.QueryRow(context.Background(), query)
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
	query := fmt.Sprintf("SELECT id, title, content, created, expires FROM snippets ORDER BY created ASC LIMIT %d;",
		amount)

	rows, err := model.DB.Query(context.Background(), query)
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
