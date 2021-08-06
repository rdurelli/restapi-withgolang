package models

type Produto struct {
	Id          string  `json:"id,omitempty"`
	Nome        string  `json:"nome,omitempty"`
	Preco       float64 `json:"preco,omitempty"`
	Catergorias []Categoria
}
