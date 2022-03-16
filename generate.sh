#!/usr/bin/env bash
# Shell script to autogenerate files needed for both th frontend and backend.
(cd backend;  go generate ./...)
(cd frontend; npm run --silent generate)
