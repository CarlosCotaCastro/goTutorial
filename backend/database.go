package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Database represents the database connection
type Database struct {
	conn *sql.DB
}

// NewDatabase creates a new database connection
func NewDatabase() (*Database, error) {
	// Create data directory if it doesn't exist
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	// Database file path
	dbPath := filepath.Join(dataDir, "progress.db")

	// Open database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	database := &Database{conn: db}

	// Initialize database schema
	if err := database.initSchema(); err != nil {
		return nil, err
	}

	log.Printf("✅ Database connected: %s", dbPath)
	return database, nil
}

// initSchema creates the necessary tables
func (db *Database) initSchema() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS user_progress (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT NOT NULL,
		lesson_id INTEGER NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0,
		completed_at DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(user_id, lesson_id)
	);

	CREATE INDEX IF NOT EXISTS idx_user_progress_user_id ON user_progress(user_id);
	CREATE INDEX IF NOT EXISTS idx_user_progress_lesson_id ON user_progress(lesson_id);
	CREATE INDEX IF NOT EXISTS idx_user_progress_completed ON user_progress(completed);
	`

	_, err := db.conn.Exec(createTableSQL)
	if err != nil {
		return err
	}

	log.Println("✅ Database schema initialized")
	return nil
}

// GetUserProgress retrieves all progress for a user
func (db *Database) GetUserProgress(userID string) ([]UserProgress, error) {
	query := `
		SELECT user_id, lesson_id, completed, completed_at 
		FROM user_progress 
		WHERE user_id = ? 
		ORDER BY lesson_id
	`

	rows, err := db.conn.Query(query, userID)
	if err != nil {
		return []UserProgress{}, err
	}
	defer rows.Close()

	progress := []UserProgress{}
	for rows.Next() {
		var p UserProgress
		var completedAt sql.NullString

		err := rows.Scan(&p.UserID, &p.LessonID, &p.Completed, &completedAt)
		if err != nil {
			return []UserProgress{}, err
		}

		if completedAt.Valid {
			p.CompletedAt = &completedAt.String
		}

		progress = append(progress, p)
	}

	return progress, nil
}

// UpdateUserProgress updates or creates user progress
func (db *Database) UpdateUserProgress(progress UserProgress) error {
	query := `
		INSERT OR REPLACE INTO user_progress 
		(user_id, lesson_id, completed, completed_at, updated_at) 
		VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)
	`

	var completedAt interface{}
	if progress.CompletedAt != nil {
		completedAt = *progress.CompletedAt
	}

	_, err := db.conn.Exec(query, progress.UserID, progress.LessonID, progress.Completed, completedAt)
	return err
}

// Close closes the database connection
func (db *Database) Close() error {
	return db.conn.Close()
}
