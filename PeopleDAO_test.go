package main

import (
    "testing"
)

var EmptyPerson = Person{}

var nick = Person{ID: "a", Firstname: "Nick", Lastname: "Charles", Address: &Address{City: "San Francisco", State: "CA"}}
var nora = Person{ID: "b", Firstname: "Nora", Lastname: "Charles", Address: &Address{City: "San Francisco", State: "CA"}}

var test = Person{ID: "z", Firstname: "Test", Lastname: "Person"}

// Unless defined as parallel (with the t.Parallel() call), tests
// run sequentially. BUT the order of execution is not guaranteed,
// so each test should be independent of the others & take care of
// its own setup and teardown.

// The init() function in RestfulGorillaMux.go sets up 4 people
// which will affect the expected results of the following tests.

func TestGetPeople(t *testing.T) {
    people := getPeople()
    if len(people) != 4 {
        t.Errorf("Expected len(getPeople()) = %d, got: %d", 4, len(people))
	}
}

func TestGetNullPerson(t *testing.T) {
    person := getPerson("")
    // one way to check for a non-empty Person
    if person.ID != "" {
		t.Errorf(`getPerson("") returned ID "%s"`, person.ID)
    }
    // another way to check for a non-empty Person
	if person != (EmptyPerson) {
		t.Error(`getPerson("") = EmptyPerson`)
	}
}

func TestGetNonExistentPerson(t *testing.T) {
	if getPerson("z") != (EmptyPerson) {
		t.Error(`getPerson("z") not empty`)
	}
}

func TestModifyNonExistentPerson(t *testing.T) {
	matched, people := modifyPerson(test)
	if matched {
		t.Error(`modifyPerson(test) matched`)
	}
    if len(people) != 4 {
        t.Errorf("Expected len(getPeople()) = %d, got: %d", 4, len(people))
	}
}

func TestDeleteNonExistentPerson(t *testing.T) {
	matched, people := deletePerson("z")
	if matched {
		t.Error(`deletePerson("z") matched`)
	}
    if len(people) != 4 {
        t.Errorf("Expected len(getPeople()) = %d, got: %d", 4, len(people))
	}
}

func TestCreateAndGetPerson(t *testing.T) {
    defer deletePerson("a")
    createPerson(nick)
	if getPerson("a") != (nick) {
		t.Error(`getPerson("a") != nick`)
	}
}

func TestCreateAndGetTwoPeople(t *testing.T) {
    defer deletePerson("a")
    createPerson(nick)
    defer deletePerson("b")
    createPerson(nora)
    people = getPeople()
    if len(people) != 6 {
        t.Errorf("Expected len(getPeople()) = %d, got: %d", 6, len(people))
	}
    matched, person := matchPerson("a", people)
    if !matched {
		t.Error(`getPeople() does not contain nick`)
    }
	if person != (nick) {
		t.Error(`getPeople() did not return nick as expected`)
	}
    matched, person = matchPerson("b", people)
    if !matched {
		t.Error(`getPeople() does not contain nora`)
    }
	if person != (nora) {
		t.Error(`getPeople() did not return nora as expected`)
	}
}

func TestModifyPerson(t *testing.T) {
    defer deletePerson("b")
    createPerson(Person{ID: "b"})
    matched, people := modifyPerson(nora)
    if !matched {
		t.Error(`modifyPerson(nora) not matched`)
    }
    if len(people) != 5 {
        t.Errorf("Expected len(getPeople()) = %d, got: %d", 5, len(people))
	}
    matched, person := matchPerson("b", people)
    if !matched {
		t.Error(`getPeople() does not contain nora`)
    }
	if person != (nora) {
		t.Error(`getPeople() did not return nora as expected`)
	}
}

func TestDeletePerson(t *testing.T) {
    defer deletePerson("a")
    createPerson(nick)
    defer deletePerson("b")
    createPerson(nora)
    matched, people := deletePerson("b")
    if !matched {
		t.Error(`deletePerson("b") not matched`)
    }
    if len(people) != 5 {
        t.Errorf("Expected len(getPeople()) = %d, got: %d", 5, len(people))
	}
    matched, person := matchPerson("a", people)
    if !matched {
		t.Error(`getPeople() does not contain nick`)
    }
	if person != (nick) {
		t.Error(`getPeople() did not return nick as expected`)
	}
}

func matchPerson(id string, people []Person) (bool, Person) {
    for _, person := range people {
        if person.ID == id {
            return true, person
        }
    }
    return false, EmptyPerson
}
