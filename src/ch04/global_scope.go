package ch04

var a1 string

func DoGlobalScope()  {
	a1 = "G"
	println(a)
	f1()
}

func f1()  {
	a1 := "O"
	println(a1)
	f2()
}

func f2()  {
	println(a1)
}
