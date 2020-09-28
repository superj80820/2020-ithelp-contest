#! /bin/bash
docker run -v ${PWD}/../:/defs namely/gen-grpc-gateway:1.29_4 -f ./schemas/digimon/schema.proto -s Service -o ./go-server
docker run -v ${PWD}/../:/defs namely/gen-grpc-gateway:1.29_4 -f ./schemas/weather/schema.proto -s Service -o ./go-server/gen/weather