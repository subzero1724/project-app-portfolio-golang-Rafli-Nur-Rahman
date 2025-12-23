# Portfolio Templates

These HTML templates are designed to work with Go's `html/template` package and Bootstrap 5.

## Templates Available:

1. **layout.html** - Base layout with navbar and footer
2. **index.html** - Homepage with hero, projects, skills, experience, and contact
3. **projects.html** - Full projects listing page
4. **contact.html** - Dedicated contact page
5. **admin.html** - Admin dashboard for content management

## Features:

- Modern, gradient-based design
- Fully responsive (mobile-first)
- Bootstrap 5.3.2
- Font Awesome 6.5.1 icons
- Google Fonts (Inter & Space Grotesk)
- Smooth animations and transitions
- Clean, professional aesthetic

## Usage in Go:

```go
tmpl := template.Must(template.ParseFiles(
    "templates/layout.html",
    "templates/index.html",
))

data := struct {
    Title      string
    Projects   []Project
    Skills     []Skill
    Experience []Experience
}{
    Title: "Home",
    // ... your data
}

tmpl.ExecuteTemplate(w, "layout", data)
