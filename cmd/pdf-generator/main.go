package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	myminio "github.com/GutsButcher/html-to-pdf/pkg/minio"
	"github.com/GutsButcher/html-to-pdf/pkg/server"
	minio "github.com/minio/minio-go/v7"
)

const bucketName = "default"

type Server struct {
	MinIOClient *minio.Client
}

func main() {
	host := flag.String("h", "localhost", "Server host")
	port := flag.Int("p", 8080, "Server port")
	flag.Parse()
	minioClient, err := myminio.InitMinIOClient(
		os.Getenv("MINIO_ENDPOINT"),
		os.Getenv("MINIO_ROOT_USER"),
		os.Getenv("MINIO_ROOT_PASSWORD"),
		false)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	clients := Server{
		MinIOClient: minioClient,
	}
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", *host, *port),
		Handler:      server.NewMux(clients.MinIOClient, bucketName),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

}
