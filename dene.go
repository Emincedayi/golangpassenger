package main

import "fmt"

type Passenger struct {
	name        string
	age         int
	job         string
	is_married  bool
	destination string
	weight      int
	status      string
	cars        []string
}

/*type Features struct {
	Passenger

}*/

func main() {

	var pas1 Passenger
	pas1.name = "Emin"
	pas1.age = 21
	pas1.job = "Student"
	pas1.is_married = false
	pas1.destination = "Dubai"
	pas1.weight = 81
	pas1.status = "Economy"

	/*pas2 := Features{
		Passenger: Passenger{
			name: "Ahmet",
			age:  18,
			job:  "Student",
		},
		destination: "Germany",
		weight:      98,
		status:      "Bussiness",
		cars: []string{
			"BMW",
			"MERCEDES",
		},
	}*/

	var pas3 Passenger
	pas3.name = "Yusuf"
	pas3.age = 28
	pas3.job = "Employee"
	pas3.is_married = true
	pas3.destination = "U.S.A"
	pas3.weight = 89
	pas3.status = "Economy"

	//fmt.Printf("%#v\n", pas1)
	//fmt.Printf("%#v\n", pas2)
	value1 := 0
	var dest1 = [4]string{"İstanbul", "Germany", "Dubai", "London"}
	for a := 0; a < 4; a++ {
		if dest1[a] == pas1.destination {
			value1 += 5
		}
	}

	if pas1.weight > 80 {
		value1 += 7
	}
	if pas1.weight <= 80 {
		value1 += 12
	}
	var prof1 = [4]string{"Student", "Enterpreneur", "Employee", "Employer"}
	for a := 0; a < 4; a++ {
		if prof1[a] == pas1.job {
			value1 += 12
		} else {
			value1 += 1
		}
	}
	var stat1 = [2]string{"Bussiness", "First_Class"}
	for a := 0; a < 2; a++ {
		if stat1[a] == pas1.status {
			value1 += 20
		} else {
			value1 += 10
		}
	}

	value2 := 0
	var dest2 = [4]string{"İstanbul", "Germany", "Dubai", "London"}
	for a := 0; a < 4; a++ {
		if dest2[a] == pas3.destination {
			value2 += 5
		}
	}

	if pas3.weight > 80 {
		value2 += 7
	}
	if pas3.weight <= 80 {
		value2 += 12
	}
	var prof2 = [4]string{"Student", "Enterpreneur", "Employee", "Employer"}
	for a := 0; a < 4; a++ {
		if prof2[a] == pas3.job {
			value2 += 12
		} else {
			value2 += 1
		}
	}
	var stat2 = [2]string{"Bussiness", "First_Class"}
	for a := 0; a < 2; a++ {
		if stat2[a] == pas3.status {
			value2 += 20
		} else {
			value2 += 10
		}
	}
	fmt.Printf("First passengers value=%d\n", value1)
	fmt.Printf("Second passengers value=%d\n", value2)
}
