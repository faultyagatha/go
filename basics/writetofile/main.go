package main

import "os"

// to see real-time changes, use terminal:
// tail -f ./writefoo

// os.WriteFile doesn't require to fluch the buffer.
// Writes to os are direct syscalls to write to the file.
func writeFile(name string, data []byte) error {
	err := os.WriteFile(name, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	writeFile("writefoo", []byte("foo"))
	writeFile("writefoo", []byte("bar"))
}
