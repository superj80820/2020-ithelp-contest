const { QueryRequest, CreateRequest } = require("../proto/schema_pb.js");
const { DigimonPromiseClient } = require("../proto/schema_grpc_web_pb.js");

async function createDigimon(digimonPromiseClient, name) {
  let createRequest = new CreateRequest();
  createRequest.setName(name);
  const createResponse = await digimonPromiseClient.create(createRequest, {});
  return createResponse;
}

async function queryDigimonStream(digimonPromiseClient, digimonID) {
  let queryRequest = new QueryRequest();
  queryRequest.setId(digimonID);

  const queryStream = await digimonPromiseClient.queryStream(queryRequest, {});
  return queryStream;
}

(async () => {
  try {
    const digimonPromiseClient = new DigimonPromiseClient(
      "http://localhost:8080"
    );

    const createResponse = await createDigimon(digimonPromiseClient, "Agumon");
    const queryStream = await queryDigimonStream(
      digimonPromiseClient,
      createResponse.getId()
    );

    queryStream.on("data", function (response) {
      console.log(
        response.getId(),
        response.getName(),
        response.getStatus(),
        response.getLocation(),
        response.getWeather().toString()
      );
    });
    queryStream.on("status", function (status) {
      console.log(status.code);
      console.log(status.details);
      console.log(status.metadata);
    });
    queryStream.on("end", function (end) {
      // stream end signal
    });
  } catch (err) {
    console.error(err.code);
    console.error(err.message);
  }
})();
