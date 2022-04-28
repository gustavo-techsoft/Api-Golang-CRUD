package apicrud

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func getProductsDB(tx *sqlx.Tx) (products []Product, err error) {
	q := `SELECT id, nome_prod, preco FROM produto`

	err = tx.Select(&products, q)

	return
}

func getProductDB(codeProduct int, tx *sqlx.Tx) (product Product, err error) {
	q := `SELECT id, nome_prod, preco FROM produto where id = ?`

	err = tx.Get(&product, q, codeProduct)
	if err == sql.ErrNoRows {
		return
	}

	return
}

func postProductDB(product Product, tx *sqlx.Tx) (err error) {
	q := `INSERT INTO Produto (nome_prod, preco) values (?,?)`

	_, err = tx.Exec(q, product.NameProduct, product.Price)

	return
}

func deleteProductDB(codeProduct int, tx *sqlx.Tx) (err error) {
	q := `DELETE FROM Produto WHERE Id = ?`

	_, err = tx.Exec(q, codeProduct)

	return
}

func putProductDB(product Product, codeProduct int, tx *sqlx.Tx) (err error) {
	q := `UPDATE Produto SET nome_prod = ?, preco = ? WHERE id = ?`

	_, err = tx.Exec(q, product.NameProduct, product.Price, codeProduct)

	return
}
