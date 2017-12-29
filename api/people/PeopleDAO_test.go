package people

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

func TestGetPeople(t *testing.T) {
    people := GetPeople()
    if len(people) != 0 {
        t.Errorf("Expected len(GetPeople()) = %d, got: %d", 0, len(people))
	}
}

func TestGetNullPerson(t *testing.T) {
    person := GetPerson("")
    // one way to check for a non-empty Person
    if person.ID != "" {
		t.Errorf(`GetPerson("") returned ID "%s"`, person.ID)
    }
    // another way to check for a non-empty Person
	if person != (EmptyPerson) {
		t.Error(`GetPerson("") = EmptyPerson`)
	}
}

func TestGetNonExistentPerson(t *testing.T) {
	if GetPerson("z") != (EmptyPerson) {
		t.Error(`GetPerson("z") not empty`)
	}
}

func TestModifyNonExistentPerson(t *testing.T) {
	matched, people := ModifyPerson(test)
	if matched {
		t.Error(`ModifyPerson(test) matched`)
	}
    if len(people) != 0 {
        t.Errorf("Expected len(ModifyPerson(test)) = %d, got: %d", 0, len(people))
	}
}

func TestDeleteNonExistentPerson(t *testing.T) {
	matched, people := DeletePerson("z")
	if matched {
		t.Error(`DeletePerson("z") matched`)
	}
    if len(people) != 0 {
        t.Errorf(`Expected len(DeletePerson("z")) = %d, got: %d`, 0, len(people))
	}
}

func TestCreateAndGetPerson(t *testing.T) {
    defer DeletePerson("a")
    CreatePerson(nick)
	if GetPerson("a") != (nick) {
		t.Error(`GetPerson("a") != nick`)
	}
}

func TestCreateAndGetTwoPeople(t *testing.T) {
    defer DeletePerson("a")
    CreatePerson(nick)
    defer DeletePerson("b")
    CreatePerson(nora)
    people = GetPeople()
    if len(people) != 2 {
        t.Errorf("Expected len(GetPeople()) = %d, got: %d", 2, len(people))
	}
    matched, person := matchPerson("a", people)
    if !matched {
		t.Error(`GetPeople() does not contain nick`)
    }
	if person != (nick) {
		t.Error(`GetPeople() did not return nick as expected`)
	}
    matched, person = matchPerson("b", people)
    if !matched {
		t.Error(`GetPeople() does not contain nora`)
    }
	if person != (nora) {
		t.Error(`GetPeople() did not return nora as expected`)
	}
}

func TestModifyPerson(t *testing.T) {
    defer DeletePerson("b")
    CreatePerson(Person{ID: "b"})
    matched, people := ModifyPerson(nora)
    if !matched {
		t.Error(`ModifyPerson(nora) not matched`)
    }
    if len(people) != 1 {
        t.Errorf("Expected len(ModifyPerson(nora)) = %d, got: %d", 1, len(people))
	}
    matched, person := matchPerson("b", people)
    if !matched {
		t.Error(`ModifyPerson(nora) does not contain nora`)
    }
	if person != (nora) {
		t.Error(`ModifyPerson(nora) did not return nora as expected`)
	}
}

func TestDeletePerson(t *testing.T) {
    defer DeletePerson("a")
    CreatePerson(nick)
    defer DeletePerson("b")
    CreatePerson(nora)
    matched, people := DeletePerson("b")
    if !matched {
		t.Error(`DeletePerson("b") not matched`)
    }
    if len(people) != 1 {
        t.Errorf(`Expected len(DeletePerson("b")) = %d, got: %d`, 1, len(people))
	}
    matched, person := matchPerson("a", people)
    if !matched {
		t.Error(`DeletePerson("b") does not contain nick`)
    }
	if person != (nick) {
		t.Error(`DeletePerson("b") did not return nick as expected`)
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
