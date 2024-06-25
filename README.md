# HTMLPDFgg
## ease of converting html to pdf

### Get Started
#### Basic usage

`hpg.Generate()` will write to a buffer `pdfBuffer`
```
pdfBuffer, err := hpg.Generate()
if err != nil {
 panic(err)
}

```

\
&nbsp;
\
&nbsp;


write to file from buffer, see on `example/write_to_file.go`
```
templateString := `
    <html>
    <head>
        <title>{{ title }}</title>
    </head>
    <body>
        <h1>{{ header }}</h1>
        <p>{{ content }}</p>
		<table border="1">
            <tr>
                <th>Name</th>
                <th>Age</th>
            </tr>
            {% for order in orders %}
            <tr>
                <td>{{ order.id }}</td>
                <td>
				 {% for item in order.items %}
				 	{{ item.name }}<br/>
				 {% endfor %}
				 </td>
            </tr>
            {% endfor %}
        </table>
    </body>
    </html>
    `

	orderItems := []map[string]interface{}{
		{"id": "123", "name": "item1"},
		{"id": "124", "name": "item2"},
		{"id": "125", "name": "item3"},
	}
	orders := []map[string]interface{}{
		{"id": "123", "items": orderItems},
		{"id": "124", "items": orderItems},
		{"id": "125", "items": orderItems},
	}

	data := map[string]interface{}{
		"title":   "htmlpdf",
		"header":  "Generate from htmlpdf",
		"content": "This is a sample PDF created with Pongo2 and chromedp.",
		"orders":  orders,
	}

	hpg := htmlpdfgg.HtmlPDF{}
	hpg.SetTemplateData(templateString, data)
	pdfBuffer, err := hpg.Generate()
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("output.pdf", pdfBuffer, 0644); err != nil {
		panic(err)
	}
```

### TODO:
- supported with watermark