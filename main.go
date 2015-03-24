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
	"gopkg.in/yaml.v2"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	t_tmp "text/template"
	h_tmp "html/template"
)

const TEMPLATE_PATH = "src/github.com/cdiener/govitae/templates"
const SPACER = "    "
const WRAP_AFTER = 60

func check(e error) {
	if e != nil {	
		panic(e)
	}
}

// to_json formats a given Resume to JSON format.
// The resulting files is written to 'filename'.
func to_json(cv Resume, filename string) error {
	out, err := json.Marshal(cv)
	if err != nil {
		return err
	}
	
	err = ioutil.WriteFile(filename, out, 0666)
	
	return err
} 

// to_yaml formats a given Resume to YAML format.
// The resulting files is written to 'filename'.
func to_yaml(cv Resume, filename string) error {
	out, err := yaml.Marshal(cv)
	if err != nil {
		return err
	}
	
	err = ioutil.WriteFile(filename, out, 0666)
	
	return err
} 

// build_latex takes a Resume object and generates a nice Latex version.
// You will need moderncv installed on your Tex distribution to compile the
// resulting document. Publications are automatically extracted into a bibtex
// file and included into the cv.
func build_latex(cv Resume, name string) error {
	dir := path.Join(os.Getenv("GOPATH"), TEMPLATE_PATH)
	fs := t_tmp.FuncMap{"join": strings.Join}
	t := t_tmp.Must(t_tmp.New("template.tex").Delims("#(", ")#").
		Funcs(fs).ParseFiles(path.Join(dir, "template.tex")))

	file, err := os.Create(fmt.Sprintf("%s.tex", name))
	if err != nil {
		return err
	}
	defer file.Close()

	err = t.Execute(file, cv)
	if err != nil {
		return err
	}

	if len(cv.Publications) > 0 {
		t = t_tmp.Must(t_tmp.New("template.bib").Delims("#(", ")#").
			Funcs(fs).ParseFiles(path.Join(dir,"template.bib")))
		bib, err := os.Create("pubs.bib")
		if err != nil {
			return err
		}
		defer bib.Close()

		err = t.Execute(bib, cv)
		if err != nil {
			return err
		}
	}

	return nil
}

func wrap(s string) string {
	splitsies := strings.Fields(s)
	n := 0
	for i:=0; i<len(splitsies)-1; i++ {
		n += len(splitsies[i])
		if n>=WRAP_AFTER {
			n = 0
			splitsies[i] += "\n"
		} else { splitsies[i] += " " }
	}
	
	return strings.Join(splitsies, "")
}

func text_header(cv Resume) string {
	ascii_lines := []string{"", ""} 
	has_ascii := false
	ascii_pic, err := ioutil.ReadFile(cv.Basics.Picture+".txt")
	if err == nil {
		ascii_lines = strings.Split(string(ascii_pic), "\n")
		has_ascii = true
	}
	upper_space := int( (len(ascii_lines)-12)/2 )
	
	out := "Curriculum vitae\n=================\n\n"
	right := make([]string, 0, 10)
	for i:=0;i<upper_space;i++ {
		right = append(right,"")
	}
	right = append(right, fmt.Sprintf("%s%-12s %s %s", SPACER, "Name:", cv.Basics.First, cv.Basics.Last))
	right = append(right, fmt.Sprintf("%s%-12s %s", SPACER, "Address:", cv.Basics.Location.Address))
	right = append(right, fmt.Sprintf("%s%-12s %s %s", SPACER, "", cv.Basics.Location.PostalCode, cv.Basics.Location.City))
	right = append(right, fmt.Sprintf("%s%-12s %s", SPACER, "", cv.Basics.Location.Country))
	right = append(right, fmt.Sprintf("%s%-12s %s", SPACER, "Email:", cv.Basics.Email))
	right = append(right, fmt.Sprintf("%s%-12s %s", SPACER, "Phone:", cv.Basics.Phone))
	for _, p := range cv.Basics.Profiles {
		right = append(right, fmt.Sprintf("%s%-12s %s", SPACER, p.Network+":", p.User))
	}

	n_left := len(ascii_lines)
	n_ascii := len(ascii_lines[0])
	left := ascii_lines
	n_right := len(right)
	n := n_left
	if n_right>n_left { n = n_right }

	for i:=0; i<n; i++ {
		if i<n_left && i<n_right && has_ascii {
			out += SPACER+left[i]+SPACER+right[i]+"\n"
		} else if i<n_left && has_ascii {
			out += SPACER+left[i]+"\n"
		} else if i<n_right {
			out += fmt.Sprintf("%-[1]*[2]s", n_ascii, right[i]) + "\n"
		}
	}
	
	return out
}

// build_text generates a nice text version for the Resume object. 
// If there is a text version of your picture (picture entry + .txt) it
// will be included as ascii art in your cv. In a valid ascii art each line
// of the file *must* have the same number of characters. 
func build_text(cv Resume, name string) error {
	dir := path.Join(os.Getenv("GOPATH"), TEMPLATE_PATH)
	fs := t_tmp.FuncMap{"join": strings.Join, "wrap": wrap}
	t := t_tmp.Must(t_tmp.New("template.txt").Funcs(fs).
			ParseFiles(path.Join(dir, "template.txt")))

	file, err := os.Create(fmt.Sprintf("%s.txt", name))
	if err != nil {
		return err
	}
	defer file.Close()

	
	file.WriteString( text_header(cv) )

	err = t.Execute(file, cv)
	if err != nil {
		return err
	}
	
	return nil
}

func findpic(pic_path string) string {
	var err error
	exts := []string{".jpg",".png",".svg"}
	found := ""
	dir := path.Join(os.Getenv("GOPATH"), TEMPLATE_PATH)
	
	for _, s := range exts {
		if _, err = os.Stat(pic_path+s); err == nil {
			found = pic_path+s
		}
	}
	
	if len(found)>0 { 
		return found 
	} else { 
		png, _ := ioutil.ReadFile(path.Join(dir, "template.png"))
		ioutil.WriteFile("default.png", png, 0600)
		return "default.png"
	}
}

func build_html(cv Resume, name string) error {
	dir := path.Join(os.Getenv("GOPATH"), TEMPLATE_PATH)
	fs := h_tmp.FuncMap{"join": strings.Join, "findpic": findpic}
	t := h_tmp.Must(h_tmp.New("template.html").Funcs(fs).
			ParseFiles(path.Join(dir, "template.html")))

	file, err := os.Create(fmt.Sprintf("%s.html", name))
	defer file.Close()
	if err != nil {
		return err
	}

	err = t.Execute(file, cv)
	if err != nil {
		return err
	}
	
	return nil
}

func main() {

	if len(os.Args) < 2 {
		panic("Need a file to parse :(")
	}

	text, err := ioutil.ReadFile(os.Args[1])
	check(err)
	
	basename := strings.ToLower(path.Base(os.Args[1]))
	name := strings.TrimSuffix( basename, path.Ext(basename) )

	var cv Resume
	if x,_ := path.Match("*.j*", basename); x { 
		err = json.Unmarshal(text, &cv)
	} else { 
		err = yaml.Unmarshal(text, &cv) 
	}
	check(err)

	err = build_latex(cv, name)
	check(err)
	err = build_text(cv, name)
	check(err)
	err = build_html(cv, name)
	check(err)

	fmt.Printf("Parsed cv for %s %s from %s.\n", cv.Basics.First, cv.Basics.Last, basename)
	check(err)
}
