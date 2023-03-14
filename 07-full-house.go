package main

import "fmt"

type HousePerson struct {
	Name string
}

type House struct {
	People map[string]*HousePerson
}

func (h *House) addPeople(people []HousePerson) {
	for _, person := range people {
		h.People[person.Name] = &person
	}
}

func FullHouse() {
	h := House{
		People: make(map[string]*HousePerson),
	}
	h.addPeople([]HousePerson{
		{Name: "Alice"},
		{Name: "Bob"},
	})

	for _, person := range h.People {
		fmt.Println(person)
	}
}

func (h *House) addPeopleAlt(people []HousePerson) {
	for i := range people {
		h.People[people[i].Name] = &people[i]
	}
}

func (h *House) addTwoPeople(people []HousePerson) {
	var person HousePerson
	// &person is (for example) 0x0000000001

	person = people[0]              // value of 'person' changes, not it's memory address
	h.People[person.Name] = &person // &person is still 0x0000000001

	person = people[1]              // again, value changes, but not the memory address
	h.People[person.Name] = &person // again, &person is still 0x0000000001

	// at the end we have two people entries that are the same pointer
}

/*
 * For-range loop uses a single variable that is populated with different value for each iteration. This means that it
 * has the same memory address. Therefore, even if we change value, variable's address stays the same.
 */
