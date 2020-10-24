package ch13_test

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)
import "testing"

func MustPanic()  {
	panic("panic from MustPanic")
}

func protect(f func())  {
	fp, err := os.Create("profile.out")
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()


	defer func() {
		println("Done")
		if e := recover(); e != nil {
			fmt.Printf("recover from panic: %v\n", e)
		}
	}()

	f()
}

func TestRecover(t *testing.T) {
	protect(MustPanic)
	println("should see me.")
}

func TestNotRecover(t *testing.T) {
	MustPanic()
	println("should not see me.")
}

func TestRecover2(t *testing.T) {
	defer func() {recover()}()
	MustPanic()
	println("should see me.")
}
