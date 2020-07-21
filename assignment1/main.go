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
	//i := -1
	for ;; {
		//i++
		fmt.Println("enter employee id")
		id, _ := reader.ReadString('\n')
		fmt.Println("enter employee salary")
		sal, _ := reader.ReadString('\n')

		idInt,_ := strconv.ParseInt(strings.TrimSpace(id),10,64)
		salInt,_ := strconv.ParseInt(strings.TrimSpace(sal) ,10,64)
		emp.empId = idInt
		emp.salary = salInt
		fmt.Println("Do you want to quit?")
		res,_ := reader.ReadString('\n')
		empSlice = append(empSlice, emp)
		if strings.TrimSpace(res) == "Q" {
			fmt.Println("inside if")
			break
		}
	}
	for i := 0; i < len(empSlice) ;i++ {
		fmt.Println(empSlice[i].salary,empSlice[i].empId)
	}
}

