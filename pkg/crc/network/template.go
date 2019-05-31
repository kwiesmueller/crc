package network

import (
	"bytes"
	"text/template"
)

const (
	resolvFileTemplate = `# Generated by CRC
search {{ range .SearchDomains }}{{ .Domain }}{{ end }}
{{ range .NameServers }}nameserver {{ .IPAddress }}
{{ end }}
`
)

func createResolvFile(values ResolvFileValues) (string, error) {
	var resolvFile bytes.Buffer

	t, err := template.New("resolvfile").Parse(resolvFileTemplate)
	if err != nil {
		return "", err
	}
	t.Execute(&resolvFile, values)

	return resolvFile.String(), nil
}
