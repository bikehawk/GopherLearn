#!/bin/bash

go run jose-verify-token.go $(go run jose-create-token.go)
