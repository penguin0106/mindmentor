#!/bin/bash
cd ./migrations

migrate -database "postgres://postgers@localhost:5432/mindmentor?sslmode=disable"