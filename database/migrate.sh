#!/bin/bash

migrate -path migrations -database "postgres://user:password@localhost:5432/database?sslmode=disable" up
