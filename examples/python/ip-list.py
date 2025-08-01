#!/usr/bin/env python

import os
import sys

from connecpy.exceptions import ConnecpyServerException

from metalstack.client import client as apiclient
from metalstack.api.v2 import ip_pb2
from metalstack.admin.v2 import network_pb2

timeout_s = 5
baseurl = os.environ['METAL_APISERVER_URL']
token = os.environ['API_TOKEN']
project = os.environ['PROJECT_ID']

client = apiclient.Client(baseurl=baseurl, token=token, timeout=timeout_s)

try:
    resp = client.apiv2().ip().List(request=ip_pb2.IPServiceListRequest(
        project=project))
except ConnecpyServerException as e:
    print(e.code, e.message, e.to_dict())
    sys.exit(1)


for ip in resp.ips:
    print(ip.ip, ip.name, ip.project, ip.network)

resp = client.adminv2().network().List(
    request=network_pb2.NetworkServiceListRequest())

for nw in resp.networks:
    print(nw.id, nw.name, nw.project)
