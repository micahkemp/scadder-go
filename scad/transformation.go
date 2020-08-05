package scad

const transformationTemplate = `{{ range .Children }}import <{{ .Name.Filename }}>
{{ end }}
module {{ .Name.String }} {
	{{ .CallName }}({{ .Fields }}) { {{- range .Children }}
		{{ .Name }}();
	{{- end }}
	}
}`

func (t Template) Translate(x, y, z float64, n string) Template {
	return NewTranslate(x, y, z, n, t)
}
