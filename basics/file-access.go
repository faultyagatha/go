package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/* file access is linear access, not random access. Basic operations:
1. Open - get handle for access
2. Read - read bytes into []byte
3. Write - write []byte into file
4. Close - release handle
5. Seek - move read/write head
*/

func simpleRead() {
	//with this approach, large files cause problems
	dat, e := ioutil.ReadFile("test.txt")
	fmt.Println("dat: ", dat, "err: ", e)
}

func simpleWrite() {
	dataStr := "Hello, world"
	dataBytes := []byte(dataStr) //convert to the needed type
	e := ioutil.WriteFile("test.txt", dataBytes, 0777)
	fmt.Println("err: ", e)
}

func readWriteOs() {
	f, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	l, err := f.WriteString("Hello, World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		os.Exit(1)
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	simpleWrite()
	simpleRead()
	readWriteOs()
}
