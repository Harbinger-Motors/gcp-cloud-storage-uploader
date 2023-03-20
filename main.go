package main

import (
	"context"
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

	flag.Parse()

	args := flag.Args()

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
