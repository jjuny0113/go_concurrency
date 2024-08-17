package advanced

import "fmt"

type Student struct {
	name string
	age  int
}

func Pointer() {
	var s Student
	var p *Student

	p = &s
	p.name = "bbb"

	fmt.Println(p)
	fmt.Println(s)
}

func ArrAndSlicePointer() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	var b [5]int
	b = array

	var c []int
	c = array[:]

	b[0] = 1000
	c[3] = 500

	fmt.Println("array:", array)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

}

func callbyCopy(n int, b [5]int, s []int) {
	n = 3000
	b[0] = 1000
	s[3] = 500
}

func CallByCopy() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	var c []int
	c = array[:]
	callbyCopy(100, array, c)

	fmt.Println("array:", array)
	fmt.Println("c:", c)
}

type Temperature struct {
	Value int
	Type  string
}

func newTemperature(v int, t string) Temperature {
	return Temperature{Value: v, Type: t}
}

func (t Temperature) String(v int) Temperature {
	return Temperature{Value: t.Value + v, Type: t.Type}
}

type StudentV2 struct {
	Age  int
	Name string
}

func newStudent(name string, age int) *StudentV2 {
	return &StudentV2{Age: age, Name: name}
}

func (s *StudentV2) addAge(a int) {
	s.Age += a
}

func DeepDivePointer() {

}
