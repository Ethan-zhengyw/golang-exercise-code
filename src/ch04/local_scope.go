package ch04

var a string = "G"

func printGlobalVar()  {
	println(a)
}

func defineLocalVarAndPrintGlobalVar()  {
	a := "L"
	println(a)
}

func Driver()  {
	printGlobalVar()
	defineLocalVarAndPrintGlobalVar()
	printGlobalVar()
}
