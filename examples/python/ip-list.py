#!/usr/bin/env python

# client.py
from connecpy.context import ClientContext
from connecpy.exceptions import ConnecpyServerException
import os
# import sys

from metalstack.api.v2 import ip_pb2, ip_connecpy

timeout_s = 5
baseurl = os.environ['METAL_APISERVER_URL']
token = os.environ['API_TOKEN']
project = os.environ['PROJECT_ID']

def main():
    with ip_connecpy.IPServiceClient(baseurl, timeout=timeout_s) as client:
        try:
            response = client.List(
                ctx=ClientContext(),
                request=ip_pb2.IPServiceListRequest(project=project),
                headers={
                    "Authorization": "Bearer " + token,
                }
            )
            for ip in response.ips:
                print(ip.ip, ip.name, ip.project, ip.network)
        except ConnecpyServerException as e:
            print(e.code, e.message, e.to_dict())


if __name__ == "__main__":
    main()
