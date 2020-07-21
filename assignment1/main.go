package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Animal interface {
	Eat() string
	Sleep() string
	Breathe() string
}

type Emp struct {
	empId int64
	salary int64
	empName string
	address Address
}

type Address struct {
	houseNumber int64
	street string
	city string
	pin int64
}

func (emp *Emp) Eat() string {
	return "This is Eat method"
}

func (emp *Emp) Sleep() string {
	return "This is Sleep Method"
}

func (emp *Emp) Breathe() string {
	return "This is breathe method"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	empSlice := []Emp{}
	var emp Emp
	var address Address
	for ;; {
		fmt.Println("enter employee id")
		id, _ := reader.ReadString('\n')

		fmt.Println("enter employee salary")
		sal, _ := reader.ReadString('\n')

		fmt.Println("enter employee name")
		name, _ := reader.ReadString('\n')

		fmt.Println("enter employee address")
		add, _ := reader.ReadString('\n')

		addressSlice := strings.Split(add, " ")
		houseNumber,_ :=  strconv.ParseInt(strings.TrimSpace(addressSlice[0]),10,64)
		pin,_ := strconv.ParseInt(strings.TrimSpace(addressSlice[3]),10,64)

		address.houseNumber = houseNumber
		address.street = addressSlice[1]
		address.city = addressSlice[2]
		address.pin = pin


		idInt,_ := strconv.ParseInt(strings.TrimSpace(id),10,64)
		salInt,_ := strconv.ParseInt(strings.TrimSpace(sal) ,10,64)

		emp.empId = idInt
		emp.salary = salInt
		emp.empName = name
		emp.address = address

		fmt.Println("Do you want to quit?")
		res,_ := reader.ReadString('\n')
		empSlice = append(empSlice, emp)

		if strings.TrimSpace(res) == "Q" {
			fmt.Println("inside if")
			break
		}
	}
	for i := 0; i < len(empSlice) ;i++ {
		fmt.Println(empSlice[i].salary,empSlice[i].empId, empSlice[i].empName,empSlice[i].address)
	}
}

