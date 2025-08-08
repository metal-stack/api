#!/usr/bin/env bash

curl \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer ${API_TOKEN}" \
    -d "{\"project\":\"${PROJECT_ID}\"}" \
    "${METAL_APISERVER_URL}"/metalstack.api.v2.IPService/List \
    | jq '.ips[] | .ip, .name'