package scad

const shapeTemplate = `module {{ .Name.String }}() {
	{{ .Function }}({{ .Fields }});
}
// call module when loaded directly
{{ .Name.String }}();
`

const transformationTemplate = `{{ range .Children }}use <{{ .SCADWriter.Name }}/{{ .SCADWriter.Name.SCADFileName }}>
{{ end }}
module {{ .Name.String }}() {
	{{ .Function }}({{ .Fields }}) { {{- range .Children }}
		{{ .SCADWriter.Name }}();
	{{- end }}
	}
}
// call module when loaded directly
{{ .Name.String }}();
`
