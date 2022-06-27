// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 115.

// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"html/template"
)

//!+template

var itemslist = template.Must(template.New("itemslist").Parse(`

<h1>Gerei a partir daqui<h1>
{{range .Items}}
</tr>

<tr class="item">
<td>{{.Name}} </td>

<td>{{.Value}}</td>

</tr>
{{end}}

<tr class="total">
<td></td>

<td>Total: R$300.00</td>
</tr>
</table>
</div>
</body>
</html>
`))

type Item struct {
	Name  string
	Value string
}
type ItemsResult struct {
	TotalCount int `json:"total_count"`
	Items      []Item
}

func main() {

	var result = []Item{
		{
			Name:  "Espelho",
			Value: "300",
		},
		{
			Name:  "Vidro",
			Value: "400",
		},
	}

	final := ItemsResult{

		TotalCount: 2,
		Items:      result,
	}

	generateTemplate()

	file, err := os.OpenFile("invoice_2022.html", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	fmt.Println("File created successfully")
	defer file.Close()
	if err := itemslist.Execute(file, final); err != nil {
		log.Fatal(err)
	}
}

func generateTemplate() {
	fin, err := os.Open("template.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create("invoice_2022.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	if err != nil {
		log.Fatal(err)
	}
}

//!-
