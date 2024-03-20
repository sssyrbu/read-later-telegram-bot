// Database manipulations logic
package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitializeSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "articles_db.db")
	if err != nil {
		return nil, err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS articles (
		user_id BIGINT,
		article TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func AddArticle(db *sql.DB, userID int64, article string) error {
	insertSQL := `
	INSERT INTO articles(user_id, article) VALUES (?, ?);`

	_, err := db.Exec(insertSQL, userID, article)
	if err != nil {
		return err
	}

	return nil
}

func GetUserArticles(db *sql.DB, userID int64) ([]string, error) {
	query := `SELECT article FROM articles WHERE user_id = ?;`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []string
	for rows.Next() {
		var article string
		if err := rows.Scan(&article); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func GetUserIDs(db *sql.DB) ([]int64, error) {
	query := `SELECT user_id FROM articles;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idSet := make(map[int64]bool)
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		idSet[id] = true
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(idSet))
	for id := range idSet {
		ids = append(ids, id)
	}

	return ids, nil
}
