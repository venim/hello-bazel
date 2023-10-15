import { createPromiseClient } from "@connectrpc/connect";
import { createGrpcWebTransport } from "@connectrpc/connect-web";
import { HelloService } from '@hello-bazel/proto/hello_connect.js';
import { HelloRequest } from '@hello-bazel/proto/hello_pb.js';

const transport = createGrpcWebTransport({
  baseUrl: "http://localhost:8080",
});

const client = createPromiseClient(HelloService, transport);

let req = new HelloRequest({ name: "World!" });

console.log(`sending ${req.toJsonString()}`);

client.sayHello(req).then((resp) => {
  console.log(resp);
  console.log(resp.message);
});