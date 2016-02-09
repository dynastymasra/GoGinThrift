# Go Microservice with Gin Framework and Apache Thrift

[![Go](https://img.shields.io/badge/Go-1.4-00E5E6.svg)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Framework-Gin-blue.svg)](https://github.com/gin-gonic/gin)
[![Thrift](https://img.shields.io/badge/Apache%20Thrift-0.9.3-yellow.svg)](https://thrift.apache.org/)
[![License](https://img.shields.io/badge/license-MIT-44897A.svg)](https://github.com/dynastymasra/GolangThrift/blob/master/LICENSE)

Golang (Go Programming Language) microservice with Apache Thrift and Gin Framework.

#### Install
* Golang (Go Programming Language) <a href="https://golang.org/" target="_blank">Golang</a>
* Apache Thrift  <a href="https://thrift.apache.org/" target="_blank">Apache Thrift</a>

#### Generate Thrift
* thrift --gen go {filename}.thrift

#### Golang workspaces
* bin/
* pkg/
* src/

#### GOPATH environment variable
* $ export GOPATH=workspaces
* $ export PATH=$PATH:$GOPATH/bin

#### Import library
* go get github.com/gin-gonic/gin
* go get git.apache.org/thrift.git/lib/go/thrift/...
* go build github.com/dynastymasra/package_name or library_name < This won't produce an output file
* go install github.com/dynastymasra/package_name or library_name
