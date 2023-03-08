import net from "net";
import { randomUUID } from "crypto";

const client = net.connect(3001, "127.0.0.1", () => {
  console.log("connected to server!");
});

client.write(
  JSON.stringify({
    Type: "write",
    Data: {
      key: "key",
      value: "value",
      commend: "SET",
    },
    id: randomUUID(),
  })
);
client.write(
  JSON.stringify({
    Type: "write",
    Data: {
      key: "key",
      commend: "GET",
    },
    id: randomUUID(),
  })
);

client.on("data", (data) => {
  // client.write("data " + data);
  console.log(data.toString(), "data");
});
