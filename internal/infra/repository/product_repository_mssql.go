package repository

import (
	"api/internal/entity"
	"database/sql"
	"fmt"
)

type ProductRepositoryMssql struct {
	DB *sql.DB
}

func NewProductRepositoryMssql(db *sql.DB) *ProductRepositoryMssql {
	return &ProductRepositoryMssql{DB: db}
}

func (r *ProductRepositoryMssql) Create(product *entity.Product) error {
	fmt.Println("AGORA VAI")
	fmt.Println(product.ID)
	fmt.Println(product.Name)
	fmt.Println(product.Price)
	_, err := r.DB.Exec("Insert into products (id, name, price) values(@P1,@P2,@P3)",
		product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryMssql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("select id, name, price from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)

	}

	return products, nil

}
