package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
)

func main() {
	ageP := new(int)
	ageP = flag.Int("age", 0, "Enter the start age")

	fmt.Fprintf(os.Stdout, "The value is %v", *ageP)
}

type myIntValue int

func (v *myIntValue) String() string {
	return strconv.Itoa(int(*v))
}

func (v *myIntValue) Set(val string) error {
	nv, err := strconv.ParseInt(val, 10, 0)
	*v = myIntValue(nv)
	return err
}

func (v *myIntValue) Get() interface{} {
	return int(*v)
}

func NewMyIntValue(val int, p *int) *myIntValue {
	*p = val
	return (*myIntValue)(p)
}


