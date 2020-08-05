package internal

const TransformationTemplate = `{{ range .Children }}import <{{ .Name.Filename }}>
{{ end }}
module {{ .Name.String }} {
	{{ .CallName }}({{ .Fields }}) { {{- range .Children }}
		{{ .Name }}();
	{{- end }}
	}
}`
