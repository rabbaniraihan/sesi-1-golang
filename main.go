package main

import "fmt"

func main() {
	var i int = 21
	fmt.Printf("%v\n", i)  //nilai 21
	fmt.Printf("%T\n", i)  //tipe data dari 21
	fmt.Printf("%%\n")     //menampilkan %
	fmt.Printf("%d\n", i)  //base 10 dari 21
	fmt.Printf("%o\n", 25) //base 8 dari 25

	fmt.Printf("%x\n", "f")    //base 16 f
	fmt.Printf("%X\n", "F 13") //base 16 F 13

	var j bool = true
	fmt.Println(j)             //j = true
	fmt.Println("\u042f")      //unicode
	fmt.Printf("%U\n", 0x042F) //format unicode

	var k float64 = 123.456
	fmt.Printf("%f\n", k) //float
	fmt.Printf("%E\n", k) //float scientific

}
