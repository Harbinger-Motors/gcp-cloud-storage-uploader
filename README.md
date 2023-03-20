# gcp-cloud-storage-uploader
A dead simple CLI tool to upload 

Easy to integrate into CI since it's just a static binary to download.

In order to setup authentication for a CI pipeline, define an environment variable with the base64 encoding of your service account key (ex. `GCP_AUTH_FILE_CONTENTS`). Then run:

``` bash
$ echo "$GCP_AUTH_FILE_CONTENTS" | base64 --decode > gcp-auth-file.json
$ export GOOGLE_APPLICATION_CREDENTIALS=gcp-auth-file.json
```

Now you can fetch the storage uploader: TODO