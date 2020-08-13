package internal

const ShapeTemplate = `module {{ .Name.String }}() {
	{{ .CallName }}({{ .Fields }});
}

// call module when loaded directly
{{ .Name.String }}();
`
