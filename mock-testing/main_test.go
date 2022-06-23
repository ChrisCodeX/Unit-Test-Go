package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "123456",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "Christian",
						Age:  24,
						DNI:  "123456",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					DNI:  "123456",
					Name: "Christian",
					Age:  24,
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI
	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("Error when getting Employee")
		}
	}
}
