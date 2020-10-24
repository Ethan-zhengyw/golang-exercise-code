package ch11_test

import (
	"ch11/sort"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type List []int

func (l List) Len() int {
	return len(l)
}

//func (l *List) Append(val int) {
//	*l = append(*l, val)
//}



func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

type Popper interface {
	Pop() int
}

func (l List) Pop() int {
	tmp := l[0]
	return tmp
}

type Lenner interface {
	Len() int
}

func CountInto(a Appender, start, end int)  {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func LongEnough(l Lenner) bool {
	return l.Len() * 10 > 42
}

func TestList(t *testing.T) {
	var appender Appender
	var lst List

	fmt.Printf("appenders: %v\n", appender)
	CountInto(&lst, 1, 2)
	fmt.Printf("appenders: %v\n", appender)

	//CountInto(lst, 1, 5)
	//CountInto(&(&lst), 1, 5)
	//if LongEnough(lst) {
	//	fmt.Printf("list %v is long enough, len %v\n", lst, lst.Len())
	//}

	//var plst *List
	//plst = new(List)
	//CountInto(plst, 1, 5)
	//
	//if LongEnough(plst) {
	//	fmt.Printf("plist %v is long enough, len %v\n", plst, plst.Len())
	//}

}

func TestPointerToInterface(t *testing.T) {
	list := new(List)
	var appender Appender
	appender = list
	fmt.Printf("%s\n", appender)
}

func TestReceiverWithoutInterface(t *testing.T) {
	// without interface
	// everything seems right

	list := List{}
	plist := &list
	CountInto(plist, 1, 5)

	// value receiver
	list.Pop()
	plist.Pop()

	// pointer receiver
	(*(&list)).Append(1)
	plist.Append(2)

	fmt.Printf("%v\n", list)

}

func TestValueToInterface(t *testing.T) {
	list := List{}
	var popper Popper
	popper = list
	popper = &list
	fmt.Printf("%s\n", popper)

	CountInto(&list, 1, 5)

	fmt.Printf("%v\n", list.Len())
	//list.Pop()

	var appender Appender
	//appender = list  // Append method has a pointer receiver
	appender = &list
	fmt.Printf("%s\n", appender)
}

func TestSortInts(t *testing.T) {
	data := []int{23, 23, 231, 21, 23, 56, 865, -1, 0}
	fmt.Printf("%v\n", data)
	sort.Sort(sort.IntArray(data))
	if !sort.IntsAreSorted(data) {
		panic("data are not sorted!")
	}
	fmt.Printf("%v\n", data)
}

func TestSortStrings(t *testing.T) {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "saturday"}
	fmt.Printf("%v\n", data)
	sort.Sort(sort.StringArray(data))
	if !sort.StringAreSorted(data) {
		panic("data are not sorted!")
	}
	fmt.Printf("%v\n", data)
}

type day struct {
	num int
	shortName string
	longName string
}

type dayArray []*day

func (data *dayArray) Len() int {
	return len(*data)
}

func (data *dayArray) Less(i, j int) bool {
	return (*data)[i].num < (*data)[j].num
}

func (data *dayArray) Swap(i, j int) {
	(*data)[i], (*data)[j] = (*data)[j], (*data)[i]
}

func (data *dayArray) Print()  {
	for _, day := range *data {
		fmt.Printf("%s\t", day.longName)
	}
	fmt.Println("")
}

func TestSortDayArray(t *testing.T) {
	data := dayArray{
		&day{0, "SUN", "Sunday"},
		&day{2, "TUE", "Tuesday"},
		&day{1, "MON", "Monday"},
		&day{4, "THU", "Thursday"},
		&day{3, "WED", "Wednesday"},
		&day{5, "FRI", "Friday"},
		&day{6, "SAT", "Saturday"},
	}

	data.Print()
	sort.Sort(&data)
	data.Print()
}


type Any interface {}

func TestEmptyInterface(t *testing.T) {
	var val Any
	a := 5
	b := "hello"
	c := &day{6, "SAT", "Saturday"}

	val = a
	fmt.Printf("%v\n", val)

	val = b
	fmt.Printf("%v\n", val)

	val = c
	fmt.Printf("%v\n", val)

	switch t := val.(type) {
	case *day:
		fmt.Printf("Pointer to day: %v\n", t)
	default:
		fmt.Printf("%t: %v\n", t, t)
	}
}

func TestEmptyInterfaceArray(t *testing.T) {
	data1 := []Any{
		1,
		//"hello",
		//&day{6, "SAT", "Saturday"},
	}

	data2 := dayArray{
		//&day{0, "SUN", "Sunday"},
		//&day{2, "TUE", "Tuesday"},
		//&day{1, "MON", "Monday"},
	}

	data3 := [3]int{ 1, 2, 3}

	fmt.Printf("%v\n", data1)
	fmt.Printf("%v\n", data2)

	//data1 = data2
	//data1 = data3

	for i, d := range data2 {
		data1[i] = d
	}

	fmt.Printf("%v %d\n", data1, unsafe.Sizeof(data1))
	fmt.Printf("%v %d\n", data2, unsafe.Sizeof(data2))
	fmt.Printf("%v %d\n", data3, unsafe.Sizeof(data3))

}

func TestReflect(t *testing.T) {
	data1 := []Any{
		1,
		//"hello",
		//&day{6, "SAT", "Saturday"},
	}

	data2 := dayArray{
		&day{0, "SUN", "Sunday"},
		&day{2, "TUE", "Tuesday"},
		&day{1, "MON", "Monday"},
	}

	data3 := [3]int{ 1, 2, 3}

	fmt.Printf("%v, %v, %v\n", reflect.TypeOf(data1), reflect.TypeOf(data1).Kind(), reflect.ValueOf(data1))
	fmt.Printf("%v, %v, %v\n", reflect.TypeOf(data2), reflect.TypeOf(data2).Kind(), reflect.ValueOf(data2))
	fmt.Printf("%v, %v, %v\n", reflect.TypeOf(data3), reflect.TypeOf(data3).Kind(), reflect.ValueOf(data3))

	v := reflect.ValueOf(&data1)
	fmt.Printf("v: %v\n", v)
	fmt.Printf("v: %v\n", v.Elem())

	data4 := 4
	v4 := reflect.ValueOf(&data4)
	fmt.Printf("v: %v\n", v4)
	fmt.Printf("v: %v\n", v4.Elem())

	fmt.Printf("%v, %v\n", v, v.Interface())
}
