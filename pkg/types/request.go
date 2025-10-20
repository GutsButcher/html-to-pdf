package types

type PDFRequest struct {
	Content      string `json:"content"`
	Filename     string `json:"filename,omitempty"`
	StoreInMinio bool   `json:"store_in_minio,omitempty"`
}
