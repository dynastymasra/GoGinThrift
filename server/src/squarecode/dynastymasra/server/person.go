package server

import (
  "log"
  "fmt"
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

type PersonImpl struct {
  persons map[int64]*person.Person
}

func NewPersonImpl() *PersonImpl {
  return &PersonImpl {
    persons: make(map[int64]*person.Person),
  }
}

func (personImpl *PersonImpl) Create(person *person.Person) (*person.Person, error) {
  log.Println("Create new person...")
  personImpl.persons[person.ID] = person
  return person, nil
}

func (personImpl *PersonImpl) Read(id int64) (*person.Person, error) {
  log.Printf("Read person ID %s...\n", id)
  person, err := personImpl.persons[id]
  if !err {
    log.Printf("Person ID %s does not exist...", id)
    return nil, fmt.Errorf("Person ID %s does not exist", id)
  }
  return person, nil
}

func (personImpl *PersonImpl) Update(person *person.Person) (*person.Person, error) {
  log.Println("Update person...")
  personImpl.persons[person.ID] = person
  return person, nil
}

func (personImpl *PersonImpl) Destroy(id int64) error {
  log.Printf("Delete person ID %s...\n", id)
  if _, ok := personImpl.persons[id]; ok {
    delete(personImpl.persons, id)
  }
  return nil
}

func (personImpl *PersonImpl) GetAll() ([]*person.Person, error) {
  log.Println("Get all person...")
  var persons []*person.Person
  for _, person := range personImpl.persons {
    persons = append(persons, person)
  }
  return persons, nil
}
