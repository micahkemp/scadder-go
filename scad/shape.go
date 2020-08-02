package scad

const shapeTemplate = `module {{ .Name.String }} {
	{{ .CallName }}({{ .Fields }});
}`
