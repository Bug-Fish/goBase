package main

import (
	"fmt"
	"goBasic/package2"
	"io"
	"math"
	"net"
	"strconv"
	"time"
)

// https://juejin.cn/post/6844904119774216206

func main() {
	pktest.Func1()
	pktest.Func2()

	//var a int
	//var b float32
	//var c, d float64
	//e, f := 9, 10
	//var g = "jiaoyuzhang"

	_, v, _ := getData()
	fmt.Println(v)

	const zjy = "3243"
	fmt.Println(zjy)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	m := 1
	for m < 200 {
		m += m
	}
	fmt.Println(m)

	//for {
	//	fmt.Println("wait")
	//}

	fmt.Println(add(23, 452))

	fmt.Println(swap("123", "314"))

	fmt.Println(split(12))

	defer fmt.Println("end")
	fmt.Println("hello")

	//var arr [10]int

	str := [4]string{
		"aaa",
		"bbb",
		"ccc",
		"ddd",
	}
	fmt.Println(str)

	a := str[0:4]
	b := str[1:3]
	fmt.Println(a, b)
	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(str)
	fmt.Println(append(a, "e3"))

	abc := make([]int, 5, 12)
	abc = abc[:cap(abc)]
	abc = abc[1:]

	ming := new(people)
	ming.name = "jiaoyuzhang"
	ming.age = 19
	fmt.Println(ming)

	var p people
	p.age = 43
	p.name = "cxr"
	fmt.Println(p)

	pp := &people{"retn", 413}
	fmt.Println(pp)

	yu := Vertex{32, 3}
	fmt.Println(yu.Abs())

	v1 := Vertex{1, 1}
	v2 := &Vertex{1, 1}
	v1.test1()
	v2.test2()
	fmt.Println(v1)
	fmt.Println(v2)

	chick := Chicken{}
	var duck Duck
	duck = chick
	duck.DuckGo()
	duck.Quack()

	go f("hello goroutine")

	//listener, err := net.Listen("tcp", "localhost:8000")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		log.Print(err)
	//		continue
	//	}
	//	go handleConn(conn)
	//}

	message := make(chan string, 6)
	go func() {
		message <- "ping"
	}()
	msg := <-message
	fmt.Println(msg)

	cs := make(chan string)
	go makeCakeAndSend(cs, 5)
	go receiveCakeAndPack(cs)
	time.Sleep(3 * 1e9)

	strbry_cs := make(chan string)
	choco_cs := make(chan string)
	go makeCake(choco_cs, "qiaokeli", 3)
	go makeCake(strbry_cs, "choco_cs", 3)
	go receiveCake(strbry_cs, choco_cs)
	time.Sleep(2 * 1e9)

}

func getData() (int, int, int) {
	return 2, 4, 8
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

type people struct {
	name string
	age  int
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) test1() {
	v.Y++
	v.X++
}

func (v *Vertex) test2() {
	v.Y++
	v.X++
}

type Duck interface {
	Quack()
	DuckGo()
}

type Chicken struct {
}

func (c Chicken) Quack() {
	fmt.Println("gagaga")
}
func (c Chicken) DuckGo() {
	fmt.Println("gogogogo")
}

func f(msg string) {
	fmt.Println(msg)
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func makeCakeAndSend(cs chan string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := "Strawberry Cake" + strconv.Itoa(i)
		cs <- cakeName
	}
	close(cs)
}

func receiveCakeAndPack(cs chan string) {
	for s := range cs {
		fmt.Println("Packing received cake: ", s)
	}
}

func makeCake(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + "cake " + strconv.Itoa(i)
		cs <- cakeName
	}
	close(cs)
}

func receiveCake(strbry_cs chan string, choco_cs chan string) {
	strbry_closed, choco_closed := false, false
	for {
		if strbry_closed && choco_closed {
			return
		}
		fmt.Println("wait for nrw cake... ")
		select {
		case cakeName, strbry_ok := <-strbry_cs:
			fmt.Println(strbry_ok)
			if !strbry_ok {
				strbry_closed = true
				fmt.Println("...chaomei close")
			} else {
				fmt.Println("receive a new chaomei cake", cakeName)
			}
		case cakeName, choco_ok := <-choco_cs:
			if !choco_ok {
				choco_closed = true
				fmt.Println("...chotolate close")
			} else {
				fmt.Println("receive a new chotolate cake", cakeName)
			}
		}
	}
}
