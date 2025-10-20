package types

type PDFResponse struct {
	Success  bool   `json:"success"`
	URL      string `json:"url,omitempty"`
	Filename string `json:"filename,omitempty"`
	Message  string `json:"message,omitempty"`
}
