package service

import (
	"log"
	"squarecode/dynastymasra/rhino/thrift/person"

	"git.apache.org/thrift.git/lib/go/thrift"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */

func Create(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, variable *person.Person) (*person.Person, error) {
	socket, err := thrift.NewTSocket("localhost:4000")
	if err != nil {
		log.Println(err)
	}
	transport := transportFactory.GetTransport(socket)
	client := person.NewPersonServiceClientFactory(transport, protocolFactory)
	defer client.Transport.Close()

	if err := client.Transport.Open(); err != nil {
		log.Println(err)
	}

	result, err := client.Create(variable)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func GetAll(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory) ([]*person.Person, error) {
	socket, err := thrift.NewTSocket("localhost:4000")
	if err != nil {
		log.Println(err)
	}
	transport := transportFactory.GetTransport(socket)
	client := person.NewPersonServiceClientFactory(transport, protocolFactory)
	defer client.Transport.Close()

	if err := client.Transport.Open(); err != nil {
		log.Println(err)
	}

	result, err := client.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func Delete(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, id int64) error {
	socket, err := thrift.NewTSocket("localhost:4000")
	if err != nil {
		log.Println(err)
	}
	transport := transportFactory.GetTransport(socket)
	client := person.NewPersonServiceClientFactory(transport, protocolFactory)
	defer client.Transport.Close()

	if err := client.Transport.Open(); err != nil {
		log.Println(err)
	}

	err = client.Destroy(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Update(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, variable *person.Person) (*person.Person, error) {
	socket, err := thrift.NewTSocket("localhost:4000")
	if err != nil {
		log.Println(err)
	}
	transport := transportFactory.GetTransport(socket)
	client := person.NewPersonServiceClientFactory(transport, protocolFactory)
	defer client.Transport.Close()

	if err := client.Transport.Open(); err != nil {
		log.Println(err)
	}

	result, err := client.Update(variable)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
