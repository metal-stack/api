from __future__ import annotations

from dataclasses import dataclass

from connectrpc.interceptor import UnaryInterceptorSync
from google.protobuf.json_format import MessageToDict
from google.protobuf.message import Message


def _messages_equal(a: Message, b: Message) -> bool:
    return MessageToDict(a, preserving_proto_field_name=True) == MessageToDict(
        b, preserving_proto_field_name=True
    )


@dataclass
class RpcCall:
    request: Message | None = None
    response: Message | None = None
    error: Exception | None = None


class TestClientInterceptor(UnaryInterceptorSync):
    def __init__(self, calls: list[RpcCall] | None = None):
        self._calls: list[RpcCall] = list(calls) if calls else []
        self._call_count: int = 0

    def intercept_unary_sync(self, call_next, request, ctx):
        if self._call_count >= len(self._calls):
            raise AssertionError(
                f"unexpected RPC call #{self._call_count}: {request.DESCRIPTOR.full_name}"
            )

        expected = self._calls[self._call_count]
        self._call_count += 1

        if expected.request is not None and not _messages_equal(expected.request, request):
            import json

            raise AssertionError(
                f"request mismatch for {request.DESCRIPTOR.full_name}\n"
                f"  expected: {json.dumps(MessageToDict(expected.request, preserving_proto_field_name=True), indent=2)}\n"
                f"  got:      {json.dumps(MessageToDict(request, preserving_proto_field_name=True), indent=2)}"
            )

        if expected.error is not None:
            raise expected.error

        return expected.response

    def assert_all_calls_used(self):
        remaining = len(self._calls) - self._call_count
        if remaining != 0:
            raise AssertionError(
                f"{remaining} expected RPC call(s) were not made "
                f"(made {self._call_count} of {len(self._calls)})"
            )
