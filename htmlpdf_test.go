package htmlpdfgg

import (
	"os"
	"testing"
)

var hp *HtmlPDF

func TestGenerate(t *testing.T) {
	hp = &HtmlPDF{}

	templateString := `
    <html>
    <head>
        <title>{{ title }}</title>
    </head>
    <body>
        <h1>{{ header }}</h1>
        <p>{{ content }}</p>
    </body>
    </html>
    `

	data := map[string]interface{}{
		"title":   "htmlpdf",
		"header":  "Generate from htmlpdf",
		"content": "This is a sample PDF created with Pongo2 and chromedp.",
	}

	hp.SetTemplateData(templateString, data)
	pdfBuffer, err := hp.Generate()
	if err != nil {
		t.Error(err)
	}

	if err := os.WriteFile("output.pdf", pdfBuffer, 0644); err != nil {
		t.Error(err)
	}

	t.Logf("success")
}
