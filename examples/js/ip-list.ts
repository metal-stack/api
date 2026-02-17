import * as apiv2 from "@metal-stack/api/js/metalstack/api/v2/ip_pb";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

import { Code, ConnectError, Interceptor } from "@connectrpc/connect";

class AuthInterceptor {
  private authToken: string;

  constructor(authToken: string) {
    this.authToken = authToken;
  }

  interceptor: Interceptor = (next) => async (req) => {
    if (!this.authToken) {
      throw new ConnectError("Missing auth token", Code.Unauthenticated);
    }

    req.header.append("Authorization", `Bearer ${this.authToken}`);

    try {
      const res = await next(req);
      return res;
    } catch (e) {
      if (e instanceof ConnectError && e.code === Code.Unauthenticated) {
        // e.g. message: "token has expired"
        console.error("unauthenticated", e);
      }
      throw e;
    }
  };
}

async function main() {
  const token = process.env["API_TOKEN"];
  const project = process.env["PROJECT_ID"];
  const baseUrl = process.env["METAL_APISERVER_URL"];

  const auth = new AuthInterceptor(token!);
  const client = createClient(
    apiv2.IPService,
    createConnectTransport({
      baseUrl: baseUrl!,
      interceptors: [auth.interceptor],
    }),
  );

  const listResp = await client.list({ project });

  for (const ip of listResp.ips) {
    console.log("ip", ip);
  }
}

main();
