package rstring

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	fmt.Println(RandomString(Alphabet, 10))
	fmt.Println(RandomString(Upper, 10))
	fmt.Println(RandomString(Lower, 10))
	fmt.Println(RandomString(Num, 10))
	fmt.Println(RandomString(SliceAppend(Lower, Num, Upper), 10))
}
