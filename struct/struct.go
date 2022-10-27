package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "Pradoo",
		phone:        "0901234001",
	}
	employee2 := employee{
		employeeID:   "102",
		employeeName: "Prayad",
		phone:        "0901234002",
	}
	employee3 := employee{
		employeeID:   "103",
		employeeName: "Pranee",
		phone:        "0901234003",
	}

	employeeList = append(employeeList, employee1, employee2, employee3)

	fmt.Println("employ =", employeeList)
}
