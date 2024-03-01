#! /usr/bin/env bash

docker rm esquilo-aniquilador-nginx-1
docker rm esquilo-aniquilador-api01-1
docker rm esquilo-aniquilador-api02-1
docker rm -v esquilo-aniquilador-db-1
docker compose up
