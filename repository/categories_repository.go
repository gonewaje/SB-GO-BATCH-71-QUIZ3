package repository

import (
	"database/sql"
	"errors"
	"library/structs"
)

func ListCategories(db *sql.DB) ([]structs.Category, error) {
	rows, err := db.Query(`SELECT id, name FROM categories ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []structs.Category
	for rows.Next() {
		var c structs.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

func GetCategory(db *sql.DB, id int) (structs.Category, error) {
	var c structs.Category
	err := db.QueryRow(`SELECT id, name FROM categories WHERE id=$1`, id).Scan(&c.ID, &c.Name)
	if err == sql.ErrNoRows {
		return c, errors.New("category not found")
	}
	return c, err
}

func CreateCategory(db *sql.DB, name, createdBy string) error {
	_, err := db.Exec(`INSERT INTO categories (name, created_by) VALUES ($1, $2)`, name, createdBy)
	return err
}

func UpdateCategory(db *sql.DB, id int, name, modifiedBy string) (bool, error) {
	res, err := db.Exec(`UPDATE categories SET name=$1, modified_at=NOW(), modified_by=$2 WHERE id=$3`, name, modifiedBy, id)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}

func DeleteCategory(db *sql.DB, id int) (bool, error) {
	res, err := db.Exec(`DELETE FROM categories WHERE id=$1`, id)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}
