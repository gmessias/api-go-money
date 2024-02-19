package models

type Cash struct {
	Movimentacao string  `json:"movimentacao"`
	Descricao    string  `json:"descricao"`
	Valor        float64 `json:"valor"`
	Categoria    string  `json:"categoria"`
	Origem       string  `json:"origem"`
}
