# gcp-cloud-storage-uploader

A dead simple CLI tool to upload to Google Cloud Storage, that is easy to integrate into CI.

In order to setup authentication for a CI pipeline, define an environment variable with the base64 encoding of your service account key (ex. `GCP_AUTH_FILE_CONTENTS`).

Now you can fetch the storage uploader, and run it:

```bash
$ curl -L -O https://github.com/Harbinger-Motors/gcp-cloud-storage-uploader/releases/download/1.0.0/gcp-cloud-storage-uploader-linux-x86_64
$ chmod +x ./gcp-cloud-storage-uploader-linux-x86_64
$ ./gcp-cloud-storage-uploader-linux-x86_64 --use-env-creds --bucket=bucket-name --bucket-path=path/to/destination path/to/file
```

Removing the `--use-env-creds` flag will use the default credentials installed on the system (ex. development environment).