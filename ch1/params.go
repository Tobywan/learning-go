package ch1

import (
	"fmt"
	"os"
	"strings"
)

func main_params() {
	//	main_concat()
	//main_join()
	ex1_2()
}

func main_concat() {

	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func main_join() {

	fmt.Println(strings.Join(os.Args[1:], " "))
}

func ex1_2() {

	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}

}
