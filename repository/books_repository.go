package repository

import (
	"database/sql"
	"errors"
	"library/structs"
)

func ListBooks(db *sql.DB) ([]structs.Book, error) {
	rows, err := db.Query(`
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id
		FROM books ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []structs.Book
	for rows.Next() {
		var b structs.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID); err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}

func GetBook(db *sql.DB, id int) (structs.Book, error) {
	var b structs.Book
	err := db.QueryRow(`
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id
		FROM books WHERE id=$1`, id).
		Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID)
	if err == sql.ErrNoRows {
		return b, errors.New("book not found")
	}
	return b, err
}

func CreateBook(db *sql.DB, in structs.Book, createdBy string) error {
	_, err := db.Exec(`
		INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		in.Title, in.Description, in.ImageURL, in.ReleaseYear, in.Price, in.TotalPage, in.Thickness, in.CategoryID, createdBy)
	return err
}

func DeleteBook(db *sql.DB, id int) (bool, error) {
	res, err := db.Exec(`DELETE FROM books WHERE id=$1`, id)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}

func ListBooksByCategory(db *sql.DB, categoryID int) ([]structs.Book, error) {
	rows, err := db.Query(`
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id
		FROM books WHERE category_id=$1`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []structs.Book
	for rows.Next() {
		var b structs.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID); err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}

func GetBookByID(db *sql.DB, id int) (structs.Book, error) {
	var book structs.Book
	query := `
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id
		FROM books
		WHERE id = $1
	`
	err := db.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.ImageURL,
		&book.ReleaseYear,
		&book.Price,
		&book.TotalPage,
		&book.Thickness,
		&book.CategoryID,
	)
	return book, err
}
