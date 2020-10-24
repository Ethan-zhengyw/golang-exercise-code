package ch04

type Color int

const (
	Red Color = iota
	Yellow
	Blue
)

const (
	Unknown = iota
	Female
	Male
	male
)

func DoTypeConvert()  {
	intVal := 5
	float64Val := float64(intVal)
	println(float64Val)

	intVal2 := int(float64Val)
	println(intVal2)

	float32Val := float32(float64Val)
	println(float32Val)

	println(Female)
	println(Male)
	println(male)

	println("red", Color(Red))
	println(Color(Yellow))
	println(Color(Blue))
	println(Blue)
}
