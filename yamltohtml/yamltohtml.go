package yamltohtml

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"gopkg.in/yaml.v2"
)

type PageData struct {
	Title string `yaml:"Title"`
	Body  string `yaml:"Body"`
}

func YamltoHTML(path string) (string, error) {
	templ, err := template.New("page").Parse(`<html><head><title>{{.Title}}</title></head><body>{{.Body}}</body></html>`)

	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	var pageData PageData
	err = yaml.Unmarshal(data, &pageData)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := templ.Execute(&tpl, pageData); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
