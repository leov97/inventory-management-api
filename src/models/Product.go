package models

type Procuct struct {
	ID     int     `json: id`
	Nombre string  `json: nombre`
	precio float64 `json: precio`
}
