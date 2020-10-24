package ch04

import os2 "os"

func DoGoos()  {
	os := os2.Getenv("OS")
	path := os2.Getenv("PATH")

	println(os)
	println(path)
}
