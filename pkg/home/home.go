package model

/*import (
	"shared/src"
)*/

type Model struct {
	Name string
}

func Init(name string) Model {
	var mdl Model
	mdl.Name = name
	return mdl
}
