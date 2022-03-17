#!/usr/bin/env bash
# Shell script to autogenerate files needed for both th frontend and backend.
(cd openapi;  widdershins foodtinder.json -o README.md --omitHeader -c &> /dev/null)
(cd backend;  go generate ./...)
(cd frontend; npm run --silent generate)
