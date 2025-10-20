package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	myminio "github.com/GutsButcher/html-to-pdf/pkg/minio"
	"github.com/GutsButcher/html-to-pdf/pkg/pdf"
	"github.com/GutsButcher/html-to-pdf/pkg/types"
	"github.com/minio/minio-go/v7"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	content := "404 not found"

	replyTextContent(w, r, http.StatusOK, content)
}

func pdfGen(minioClient *minio.Client, bucketName string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PDFRequest
		if r.URL.Path != "/pdf" {
			http.NotFound(w, r)
			return
		}

		if r.Method != http.MethodPost {
			replyError(w, r, http.StatusMethodNotAllowed, "Not allowed method")
			return
		}

		// Decode JSON body
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		// Now req.Content has your HTML string
		pdfBytes, err := pdf.GeneratePDF(req.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := req.Filename
		if filename == "" {
			filename = "generated.pdf"
		}
		//

		if req.StoreInMinio {
			// create a bucket

			exists, err := minioClient.BucketExists(context.Background(), bucketName)
			if err != nil {
				http.Error(w, "Failed to check storage availability", http.StatusInternalServerError)
				log.Printf("Bucket checking error: %v", err)
				return
			}
			if !exists {
				if err := myminio.CreateBucket(bucketName, minioClient); err != nil {
					http.Error(w, "Failed to create bucket", http.StatusInternalServerError)
					log.Printf("Bucket creation: %v", err)
					return
				}
			}

			// upload to `default` bucket

			if err := myminio.UploadObject(pdfBytes, filename, bucketName, minioClient); err != nil {
				http.Error(w, "Failed to upload to bucket", http.StatusInternalServerError)
				log.Printf("Object Upload: %v", err)
				return
			}

			// test by download the object
			if err := myminio.DownloadObject(filename, bucketName, minioClient); err != nil {
				log.Printf("Object Download: %v", err)
			}
		}
		//

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		w.Write(pdfBytes)

	}

}
