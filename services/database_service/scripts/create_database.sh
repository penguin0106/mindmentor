#!/bin/bash

createdb -U postgres mindmentor

psql -U postgres mindmentor -f ./migrations/schema.sql