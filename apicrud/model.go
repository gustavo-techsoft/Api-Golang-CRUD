package apicrud

type Product struct {
	Id          int     `json:"id" db:"id"`
	NameProduct string  `json:"decription" db:"nome_prod"`
	Price       float64 `json:"price" db:"preco"`
}
