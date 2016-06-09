package main

import (
	"testing"
	"encoding/json"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func BenchmarkLatexJSON(b *testing.B) {
	var cv Resume
	tp := "templates"
    
	for i:=0; i<b.N; i++ {
		text, _ := ioutil.ReadFile("examples/resume.json")
		json.Unmarshal(text, &cv)
		build_latex(cv, "test", tp)
	}
}

func BenchmarkLatexYAML(b *testing.B) {
	var cv Resume
	tp := "templates"
    
	for i:=0; i<b.N; i++ {
		text, _ := ioutil.ReadFile("examples/resume.yaml")
		yaml.Unmarshal(text, &cv)
		build_latex(cv, "test", tp)
	}
}
