package main

import (
	"flag"
	"github.com/gpmgo/gopm/modules/log"
	"os"
	"runtime/pprof"
)

func loop(n int)  {
	for i := 0; i < n; i++ {
	}

	loopA(n / 2)
	loopB(n / 2)
	println("loop done.")
}

func loopA(n int)  {
	for i := 0; i < n; i++ {
	}
	println("loopA done.")
}

func loopB(n int)  {
	for i := 0; i < n; i++ {
	}
	println("loopB done.")
}


var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")


func main() {
	flag.Parse()
	if *cpuprofile != "" {
		println("cpu profile: %v", *cpuprofile)
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("%v", err)
		}
		if e := pprof.StartCPUProfile(f); e != nil {
			log.Fatal("Start CPU Profile failed.")
		}
		defer pprof.StopCPUProfile()
	}

	loop(10000000000)

}

