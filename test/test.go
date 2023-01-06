package main

import (
	sh1 "crypto/sha1"
	"fmt"

)

func errors(t string, err error) {
	if err != nil {
		fmt.Printf(t, err)
	}
}

func main() {
	salt := "123"
	h := sh1.New()

	m := []byte("I Live in Winnipeg")
	// h.Write(m)

	h.Write(m)
	// errors(" hash Write Error", err)

	
	// fmt.Printf("a1 = %x \n ", a1)
	s1 := h.Sum([]byte(salt))
	fmt.Printf("a1 = %x \n ", s1)

}
