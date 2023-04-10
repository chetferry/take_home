package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("please enter a file name")
		os.Exit(1)
	}
	fName := args[1]
	fInfo, err := os.Stat(fName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(fName + " file not found")
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fLength := fInfo.Size()
	// open the file in write only mode (the last argument is ignored)
	fp, err := os.OpenFile(fName, os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	// overwrite each byte in the file 3 times with random data
	fmt.Println("shredding " + fName)
	for ii := 0; ii < 3; ii++ {
		for jj := int64(0); jj < fLength; jj++ {
			b := byte(rand.Intn(255))
			_, err = fp.WriteAt([]byte{b}, jj)
			if err != nil {
				panic(err)
			}
		}
	}
	err = fp.Close()
	if err != nil {
		panic(err)
	}
	// delete the file
	err = os.Remove(fName)
	if err != nil {
		panic(err)
	}
	fmt.Println("deleted ", fName)
}
