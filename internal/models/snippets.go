package models

import (
	"database/sql"
	"errors"
	"time"
)

// Snippet type to hold the data for an individual snippet.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) 
			 VALUES($1, $2, NOW(), NOW() + INTERVAL '1 DAY' * $3) RETURNING id`

	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
			 WHERE expires > NOW() AND id = $1`

	// Initialize a pointer to a new zeroed Snippet struct.
	s := &Snippet{}

	// Use the QueryRow() method on the connection pool to execute our SQL statement
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// If the query returns no rows, return ErrNoRecord
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	// If everything went OK then return the Snippet object.
	return s, nil
}

// Latest get the latest 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	// Write the SQL statement we want to execute.
	stmt := `SELECT id, title, content, created, expires FROM snippets
			 WHERE expires > NOW() ORDER BY id DESC LIMIT 10`
	// Use the Query() method to execute our SQL statement.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// Defer rows.Close() to ensure the sql.Rows resultset is
	// always properly closed before the Latest() method returns.
	defer rows.Close()

	// Initialize an empty slice to hold the Snippet structs.
	snippets := []*Snippet{}
	// Use rows.Next to iterate through the rows in the resultset.
	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct.
		s := &Snippet{}
		// Use rows.Scan() to copy the values from each field in the row to the new Snippet object that we created.
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets.
		snippets = append(snippets, s)
	}
	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
	// error that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// return the Snippets slice.
	return snippets, nil
}
