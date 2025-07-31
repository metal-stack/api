#!/usr/bin/env python

from connecpy.context import ClientContext
from connecpy.exceptions import ConnecpyServerException
import os

from metalstack.api.v2 import network_pb2, ip_pb2, ip_connecpy
from client import client as cli


timeout_s = 5
baseurl = os.environ['METAL_APISERVER_URL']
token = os.environ['API_TOKEN']
project = os.environ['PROJECT_ID']

def main():
    # Usage Sample without driver
    print("ips with direct client")
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


    # Usage Sample with driver
    print()
    print("ips with driver")
    driver = cli.Driver(baseurl=baseurl, token=token, timeout=timeout_s)
    with driver.api().ip() as client:
        try:
            response = client.List(
                request=ip_pb2.IPServiceListRequest(project=project),
            )
            for ip in response.ips:
                print(ip.ip, ip.name, ip.project, ip.network)
        except ConnecpyServerException as e:
            print(e.code, e.message, e.to_dict())

    print()
    print("networks with admin driver")
    with driver.admin().network() as client:
        try:
            response = client.List(
                request=network_pb2.NetworkServiceListRequest(),
            )
            for nw in response.networks:
                print(nw.id, nw.name, nw.project)
        except ConnecpyServerException as e:
            print(e.code, e.message, e.to_dict())
if __name__ == "__main__":
    main()
