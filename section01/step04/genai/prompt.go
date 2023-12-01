package genai

import (
	_ "embed"
	"fmt"
	"io"
	"text/template"
)

var (
	//go:embed prompt_template.txt
	promptTmplFile string
	promptTmpl     = template.Must(template.New("prompt").Parse(promptTmplFile))
)

type Instruction struct {
	Name    string
	Message string
}

type Prompt struct {
	Name     string
	Message  string
	Template *template.Template
}

func NewPrompt(name string, message string) *Prompt {
	return &Prompt{
		Name:     name,
		Message:  message,
		Template: promptTmpl,
	}
}

func (p *Prompt) Write(w io.Writer) error {
	inst := &Instruction{
		Name:    p.Name,
		Message: p.Message,
	}

	if err := p.Template.Execute(w, inst); err != nil {
		return fmt.Errorf("(*genaichat.Prompt).Write: %w", err)
	}
	return nil
}
