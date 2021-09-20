package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              int
		mockFunc         func()
		expectedEMployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: 31,
			mockFunc: func() {
				GetEmployeeByID = func(id int) (Employee, error) {
					return Employee{ID: 1, Position: "CEO"}, nil
				}

				GetPersonByDNI = func(id int) (Person, error) {
					return Person{Name: "Andy", Age: 35, DNI: 1}, nil
				}
			},
			expectedEMployee: FullTimeEmployee{
				Employee{1, "Golang Dev"},
				Person{"Jhon", 1, 35},
			},
		},
	}

	for _, test := range table {
		test.mockFunc()
		ftemployee, err := GetFullTimeEmployeeByID(test.id, test.dni)
		if err != nil {
			t.Error("Error when getting  Employee.")
		}

		if ftemployee.Age != test.expectedEMployee.Age {
			t.Errorf("Got %d expected %d.", ftemployee.Age, test.expectedEMployee.Age)
		}
	}
}
