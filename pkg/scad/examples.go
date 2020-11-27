package scad

type ExampleTypes map[string]InitializerDirPathRenderer

var Examples = ExampleTypes{
	"cube": &CubeExample,
}
