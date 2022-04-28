package apicrud

import (
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	db, err := sqlx.Connect(DB, CONFIGDB)
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()

	products, err := getProductsDB(tx)
	if err != nil {
		log.Fatalln(err)
	}

	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	db, err := sqlx.Connect(DB, CONFIGDB)
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()

	id, _ := strconv.Atoi(c.Param("id"))

	if id > 0 {
		product, err := getProductDB(id, tx)
		if err != nil {
			//log.Fatalln(err, "teste")
			tx.Rollback()
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Falha ao buscar produto por parametro",
			})
		}

		return c.JSON(http.StatusOK, product)
	} else {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Falha ao buscar produto por parametro",
		})
	}
}

func PostProduct(c echo.Context) error {
	db, err := sqlx.Connect(DB, CONFIGDB)
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()

	product := Product{}
	c.Bind(&product)

	err = postProductDB(product, tx)
	if err != nil {
		log.Fatalln(err)
		tx.Rollback()
	}
	tx.Commit()
	return c.JSON(http.StatusCreated, product)
}

func DeleteProduct(c echo.Context) error {
	db, err := sqlx.Connect(DB, CONFIGDB)
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()

	id, _ := strconv.Atoi(c.Param("id"))

	err = deleteProductDB(id, tx)
	if err != nil {
		log.Fatalln(err)
		tx.Rollback()
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Falha ao deletar produto",
		})
	}
	tx.Commit()
	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"message": "Produto excluido com sucesso",
	})
}

func PutProduct(c echo.Context) error {
	db, err := sqlx.Connect(DB, CONFIGDB)
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()

	product := Product{}
	c.Bind(&product)

	id, _ := strconv.Atoi(c.Param("id"))

	err = putProductDB(product, id, tx)
	if err != nil {
		log.Fatalln(err)
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Erro ao alterar com produto",
		})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Produto alterado com sucesso",
	})
}
