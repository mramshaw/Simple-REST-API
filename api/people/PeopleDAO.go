package people

var people []Person

// GetPerson returns a single specified person.
func GetPerson(id string) Person {
	for _, person := range people {
		if person.ID == id {
			return person
		}
	}
	// If no match found, return 'empty' Person
	return Person{}
}

// GetPeople returns a collection of all known persons.
func GetPeople() []Person {
	return people
}

// CreatePerson is used to create a single person.
func CreatePerson(person Person) []Person {
	people = append(people, person)
	return people
}

// ModifyPerson is used to modify a specific person.
func ModifyPerson(p Person) (bool, []Person) {
	id := p.ID
	for index, person := range people {
		if person.ID == id {
			people = append(people[:index], people[index+1:]...)
			people = append(people, p)
			return true, people
		}
	}
	return false, people
}

// DeletePerson is used to delete a specific person.
func DeletePerson(id string) (bool, []Person) {
	for index, person := range people {
		if person.ID == id {
			people = append(people[:index], people[index+1:]...)
			return true, people
		}
	}
	return false, people
}
