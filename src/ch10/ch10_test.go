package ch10_test

import (
	"ch10/person"
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

type innerStruct struct {
	a int
	b int
	field1 string
}


type innerStruct2 struct {
	a int
	b int
}

type innerStruct3 struct {
	a2 int
	string
}


type StructType struct {
	field1 int "field1 int"
	field2 string "field2 string"
	field3 float32 "field3 float32"
	int "field4 int"
	innerStruct "field5 innerStruct"
	innerStruct2 "field6 innerStruct2"
	innerStruct3 "field7 innerStruct3"
}

func TestStructTag(t *testing.T) {
	a := StructType{
		1, "222", 3.222, 5,
		innerStruct{1, 2, "hello"},
		innerStruct2{3, 4},
		innerStruct3{9, "hello"},
	}

	for i := 0; i < 5; i++ {
		t := reflect.TypeOf(a)
		f := t.Field(i)
		println(f.Tag)
	}

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", a.innerStruct.a)
	fmt.Printf("%v\n", a.innerStruct2.a)
	fmt.Printf("%v\n", a.a2)
	fmt.Printf("%v\n", a.field1)
	fmt.Printf("%v\n", a.innerStruct.field1)

}

type myInt int

var a myInt = 5

func (a *myInt) addFive()  {
	*a = (*a) + myInt(5)
}

// not allowed
//func (a myInt) addFive()  {
//	a = a + 5
//}

func TestAliasTypeMethod(t *testing.T) {
	println(a)
	a.addFive()
	println(a)
}

type List []int

func (l *List) append(n int)  {
	*l = append(*l, n)
}

func (l List) len() int  {
	return len(l)
}

func TestPointerAndValueCall(t *testing.T) {
	// call by value
	l := List{}
	l.append(1)
	l.append(2)
	println(l.len())

	// call by pointer
	pl := new(List)
	pl.append(3)
	pl.append(4)
	println(pl.len())
}

func TestUnExportedField(t *testing.T) {
	p := person.Person{}
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.SetFirstName("Yiwang")
	p.SetLastName("Zheng")
	fmt.Printf("%v", p)
}

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

func TestCar(t *testing.T) {
	car := Car{}
	fmt.Printf("%v", car)
	//car.Start()
	//car.Stop()
}

type Point struct {
	x int
	y int
}

type NamedPoint struct {
	Point
	name string
}

func (npp *NamedPoint) moveLeft() {
	npp.x--
}

func TestNamedPoint(t *testing.T) {
	npp := new(NamedPoint)
	npp.x = 1
	npp.y = 1
	npp.name = "hello"
	fmt.Printf("%v\n", npp)

	npp.moveLeft()
	fmt.Printf("%#v\n", npp)

	np := NamedPoint{Point{2, 2}, "hello2"}
	fmt.Printf("%T\n", np)
}

type Base struct{}
func (Base) Magic() { fmt.Print("base magic ") }
func (b Base) MoreMagic() {
	b.Magic()
	b.Magic()
}
type Voodoo struct {
	Base
}
func (Voodoo) Magic() { fmt.Println("voodoo magic") }

func TestVoodoo(t *testing.T) {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}

func (is innerStruct) String() string  {
	//return string(is.a) + " / " + string(is.b) + " / " + is.field1
	return fmt.Sprintf("%v / %v / %v\n", is.a, is.b, is.field1)
}

func TestStructString(t *testing.T) {
	tmp := innerStruct{a:1, b:2, field1:"hello"}

	runtime.SetFinalizer(&tmp, func(ptr *innerStruct) {
		ptr.a = 0
		ptr.b = 0
		ptr.field1 = ""
		fmt.Printf("%v", tmp)
	})

	fmt.Printf("%v", tmp)

}

type Stringer interface {
	String() string
}

func TestCheckStringer(t *testing.T) {
	var v Stringer
	v = innerStruct{}
	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String())
		fmt.Printf("v implements String(): %s\n", v.String())
	}
}
