package scad

type Initializer interface {
	Initialize()
}

type InitializerDirPathRenderer interface {
	Initializer
	DirPathRenderer
}
