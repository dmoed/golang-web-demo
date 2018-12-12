package model

import "database/sql"

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//pagination
func GetAllByPage(db *sql.DB, page int, limit int) ([]Product, error) {

	products := []Product{}
	offset := (page - 1) * limit

	results, err := db.Query("select id, pr_desc from product limit ?,?", offset, limit)

	if err != nil {
		return nil, err
	}

	for results.Next() {

		product := Product{}

		err = results.Scan(&product.ID, &product.Name)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
