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
Golang:
go run go/ip-list.go

Ansible:
ansible-playbook ansible/ip-list.yaml

Curl:
sh curl/ip-list.sh

Python:
python python/ip-list.py
```

## Python support

In order to make python work with a connect server we can choose which of the unofficial connectrpc python libraries we would use:

- [connect-python](https://github.com/mattrobenolt/connect-python/tree/ref)
- [connecpy](https://github.com/i2y/connecpy)
- [sonora-connect](https://github.com/dsludwig/sonora-connect)

All of them do not have a buf plugin yet, the last one does not need one.

There is a upstream PR where this feature request is discussed: [Connect Python Implementation](https://github.com/connectrpc/connectrpc.com/pull/71)