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
	//t_tmp "text/template"
	//h_tmp "html/template"
)

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
	
	fmt.Printf("Parsed cv for %s.\n",cv.Basics.Name)
}
