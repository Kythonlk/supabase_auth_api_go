#!/bin/bash

cd cmd

go build -tags netgo -ldflags '-s -w' -o app

chmod +x ./app

./app
