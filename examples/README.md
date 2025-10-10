# Usage examples

You must first get a API Token and the Project ID to access the metalstack.cloud api then export this as environment variables.

Grab a project with metalctl.

```bash
export API_TOKEN=yourapitoken
export PROJECT_ID=yourprojectuuid
export METAL_APISERVER_URL=metal-apiserver url
```

Then you can execute them as following:

```bash
# Golang:
go run go/ip-list.go

# Curl:
sh curl/ip-list.sh

# Python:
python python/ip-list.py
```
