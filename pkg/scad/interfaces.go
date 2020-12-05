package scad

type Module interface {
	SCADWriter() SCADWriter
}
