import { describe, it, expect } from "bun:test";
import { create } from "@bufbuild/protobuf";
import type { DescMessage, MessageShape } from "@bufbuild/protobuf";
import {
  VersionServiceGetRequestSchema,
  VersionServiceGetResponseSchema,
  VersionService,
} from "./metalstack/api/v2/version_pb";
import type { UnaryResponse } from "@connectrpc/connect";
import { newClient } from "./client";
import { newTestInterceptor } from "./test-interceptor";

function unaryResponse<Req extends DescMessage, Res extends DescMessage>(
  schema: Req,
  message: MessageShape<Res>,
): UnaryResponse<Req, Res> {
  return {
    stream: false as const,
    message: message as any,
    header: new Headers(),
    trailer: new Headers(),
    service: undefined as any,
    method: undefined as any,
  };
}

describe("client", () => {
  it("works with test interceptor", async () => {
    const c = newClient({
      baseUrl: "http://this-is-just-for-testing",
      interceptors: [
        newTestInterceptor([
          {
            wantRequest: create(VersionServiceGetRequestSchema, {}),
            wantRequestSchema: VersionServiceGetRequestSchema,
            wantResponse: () =>
              unaryResponse(VersionServiceGetRequestSchema, {
                version: { version: "1.0", revision: "", gitSha1: "", buildDate: "" },
              }),
          },
        ]),
      ],
    });

    const v = await c.apiv2().version().get({});
    expect(v.version?.version).toBe("1.0");
  });
});
