package scad

type Module interface {
	SCADWriter() SCADWriter
}

type transformable struct {
	Module
}

func Transformable(m Module) transformable {
	return transformable{
		Module: m,
	}
}
