/*
 * main.go
 * 
 * Copyright 2015 Christian Diener <ch.diener@gmail.com>
 * 
 * MIT license. See LICENSE for more information.
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	t_tmp "text/template"
	//h_tmp "html/template"
)

func build_latex(cv Resume) error {
	t := t_tmp.Must(t_tmp.New("template.tex").Delims("#(",")#").ParseFiles("template.tex"))
	name := fmt.Sprintf("%s_%s", cv.Basics.First, cv.Basics.Last)
	
	file, err := os.Create( fmt.Sprintf("%s.tex", name) )
	if err!=nil { return err } 
	defer file.Close()
	
	err = t.Execute(file, cv)
	if err!=nil { return err }
	
	return nil
}

func main() {
	
	if len(os.Args)<2 {
		panic("Need a json file to parse :(")
	}
	
	text, err := ioutil.ReadFile(os.Args[1])
	if err!=nil {
		panic(err)
	}
	
	var cv Resume
	err = json.Unmarshal(text, &cv)
	if err != nil {
		panic(err)
	}
	
	err = build_latex(cv)
	if err!=nil { panic(err) }
	
	fmt.Printf("Parsed cv for %s %s.\n",cv.Basics.First, cv.Basics.Last)
}
