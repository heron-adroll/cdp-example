#!/bin/bash


gen_and_run:
	go run github.com/99designs/gqlgen generate
	go run server.go