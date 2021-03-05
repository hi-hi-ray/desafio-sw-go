#!/bin/sh


echo "Reseting the GO111MODULE to auto"

export GO111MODULE="auto"

echo "Installing required packages"

go get -v github.com/BurntSushi/toml
go get -v github.com/codegangsta/inject
go get -v github.com/go-martini/martini
go get -v github.com/stretchr/testify
go get -v go.mongodb.org/mongo-driver

echo "If everything it's okay, please make sure that your mongo server it's up. Don't forget to check the README to see more instructions"