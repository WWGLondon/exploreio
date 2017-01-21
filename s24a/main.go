// S24a: A counting reader, count the number of bytes read in total.
//
//     $ cat main.go | go run main.go
//     n (io.Copy) = 550, n (CountingReader) = 550
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// TODO: implement a reader that counts the total number of bytes read. 12 lines.
// ...
// ...
// ...
// ...

// ...
// ...
// ...
// ...
// ...

// ...
// ...
// ...

func main() {
	cr := &CountingReader{r: os.Stdin}
	n, err := io.Copy(ioutil.Discard, cr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("n (io.Copy) = %d, n (CountingReader) = %d\n", n, cr.Count())
}