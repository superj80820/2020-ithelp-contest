#! /bin/bash
docker run -v `pwd`:/defs namely/gen-grpc-gateway:1.29_4 -f schema.proto -s Service