package main

import (
	"database/sql"
	"encoding/json"
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

	// Initialize lessons database
	if err := db.initLessonsDatabase(); err != nil {
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

// initLessonsDatabase initializes the lessons database
func (db *Database) initLessonsDatabase() error {
	// Open lessons database
	lessonsDBPath := filepath.Join("data", "lessons.db")
	lessonsDB, err := sql.Open("sqlite3", lessonsDBPath)
	if err != nil {
		return err
	}
	defer lessonsDB.Close()

	// Create lessons table
	createLessonsTable := `
	CREATE TABLE IF NOT EXISTS lessons (
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		content TEXT NOT NULL,
		explanation TEXT NOT NULL,
		variants TEXT NOT NULL,
		exercise TEXT NOT NULL,
		solution TEXT NOT NULL,
		difficulty TEXT NOT NULL,
		order_index INTEGER NOT NULL,
		category TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = lessonsDB.Exec(createLessonsTable)
	if err != nil {
		return err
	}

	// Check if lessons table is empty and populate it
	var count int
	err = lessonsDB.QueryRow("SELECT COUNT(*) FROM lessons").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		log.Println("📚 Populating lessons database...")
		return db.populateLessonsDatabase(lessonsDB)
	}

	log.Println("✅ Lessons database initialized")
	return nil
}

// populateLessonsDatabase populates the lessons database with initial data
func (db *Database) populateLessonsDatabase(lessonsDB *sql.DB) error {
	lessons := getTutorialLessons()

	stmt, err := lessonsDB.Prepare(`
		INSERT INTO lessons (id, title, description, content, explanation, variants, exercise, solution, difficulty, order_index, category)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, lesson := range lessons {
		// Convert variants slice to JSON string
		variantsJSON, err := json.Marshal(lesson.Variants)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(
			lesson.ID,
			lesson.Title,
			lesson.Description,
			lesson.Content,
			lesson.Explanation,
			string(variantsJSON),
			lesson.Exercise,
			lesson.Solution,
			lesson.Difficulty,
			lesson.Order,
			lesson.Category,
		)
		if err != nil {
			return err
		}
	}

	log.Println("✅ Lessons database populated successfully")
	return nil
}

// GetLessons retrieves all lessons from the database
func (db *Database) GetLessons() ([]Lesson, error) {
	lessonsDBPath := filepath.Join("data", "lessons.db")
	lessonsDB, err := sql.Open("sqlite3", lessonsDBPath)
	if err != nil {
		return nil, err
	}
	defer lessonsDB.Close()

	query := `
		SELECT id, title, description, content, explanation, variants, exercise, solution, difficulty, order_index, category
		FROM lessons 
		ORDER BY order_index
	`

	rows, err := lessonsDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []Lesson
	for rows.Next() {
		var lesson Lesson
		var variantsJSON string

		err := rows.Scan(
			&lesson.ID,
			&lesson.Title,
			&lesson.Description,
			&lesson.Content,
			&lesson.Explanation,
			&variantsJSON,
			&lesson.Exercise,
			&lesson.Solution,
			&lesson.Difficulty,
			&lesson.Order,
			&lesson.Category,
		)
		if err != nil {
			return nil, err
		}

		// Parse variants JSON
		err = json.Unmarshal([]byte(variantsJSON), &lesson.Variants)
		if err != nil {
			return nil, err
		}

		lessons = append(lessons, lesson)
	}

	return lessons, nil
}

// GetLesson retrieves a specific lesson by ID from the database
func (db *Database) GetLesson(id int) (*Lesson, error) {
	lessonsDBPath := filepath.Join("data", "lessons.db")
	lessonsDB, err := sql.Open("sqlite3", lessonsDBPath)
	if err != nil {
		return nil, err
	}
	defer lessonsDB.Close()

	query := `
		SELECT id, title, description, content, explanation, variants, exercise, solution, difficulty, order_index, category
		FROM lessons 
		WHERE id = ?
	`

	var lesson Lesson
	var variantsJSON string

	err = lessonsDB.QueryRow(query, id).Scan(
		&lesson.ID,
		&lesson.Title,
		&lesson.Description,
		&lesson.Content,
		&lesson.Explanation,
		&variantsJSON,
		&lesson.Exercise,
		&lesson.Solution,
		&lesson.Difficulty,
		&lesson.Order,
		&lesson.Category,
	)
	if err != nil {
		return nil, err
	}

	// Parse variants JSON
	err = json.Unmarshal([]byte(variantsJSON), &lesson.Variants)
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}
