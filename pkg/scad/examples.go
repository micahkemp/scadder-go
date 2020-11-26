package scad

type ExampleTypes map[string]Initializer

var Examples = ExampleTypes{
	"cube": &CubeExample,
}
