package server

import (
  "log"
  "git.apache.org/thrift.git/lib/go/thrift"
  "squarecode/dynastymasra/thrift/person"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */

type Server struct {
  host              string
  handler           *PersonImpl
  processor         *person.PersonServiceProcessor
  transport         *thrift.TServerSocket
  transportFactory  thrift.TTransportFactory
  protocolFactory   *thrift.TBinaryProtocolFactory
  server            *thrift.TSimpleServer
}

func NewServer(host string) *Server {
  handler := NewPersonImpl()
  processor := person.NewPersonServiceProcessor(handler)
  transport, err := thrift.NewTServerSocket(host)
  if err != nil {
    panic(err)
  }

  transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
  protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
  server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
  return &Server {
    host:             host,
    handler:          handler,
    processor:        processor,
    transport:        transport,
    transportFactory: transportFactory,
    protocolFactory:  protocolFactory,
    server:           server,
  }
}

func (server *Server) Run() {
  log.Printf("Server listening on %s...\n", server.host)
  server.server.Serve()
}

func (server *Server) Stop() {
  log.Println("Stopping server...")
  server.server.Stop()
}
