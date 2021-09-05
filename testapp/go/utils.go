package main

import (
	"fmt"
	"io"
	"os"
)

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func CopyFile(src string, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func PrintSectionHeader(s string) {
	fmt.Println("\n----------------------------------")
	fmt.Println(s)
	fmt.Println("----------------------------------")
}
