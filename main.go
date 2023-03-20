package main

import (
	"context"
	b64 "encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

func main() {
	bucketFlag := flag.String("bucket", "", "the bucket to upload to")
	remotePathFlag := flag.String("bucket-path", "", "the path in the bucket to upload to")
	useEnvCredsFlag := flag.Bool("use-env-creds", false, "uses the base64 decoded value of the GCP_AUTH_FILE_CONTENTS variable as the Google Cloud .json service file credentials")

	flag.Parse()

	args := flag.Args()

	if *useEnvCredsFlag {
		base64Encoded := os.Getenv("GCP_AUTH_FILE_CONTENTS")

		decoded, err := b64.StdEncoding.DecodeString(base64Encoded)
		if err != nil {
			fmt.Printf("Error decoding GCP_AUTH_FILE_CONTENTS as base64: %v\n", err)
			os.Exit(1)
		}

		err = os.WriteFile("gcp-auth-file.json", decoded, 0644)
		if err != nil {
			fmt.Printf("Error writing to gcp-auth-file: %v\n", err)
			os.Exit(1)
		}

		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "gcp-auth-file.json")
	}

	if len(args) == 0 {
		fmt.Printf("Please specify the path to the file to upload.\n")
		os.Exit(1)
	}

	fmt.Printf("Uploading file %v to bucket %v at path %v...\n", args[0], *bucketFlag, *remotePathFlag)

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("storage.NewClient: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	f, err := os.Open(args[0])
	if err != nil {
		fmt.Printf("Could not open file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	o := client.Bucket(*bucketFlag).Object(*remotePathFlag)

	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		fmt.Printf("io.Copy: %v\n", err)
		os.Exit(1)
	}

	if err := wc.Close(); err != nil {
		fmt.Printf("Writer.Close: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Blob %v uploaded.\n", *remotePathFlag)
}
