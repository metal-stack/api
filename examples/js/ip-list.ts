import { newClient } from "../../js/client";

async function main() {
  const token = process.env["API_TOKEN"];
  const project = process.env["PROJECT_ID"];
  const baseUrl = process.env["METAL_APISERVER_URL"];

  const client = newClient({
    baseUrl: baseUrl!,
    token: token!,
  });

  const listResp = await client.apiv2().ip().list({ project });

  for (const ip of listResp.ips) {
    console.log("ip", ip);
  }
}

main();
