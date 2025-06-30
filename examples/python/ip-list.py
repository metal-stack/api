#!/usr/bin/env python

from google.protobuf.duration_pb2 import Duration
import os
import sonora.client
import ip_pb2_grpc, ip_pb2

# TODO set Token in header?

baseurl = os.environ['METAL_APISERVER_URL']
token = os.environ['API_TOKEN']
project = os.environ['PROJECT_ID']

with sonora.client.insecure_web_channel(baseurl) as c:
    x = ip_pb2_grpc.IPServiceStub(c)
    d = Duration(seconds=1)

    r = x.List(ip_pb2.IPServiceListRequest(project=project))
    print(r)
