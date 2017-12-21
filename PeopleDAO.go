package main

var people []Person

func getPerson(id string) Person {
    for _, person := range people {
        if person.ID == id {
            return person
        }
    }
    // If no match found, return 'empty' Person
    return Person{}
}

func getPeople() []Person {
    return people
}

func createPerson(person Person) []Person {
    people = append(people, person)
    return people
}

func modifyPerson(p Person) (bool, []Person) {
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

func deletePerson(id string) (bool, []Person) {
    for index, person := range people {
        if person.ID == id {
            people = append(people[:index], people[index+1:]...)
            return true, people
        }
    }
    return false, people
}
