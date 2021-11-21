package store

type Dataset struct {
	Entradas [][]interface{} `json:"entradas"`
	Targets  []string        `json:"targets"`
}

var Data *Dataset
