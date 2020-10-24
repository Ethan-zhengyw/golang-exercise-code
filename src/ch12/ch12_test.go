package ch12_test

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/gob"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestScanln(t *testing.T) {
	var s1, s2 string
	fmt.Scanln(&s1, &s2)
	fmt.Printf("%v, %v\n", s1, s2)
}

func TestSscanf(t *testing.T) {
	format := "%f / %d / %s"
	input := "32.2 / 23 / hello"
	var f float32
	var i int
	var s string
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Printf("%v %v %v\n", f, i, s)
}

func TestBufio(t *testing.T) {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err == nil {
		println(input)
	}
}

func TestWriteToScreen(t *testing.T) {
	os.Stdout.WriteString("hello world!\n")
}

func TestFlagBool(t *testing.T) {
	newLine := flag.Bool("n", false, "print on new line")
	const (
		Space = " "
		NewLine = "\n"
	)
	flag.PrintDefaults()
	fmt.Printf("%v\n", newLine)
	flag.Parse()
	var s string = ""

	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}

	if *newLine {
		s += NewLine
	}

	fmt.Printf("%s", s)
}

func cat(r *bufio.Reader)  {
	for {
		s, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", s)
	}
}

func TestBufioReadFile(t *testing.T) {
	f, err := os.Open("ch12_test.go")
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	cat(r)

}

func TestBufferFlush(t *testing.T) {
	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	buf.Flush()
}

func TestWriteCol(t *testing.T) {
	inputFile, _ := os.Open("ch12_test.go")
	outputFile, _ := os.OpenFile("tmp.txt", os.O_WRONLY | os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)

	for {
		inputSring, _, err := inputReader.ReadLine()
		if err == io.EOF {
			println("EOF")
			break
		}
		outputString := inputSring[2:5]
		_, err2 := fmt.Fprintf(outputWriter, "%v\n", string(outputString))
		//_, err2 := outputWriter.WriteString(string(outputString) + "\n")
		if err2 != nil {
			println(err2)
			break
		}
	}
	outputWriter.Flush()
}

func TestMarshal(t *testing.T) {
	type Address struct {
		Type 	string
		City 	string
		Country string
	}
	type VCard struct {
		FirstName	string
		LastName	string
		Addresses	[]*Address
		Remark		string
	}

	pa := &Address{"private", "Aartselaar","Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa,wa}, "none"}
	fmt.Printf("%v\n", vc)
	js, err := json.Marshal(vc)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("JSON format: %s\n", js)

	outputFile, _ := os.OpenFile("tmp.txt", os.O_WRONLY | os.O_CREATE, 0666)
	defer outputFile.Close()

	enc := json.NewEncoder(outputFile)
	err = enc.Encode(vc)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func TestUnMarshal(t *testing.T) {
	b := []byte("{\"FirstName\":\"Jan\",\"LastName\":\"Kersschot\",\"Addresses\":[{\"Type\":\"private\",\"City\":\"Aartselaar\",\"Country\":\"Belgium\"},{\"Type\":\"work\",\"City\":\"Boom\",\"Country\":\"Belgium\"}],\"Remark\":\"none\"}")
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", f)
}

func TestGob(t *testing.T) {
	type Address struct {
		Type 	string
		City 	string
		Country string
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	err := enc.Encode(Address{"C", "S", "Co"})
	if err != nil {
		println(err)
	}

	var q Address
	err = dec.Decode(&q)
	if err != nil {
		println(err)
	}
	fmt.Printf("%v\n", q)
}

func TestHashSum(t *testing.T) {
	res := []byte{}

	hasher := sha1.New()

	if _, err := io.WriteString(hasher, "hello"); err != nil {
		println(err)
	}

	s := hasher.Sum(res)
	fmt.Printf("%v", s)
	fmt.Printf("%x", s)


}

func TestMd5Sum(t *testing.T) {
	md5Helper := md5.New()

	_, _ = io.WriteString(md5Helper, "Hello World!")

	var buf []byte

	res := md5Helper.Sum(buf)
	fmt.Printf("%x\n", res)
}

func TestUserDefinedError(t *testing.T) {
	errNotFound := errors.New("XXX Notfound")

	println(errNotFound.Error())
}
