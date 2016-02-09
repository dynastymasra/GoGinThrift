typedef string timestamp

struct Person {
  1: required i64 ID,
  2: required string Identity,
  3: required string Firstname,
  4: optional string Middlename,
  5: optional string Lastname,
  6: required string Address,
  7: optional i32 Age,
  8: required bool Active,
  9: required string Created
}

service PersonService {
  Person Create(1: Person person),
  Person Read(1: i64 id),
  Person Update(1: Person person),
  void Destroy(1: i64 id),
  list<Person> GetAll()
}
