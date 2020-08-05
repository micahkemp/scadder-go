package internal

const ShapeTemplate = `module {{ .Name.String }} {
	{{ .CallName }}({{ .Fields }});
}`
