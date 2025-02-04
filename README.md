# GCP-Brute-Bucket

## Overview
GCP Brute Bucket is a tool designed to perform brute-force or crawling attacks on `storage.googleapis.com/` to identify exposed ACL buckets. The tool allows the user to provide a wordlist and returns categorized results based on access permissions.


### Install Go
1. Download Go from [https://go.dev/dl/](https://go.dev/dl/).
2. Follow installation steps for your operating system.
3. Verify installation:
   ```sh
   go version
   ```
4. Set up the Go environment (if needed):
   ```sh
   export PATH=$PATH:/usr/local/go/bin
   ```
5. Initialize the Go module:
   ```sh
   go mod init gcp-brute-bucket

### Build

```sh
make 
make run
```
