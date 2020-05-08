package main

import (
	"bytes"

	"github.com/anzx/pkg/protoc-gen-go-jwtauth/jwtauthoption"
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

	data := struct {
		Package string
		Methods map[pgs.Method][]*jwtauthoption.Scopes
	}{
		Package: "pkg",
		Methods: visitor.scopes,
	}
	if err := jwtauthTemplate.Execute(&buf, data); err != nil {
		panic(err)
	}
	p.AddGeneratorFile("name.jwtauth.go", buf.String())
	return p.Artifacts()
}

type JwtauthVisitor struct {
	pgs.Visitor
	pgs.DebuggerCommon
	scopes map[pgs.Method][]*jwtauthoption.Scopes
}

func NewJwtauthVisitor(debugger pgs.DebuggerCommon) *JwtauthVisitor {
	return &JwtauthVisitor{
		Visitor: pgs.NilVisitor(),
		scopes:  map[pgs.Method][]*jwtauthoption.Scopes{},
	}
}

func (j *JwtauthVisitor) VisitPackage(pgs.Package) (v pgs.Visitor, err error) { return j, nil }
func (j *JwtauthVisitor) VisitFile(pgs.File) (v pgs.Visitor, err error)       { return j, nil }
func (j *JwtauthVisitor) VisitService(pgs.Service) (v pgs.Visitor, err error) { return j, nil }

func (j *JwtauthVisitor) VisitMethod(method pgs.Method) (v pgs.Visitor, err error) {
	scopes := []*jwtauthoption.Scopes{}
	if !method.BuildTarget() {
		return nil, nil
	}
	if ok, err := method.Extension(jwtauthoption.E_Scopes, &scopes); err != nil {
		return nil, err
	} else if ok {
		j.scopes[method] = scopes
	}
	return nil, nil
}
