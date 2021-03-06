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
		// Test for DNI property
		if ft.DNI != test.expectedEmployee.DNI {
			t.Errorf("DNI was incorrect, got %s, expected %s", ft.DNI, test.expectedEmployee.DNI)
		}

		// Test for Name property
		if ft.Name != test.expectedEmployee.Name {
			t.Errorf("Name was incorrect, got %s, expected %s", ft.Name, test.expectedEmployee.Name)
		}

		// Test for Age Property
		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Age was incorrect, got %d, expected %d", ft.Age, test.expectedEmployee.Age)
		}

		// Test for Id Property
		if ft.Id != test.expectedEmployee.Id {
			t.Errorf("Id was incorrect, got %d, expected %d", ft.Id, test.expectedEmployee.Id)
		}

		// Test for Position Property
		if ft.Position != test.expectedEmployee.Position {
			t.Errorf("Position was incorrect, got %s, expected %s", ft.Position, test.expectedEmployee.Position)
		}
	}
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
