package entity

type Response struct {
	Entradas [][]interface{} `json:"entradas"`
	Targets  []string        `json:"targets"`
}
type Request struct {
	Nro_filas int `json:"nro_filas"`
}
