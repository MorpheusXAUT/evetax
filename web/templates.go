package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

// Templates stores the parsed HTTP templates used by the web app
type Templates struct {
	template *template.Template
}

// SetupTemplates parses the HTTP templates from disk and stores them for later usage
func SetupTemplates() *Templates {
	templates := &Templates{}

	templates.template = template.Must(template.New("").Funcs(templates.TemplateFunctions(nil)).ParseGlob("app/templates/*"))

	return templates
}

// ReloadTemplates re-reads the HTTP templates from disk and refreshes the output
func (templates *Templates) ReloadTemplates() {
	templates.template = template.Must(template.New("").Funcs(templates.TemplateFunctions(nil)).ParseGlob("app/templates/*"))
}

// ExecuteTemplates performs all replacement in the HTTP templates and sends the response to the client
func (templates *Templates) ExecuteTemplates(w http.ResponseWriter, r *http.Request, template string, response map[string]interface{}) error {
	return templates.template.Funcs(templates.TemplateFunctions(r)).ExecuteTemplate(w, template, response)
}

// TemplateFunctions prepares a map of functions to be used within templates
func (templates *Templates) TemplateFunctions(r *http.Request) template.FuncMap {
	return template.FuncMap{
		"IsResultNil":     func(r interface{}) bool { return templates.IsResultNil(r) },
		"FormatISK":       func(i int) string { return templates.FormatISK(i) },
		"FormatFloat":     func(f float64) string { return templates.FormatFloat(f) },
		"FormatTimestamp": func(t time.Time) string { return templates.FormatTimestamp(t) },
	}
}

// IsResultNil checks whether the given result/interface is nil
func (templates *Templates) IsResultNil(r interface{}) bool {
	return (r == nil)
}

// FormatISK formats the given integer ISK value into a formatted string
func (templates *Templates) FormatISK(i int) string {
	floatValue := float64(i) / 100.0

	return fmt.Sprintf("%s", templates.FormatFloat(floatValue))
}

// FormatFloat formats a given floating point number to a human readable string
func (templates *Templates) FormatFloat(f float64) string {
	fString := humanize.Ftoa(f)

	var formattedFloat string

	if strings.Contains(fString, ".") {
		fInt, err := strconv.ParseInt(fString[:strings.Index(fString, ".")], 10, 64)
		if err != nil {
			return fString
		}

		digitsAfterPoint := len(fString) - strings.Index(fString, ".")
		if digitsAfterPoint > 3 {
			digitsAfterPoint = 3
		}

		formattedFloat = fmt.Sprintf("%s%s", humanize.Comma(fInt), fString[strings.Index(fString, "."):strings.Index(fString, ".")+digitsAfterPoint])
	} else {
		fInt, err := strconv.ParseInt(fString, 10, 64)
		if err != nil {
			return fString
		}

		formattedFloat = fmt.Sprintf("%s.00", humanize.Comma(fInt))
	}

	return formattedFloat
}

// FormatTimestamp formats the given time.Time into a shorter timestamp
func (templates *Templates) FormatTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}
