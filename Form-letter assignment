package main

import (
    "log"
    "os"
    "text/template"
)

func main() {
    // Define a template.
    const letter = `
Dear {{.Honorific}}{{.LastName}},
{{if .Attended}}
It was a pleasure to see you at the fundraiser.{{else}}
It is a shame you couldn't attend the fundraiser.{{end}}
{{if .Donated}}
Thank you for donating.{{end}}
{{range .NextEvents}}
Remember the upcoming events: {{.}}
{{end}}
Best wishes,
Josie
`

// Prepare some data to insert into the template.
type Recipient struct {
    Honorific, LastName string
    Attended, Donated bool
    UpcEvents []string
    }

var events = []string {"Event One", "Event Two", "Event Three"}

var recipients = []Recipient{
    {"Mrs.", "Mildred", true, true, events},
    {"Mr.", "John", true, false, events},
    {"Ms.", "Rodney", false, false, events},
    }

// STEP 1 & STEP 2
// Create a new template and parse the letter into it.
t := template.Must(template.New("letter").Parse(letter))

//STEP 3
// Execute the template for each recipient.
for _, r := range recipients {
    err := t.Execute(os.Stdout, r)
      if err != nil {
          log.Println("executing template:", err)
        }
    }
}
