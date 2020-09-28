#!/bin/sh -e

psql --variable=ON_ERROR_STOP=1 --username "postgres" <<-EOSQL
    CREATE TABLE digimons (
        id uuid primary key,
        name varchar(50),
        status varchar(255)
    );

        CREATE TABLE diets (
        id uuid primary key,
        user_id uuid references digimons(id),
        name  varchar(255)
    );
EOSQL