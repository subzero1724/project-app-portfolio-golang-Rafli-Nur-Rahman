package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"time"

	"portfolio/internal/dto"
	"portfolio/internal/model"
)

func main() {
	projects := []*dto.ProjectResponse{{ID: "1", Title: "P1", Description: "d", TechStack: []string{"Go"}, ImageURL: "", DemoURL: "", GithubURL: ""}}
	skills := []*model.Skill{{ID: "s1", Name: "Go", Level: 90}}
	exps := []*model.Experience{{ID: "e1", Company: "C", Position: "P", Description: "D", StartDate: time.Now(), IsCurrent: false}}
	msgs := []*model.Contact{{ID: "m1", Name: "Foo", Email: "f@b", Subject: "Hi", Message: "msg"}}

	data := map[string]interface{}{
		"Title":      "Admin",
		"Projects":   projects,
		"Skills":     skills,
		"Experience": exps,
		"Messages":   msgs,
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/admin.html")
	if err != nil {
		fmt.Println("parse error:", err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "layout", data); err != nil {
		fmt.Println("exec error:", err)
		os.Exit(2)
	}

	fmt.Println("OK: rendered", len(buf.Bytes()), "bytes")
}
