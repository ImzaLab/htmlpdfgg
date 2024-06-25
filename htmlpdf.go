package htmlpdfgg

import (
	"context"
	"errors"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/flosch/pongo2"
)

var (
	ErrorTemplateEmpty = errors.New("template cant empty")
)

type HtmlPDF struct {
	templateStr string
	data        map[string]interface{}
}

func (h *HtmlPDF) Generate() ([]byte, error) {
	renderedHTML, err := h.parseTemplate()
	if err != nil {
		return nil, err
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run chromedp
	var buff []byte
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("data:text/html," + renderedHTML),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buff, _, err = page.PrintToPDF().Do(ctx)
			return err
		}),
	}); err != nil {
		return nil, err
	}

	return buff, nil
}

func (h *HtmlPDF) SetTemplateData(templateStr string, data map[string]interface{}) *HtmlPDF {
	h.templateStr = templateStr
	h.data = data
	return h
}

func (h *HtmlPDF) parseTemplate() (result string, err error) {
	if len(h.templateStr) == 0 {
		return result, ErrorTemplateEmpty
	}

	// Parse the template
	tpl, err := pongo2.FromString(h.templateStr)
	if err != nil {
		return result, err
	}

	result, err = tpl.Execute(h.data)
	if err != nil {
		return result, err
	}

	return result, nil
}
