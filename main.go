package main

import (
	"fmt"
	"math"
	"time"
)

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

type geometry interface {
	area() float64
	perim() float64
}

type rec struct {
	widht, height float64
}

type cir struct {
	radius float64
}

func variable() {
	var numberInt, numberInt_2 int = 1000, 2000
	var msg string = "Hello"

	numberfloat := 25.4
	fmt.Println(numberInt)
	fmt.Println(numberInt_2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	fmt.Println(numberInt + numberInt_2)
	fmt.Println(numberfloat + float64(numberInt_2))
	fmt.Println(msg + " World")
	fmt.Println("My money : ", numberInt_2)
}

func array_go() {
	var productName [4]string
	// var price [4]float32

	productName[0] = "Macbook"
	productName[1] = "Mac"
	productName[2] = "iPad"
	productName[3] = "iPhone"

	price := [4]float32{40000, 3000, 20000, 35000}

	fmt.Println(productName)
	fmt.Println(price)

}

func slice_go() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)

	courseName = append(courseName, "C", "C#", "c++")
	fmt.Println(courseName)

	courseWeb := courseName[2:5]
	fmt.Println(courseWeb)

	courseWeb2 := courseName[:3]
	fmt.Println(courseWeb2)
}

func map_go() {
	// use key and value
	var product = make(map[string]float32)

	// fmt.Println("procect", product)
	// add
	product["Macbook"] = 40000
	product["iPad"] = 30000
	product["iPhone"] = 20000
	fmt.Println("procect", product)

	// delate
	delete(product, "iPad")
	fmt.Println("procect", product)

	// update
	product["iPhone"] = 16000
	fmt.Println("procect", product)

	value1 := product["iPhone"]
	fmt.Println("value", value1, "baht")

	courseName := map[string]string{"101": "java", "102": "python"}
	fmt.Println(courseName)
}

func AddNumbers(x int, y int) int {
	sum := x + y
	return sum
}

func zeroVal(ival int) {
	ival = 0
}

func zeroPointer(iPointer *int) {
	*iPointer = 0
}

func pointer_go() {
	i := 1
	fmt.Println("i = ", i)

	zeroVal(i)
	fmt.Println("i fromzeroVal func = ", i)

	zeroPointer(&i)
	fmt.Println("i value zeroPointer func = ", i)
	fmt.Println("i address zeroPointer func = ", &i)
}

func condition_go() {
	myMoney := 100
	if myMoney > 100 {
		fmt.Println("Ok")
	} else {
		fmt.Println("abort")
	}
}

func loop_go() {
	// i := 0
	// const count = 15
	// for i <= 10 {
	// 	fmt.Println("i = ", i)
	// 	i++
	// }

	// for j := 0; j < count; j++ {
	// 	fmt.Println("j = ", j)
	// }

	// while loop
	var input string
	for {
		// fmt.Println("Hello AIT")
		fmt.Scanf("%s", &input)
		fmt.Println("input = ", input)
		if input == "exit" {
			break
		}
	}
}

func switch_go() {
	input := 2
	switch input {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")

	default:
		fmt.Println("Invalid value")
	}
}

func struc_go() {
	emp1 := employee{
		employeeID:   "101",
		employeeName: "Two",
		phone:        "0000000",
	}
	fmt.Println("employee 1", emp1)

	empList := [3]employee{}
	empList[0] = employee{
		employeeID:   "101",
		employeeName: "XXX",
		phone:        "000000",
	}
	empList[1] = employee{
		employeeID:   "102",
		employeeName: "YYY",
		phone:        "000000",
	}
	empList[2] = employee{
		employeeID:   "103",
		employeeName: "ZZZ",
		phone:        "000000",
	}
	for i := 0; i < 3; i++ {
		fmt.Println(empList[i])
	}

	// slince
	empListX := []employee{}
	empX := employee{
		employeeID:   "101",
		employeeName: "Slince",
		phone:        "000000",
	}
	empListX = append(empListX, empX)
	fmt.Println(empListX)
}

// Interface
func (r rec) area() float64 {
	return r.widht * r.height
}

func (r rec) perim() float64 {
	return 2*r.widht + 2*r.height
}

func (c cir) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c cir) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Aear : ", g.area())
	fmt.Println("Perimeter : ", g.perim())
}

func interface_go() {
	r := rec{widht: 5, height: 10}
	c := cir{radius: 25}

	measure(r)
	measure(c)
}

// rutine Thread

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func rutine_go() {
	go f("msg")
	go f("msg2")
	time.Sleep(5 * time.Second)
}

// chanel

func process1(c chan string, data string) {
	c <- data
}

func chanel_go() {
	ch := make(chan string)
	go process1(ch, "Data1")

	// <- resive data from chanel
	fmt.Println(<-ch)
	time.Sleep(5 * time.Second)
}

func defer_go() {
	defer fmt.Println("End process")
	for i := 0; i < 10; i++ {
		defer fmt.Println("Hello : ", i)
	}
}

func main() {
	// variable()
	// array_go()
	// slice_go()
	// map_go()
	// fmt.Println(AddNumbers(1, 2))
	// pointer_go()
	// condition_go()
	// loop_go()
	// switch_go()
	// struc_go()
	// interface_go()
	// rutine_go()
	// chanel_go()
	defer_go()
}
