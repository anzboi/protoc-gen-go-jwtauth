package main

import (
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
)

var (
	tmplFuncs = template.FuncMap{
		"ShortMethodName": func(method pgs.Method) string {
			return method.Name().String()
		},
		"FullMethodPath": func(method pgs.Method) string {
			servicePath := strings.TrimPrefix(method.Service().FullyQualifiedName(), ".")
			methodPath := method.Name().String()
			return servicePath + "/" + methodPath
		},
		"StringArray": func(strs []string) string {
			return `[]string{"` + strings.Join(strs, `","`) + `"}`
		},
	}

	jwtauthTemplate = template.Must(template.New("protoc-gen-go-jwtauth").Funcs(tmplFuncs).Parse(`
package {{.Package}}

import (
	"github.com/anzx/pkg/jwtauth"
)

{{ range $method, $scopesArr := .Methods}}
func Validate{{ShortMethodName $method}}Scopes(claims jwtauth.Claims) bool {
	{{- range $i, $scopes := $scopesArr}}
	if claims.HasAllScopes({{StringArray $scopes.And}}) {
		return true
	}
	{{- end}}
	return false
}
{{- end}}

func ValidateScopes(claims jwtauth.Claims, methodName string) bool {
	switch methodName {
		{{- range $method, $scopesArr := .Methods}}
	case "{{FullMethodPath $method}}":
		return Validate{{ShortMethodName $method}}Scopes(claims)
		{{- end }}
	default:
		return false
	}
}
`))
)
