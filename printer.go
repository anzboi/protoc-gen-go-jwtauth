package main

import (
	"bytes"

	"github.com/anzx/pkg/protoc-gen-gojwtauth/jwtauthoption"
	pgs "github.com/lyft/protoc-gen-star"
)

type jwtauthModule struct {
	*pgs.ModuleBase
}

func JwtauthModule() *jwtauthModule {
	return &jwtauthModule{ModuleBase: &pgs.ModuleBase{}}
}

func (p *jwtauthModule) Name() string {
	return "jwtauth"
}

func (p *jwtauthModule) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	buf := bytes.Buffer{}
	visitor := NewJwtauthVisitor(p)
	for _, pkg := range packages {
		pgs.Walk(visitor, pkg)
	}
	p.Debug(visitor.scopes)
	p.AddGeneratorFile("name.cli.go", buf.String())
	return p.Artifacts()
}

type JwtauthVisitor struct {
	pgs.Visitor
	pgs.DebuggerCommon
	scopes map[string][]string
}

func NewJwtauthVisitor(debugger pgs.DebuggerCommon) *JwtauthVisitor {
	return &JwtauthVisitor{
		Visitor: pgs.NilVisitor(),
		scopes:  map[string][]string{},
	}
}

func (j *JwtauthVisitor) VisitPackage(pgs.Package) (v pgs.Visitor, err error) { return j, nil }
func (j *JwtauthVisitor) VisitFile(pgs.File) (v pgs.Visitor, err error)       { return j, nil }
func (j *JwtauthVisitor) VisitService(pgs.Service) (v pgs.Visitor, err error) { return j, nil }

func (j *JwtauthVisitor) VisitMethod(method pgs.Method) (v pgs.Visitor, err error) {
	scopes := []string{}
	if !method.BuildTarget() {
		return nil, nil
	}
	if ok, err := method.Extension(jwtauthoption.E_Scopes, &scopes); err != nil {
		return nil, err
	} else if ok {
		j.scopes[method.FullyQualifiedName()] = scopes
	}
	return nil, nil
}
