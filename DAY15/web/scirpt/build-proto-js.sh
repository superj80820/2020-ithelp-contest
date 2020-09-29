#! /bin/bash

docker run \
    -v "${PWD}/../schemas/digimon:/protofile" \
    -e "protofile=schema.proto" \
    juanjodiaz/grpc-web-generator

mv ../schemas/digimon/generated/* ./proto

rm -r ../schemas/digimon/generated