package product

import (
	"database/sql"

	"github.com/ayinde1993/ecom/types"
)

//repository for products setup to comunicate with DB

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}
	// if products.lenght == 0{

	// }
	return products, nil

}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreateAt,
	)
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (s *Store) CreateProduct(product types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity ) VALUES(?,?,?,?,?)", product.Name, product.Description, product.Image, product.Price, product.Quantity)

	if err != nil {
		return err
	}
	return nil
}
