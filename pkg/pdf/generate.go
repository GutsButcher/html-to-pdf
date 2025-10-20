package pdf

import (
	"fmt"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// Modified to accept HTML string and return PDF bytes
func GeneratePDF(htmlContent string) ([]byte, error) {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return nil, fmt.Errorf("NewPDFGenerator error: %s", err)
	}

	// Use the HTML content passed as parameter (not from file!)
	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlContent)))

	// Create PDF in internal buffer
	err = pdfg.Create()
	if err != nil {
		return nil, fmt.Errorf("PDF creation error: %s", err)
	}

	// Return the bytes instead of writing to file
	return pdfg.Bytes(), nil
}
