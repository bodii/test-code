package proxy

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"strings"
)

func generate(file string) (string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	data := proxyData{
		Package: f.Name.Name,
	}
	cmap := ast.NewCommentMap(fset, f, f.Comments)
	for node, group := range cmap {
		name := getProxyInterfaceName(group)
		if name == "" {
			continue
		}

		data.ProxyStructName = node.(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Name.Name

		obj := f.Scope.Lookup(name)

		t := obj.Decl.(*ast.TypeSpec).Type.(*ast.InterfaceType)

		for _, field := range t.Methods.List {
			fc := field.Type.(*ast.FuncType)

			method := &proxyMethod{
				Name: field.Names[0].Name,
			}
			method.Params, method.ParamNames = getParamsOrResults(fc.Params)
			method.Results, method.ResultNames = getParamsOrResults(fc.Results)

			data.Methods = append(data.Methods, method)
		}
	}

	tpl, err := template.New("").Parse(proxyTpl)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, data); err != nil {
		return "", err
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return "", err
	}

	return string(src), nil

}

const proxyTpl = `
package {{.Package}}

type {{ .ProxyStructName }}Proxy struct {
	child *{{ .ProxyStructName }}
}

func New{{ .ProxyStructName }}Proxy(child *{{ .ProxyStructName }}) *{{ .ProxyStructName }}Proxy {
	return &{{ .ProxyStructName }}Proxy{child: child}
}

{{ range .Methods }}
func (p *{{$.ProxyStructName}}Proxy) {{ .Name }} ({{ .Params }}) ({{ .Results }}) {
	start := time.Now()

	{{ .ResultNames }} = p.child.{{ .Name }}({{ .ParamNames }})

	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return {{ .ResultNames }}
}
{{ end }}
`

type proxyData struct {
	Package         string
	ProxyStructName string
	Methods         []*proxyMethod
}

type proxyMethod struct {
	Name        string
	Params      string
	ParamNames  string
	Results     string
	ResultNames string
}

func getParamsOrResults(fields *ast.FieldList) (string, string) {
	var (
		params     []string
		paramNames []string
	)

	for i, param := range fields.List {
		var names []string
		for _, name := range param.Names {
			names = append(names, name.Name)
		}

		if len(names) == 0 {
			names = append(names, fmt.Sprintf("r%d", i))
		}

		paramNames = append(paramNames, names...)

		param := fmt.Sprintf("%s %s", strings.Join(names, ","), param.Type.(*ast.Ident).Name)
		params = append(params, strings.TrimSpace(param))
	}

	return strings.Join(params, ","), strings.Join(paramNames, ",")
}

func getProxyInterfaceName(groups []*ast.CommentGroup) string {
	for _, commentGroup := range groups {
		for _, comment := range commentGroup.List {
			if strings.Contains(comment.Text, "@proxy") {
				interfaceName := strings.TrimLeft(comment.Text, "// @proxy ")
				return strings.TrimSpace(interfaceName)
			}
		}
	}

	return ""
}
