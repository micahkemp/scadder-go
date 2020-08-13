package internal

const TransformationTemplate = `{{ range .Children }}use <{{ .Name }}/{{ .Name.Filename }}>
{{ end }}
module {{ .Name.String }}() {
	{{ .CallName }}({{ .Fields }}) { {{- range .Children }}
		{{ .Name }}();
	{{- end }}
	}
}

// call module when loaded directly
{{ .Name.String }}();
`
