package test

import (
	fm "fmt"
	"unsafe"
	"reflect"
	"time"
)

var (
	A string ="ab"
	B int = 12
)
type Type int
type Struct struct{
}
const (
	c = iota
	x string = "x"
	y = "y"
	_
	d = iota*2
	e
)

func init(){
	fm.Println("init");
}
func Test() int{
	fm.Println("test");
	return 0;
}

func tst() {
	var i, l int = 1, 2
	var b rune = 0
	var j complex128 = 1+2i
	str, _, str2 := "my", "own", "world"
	var k interface{}
	arry := []string{"you","are","mine"}

	i++
	i = l % i
	i <<= uint32(l)
	fm.Println(i)
	if i != l {
		fm.Println(i ^ l)
	} else {
		fm.Println(i & l)
	}
	k= j
	switch k.(type) {
	case int:
		fm.Println("int")
	case complex128:
		fm.Println("complex128")
	default:
		fm.Println("not match")
	}
	for i:=0;i<3;i++{
		Test()
	}
	for key,value := range(arry){
		fm.Print(key)
		fm.Println(value)
	}
	goto test
	fm.Println(str+ str2);
	fm.Println(unsafe.Sizeof(b))
	fm.Print(j);
	fm.Println(reflect.TypeOf(j))
	fm.Print(A)
	fm.Println(B)
	fm.Print(reflect.TypeOf(x))
	fm.Println(reflect.TypeOf(y))
	test:
	for {
		fm.Print(c)
		fm.Print(d)
		fm.Println(e)
		time.Sleep(1)
		break
	}
}
