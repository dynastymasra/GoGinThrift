package model

import (
	"squarecode/dynastymasra/rhino/thrift/person"
	"time"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */

type Person struct {
	ID         int64   `json:"id"`
	Identity   string  `json:"identity" binding:"required`
	Firstname  string  `json:"firtsname" binding:"required`
	Middlename *string `json:"middlename"`
	Lastname   *string `json:"lastname"`
	Address    string  `json:"address" binding:"required`
	Age        *int32  `json:"age"`
	Active     bool    `json:"active" binding:"required`
	Created    string  `json:"created"`
}

func PersonResponse(person *person.Person) Person {
	return Person{
		ID:         person.ID,
		Identity:   person.Identity,
		Firstname:  person.Firstname,
		Middlename: person.Middlename,
		Lastname:   person.Lastname,
		Address:    person.Address,
		Age:        person.Age,
		Active:     person.Active,
		Created:    person.Created,
	}
}

func PersonMapping(identity, firstname string, midlename, lastname *string, address string, age *int32, active bool) *person.Person {
	return &person.Person{
		ID:         time.Now().UTC().UnixNano(),
		Identity:   identity,
		Firstname:  firstname,
		Middlename: midlename,
		Lastname:   lastname,
		Address:    address,
		Age:        age,
		Active:     active,
		Created:    time.Now().UTC().Format(time.RFC3339),
	}
}

func PersonUpdate(id int64, identity, firstname string, midlename, lastname *string, address string, age *int32, active bool) *person.Person {
	return &person.Person{
		ID:         id,
		Identity:   identity,
		Firstname:  firstname,
		Middlename: midlename,
		Lastname:   lastname,
		Address:    address,
		Age:        age,
		Active:     active,
		Created:    time.Now().UTC().Format(time.RFC3339),
	}
}
