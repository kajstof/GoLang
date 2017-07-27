package main

import "fmt"
import "math"
import "strings"
import "sort"
import "decr1"
import "decr2"
//import "os"
//import "log"
//import "io/ioutil"
//import "strconv"

func main() {
    fmt.Println("Hello World")
    fmt.Println("ZAŻÓŁĆ GĘŚLĄ JAŹŃ")
    fmt.Println("zażółć gęślą jaźń")
    var age int = 40
    var favNumber float64 = 1.6180339
    randNum := 1
    fmt.Println(randNum)
    fmt.Println(age, " ", favNumber)
    var numOne = 1.000
    var num99 = .9999
    fmt.Println(numOne - num99)
    const pi float64 = 3.14159265
    //var (
    //    varA = 2
    //    varB = 3
    //)
    var myName string = "Krzysztof Krysiak"
    fmt.Println(len(myName))
    fmt.Println(myName + " is a robot")
    var isOver40 bool = true
    fmt.Printf("%.3f | %T \n", pi, pi)
    fmt.Printf("%t \n", isOver40)
    fmt.Printf("%d, %b, %c, %x, %e \n", 100, 100, 38, 17, pi)
    fmt.Println("true && false = ", true && false)
    fmt.Println("true || false = ", true || false)
    fmt.Println("!true = ", !true)
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i++
    }
    for j := 0; j < 5; j++ {
        fmt.Println(j)
    }
    yourAge := 18
    if yourAge >= 16 {
        fmt.Println("You can drive")
    } else {
        fmt.Println("You can't drive")
    }
    switch yourAge {
    case 16:
        fmt.Println("sixteen")
    case 18:
        fmt.Println("eighteen")
    default:
        fmt.Println("Rest of the world value")
    }
    var favNums2 [5]float64
    favNums2[0] = 163
    favNums2[1] = 24235
    favNums2[2] = 4234
    favNums2[3] = 1.11
    favNums2[4] = 32.33
    fmt.Println(favNums2[3])
    favNums3 := [5]float64{1, 2, 3, 4, 5}
    for i, value := range favNums3 {
        fmt.Println(value, i)
    }
    for _, value := range favNums3 {
        fmt.Println(value)
    }
    numSlice := []int{5, 4, 3, 2, 1}
    numSlice2 := numSlice[3:5]
    fmt.Println("numSlice2[0] =", numSlice2[0])
    numSlice3 := numSlice[:2]
    fmt.Println("numSlice3[:2] =", numSlice3[:2])
    numSlice4 := make([]int, 5, 10)
    copy(numSlice4, numSlice)
    fmt.Println(numSlice3[0])
    numSlice4 = append(numSlice4, 0, -1)
    fmt.Println(numSlice4[6])
    presAge := make(map[string]int)
    presAge["TheodoreRoosevelt"] = 42
    fmt.Println(presAge["TheodoreRoosevelt"])
    fmt.Println(len(presAge))
    presAge["John F. Kennedy"] = 43
    fmt.Println(len(presAge))
    delete(presAge, "John F. Kennedy")
    fmt.Println(len(presAge))
    listNums := []float64{1, 2, 3, 4, 5}
    fmt.Println(addThemUp(listNums))
    numb1, numb2 := next2Values(5)
    fmt.Println(numb1, numb2)
    fmt.Println(subtractThem(1, 2, 3, 4, 5))
    numb3 := 3
    doubleNum := func() int {
        numb3 *= 2
        return numb3
    }
    fmt.Println(doubleNum())
    fmt.Println(doubleNum())
    fmt.Println(factorial(5))
    defer printTwo()
    printOne()
    fmt.Println(safeDiv(3, 0))
    fmt.Println(safeDiv(3, 2))
    demPanic()
    x := 0
    changeXVal(&x)
    fmt.Println("x =", x)
    fmt.Println("Memory Address for x =", &x)
    yPtr := new(int)
    changeYValNow(yPtr)
    fmt.Println("y =", *yPtr)
    rect1 := RectangleX{leftX: 0, topY: 50, height: 10, width: 10}
    rect2 := RectangleX{0, 12, 5, 5}
    fmt.Println("Rectangle is", rect1.width, "wide")
    fmt.Println("Rectangle2 is", rect2.height, "height")
    fmt.Println("Area of rectangle =", rect1.area())
    rect := Rectangle{20, 50}
    circ := Circle{4}
    fmt.Println("Rectangle Area =", getArea(rect))
    fmt.Println("Circle Area =", getArea(circ))
    // String examples
    sampString := "Hello World"
    fmt.Println(strings.Contains(sampString, "lo"))
    fmt.Println(strings.Index(sampString, "lo"))
    fmt.Println(strings.Count(sampString, "lo"))
    fmt.Println(strings.Replace(sampString, "l", "x", 3))
    csvString := "1,2,3,4,5,6"
    fmt.Println(strings.Split(csvString, ","))
    listOfLetters := []string{"c", "a", "b"}
    sort.Strings(listOfLetters)
    fmt.Println("Letters: ", listOfLetters)
    listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")
    fmt.Println("Numbers: ", listOfNums)
    fmt.Println(":)")
    decr1.decr1()
    decr2.decr2()
}

func addThemUp(numbers []float64) float64 {
    sum := 0.0
    for _, val := range numbers {
        sum += val
    }
    return sum
}

func next2Values(number int) (int, int) {
    return number + 1, number + 2
}

func subtractThem(args ...int) int {
    finalValue := 0
    for _, value := range args {
        finalValue -= value
    }
    return finalValue
}

func factorial(num int) int {
    if num == 0 {
        return 1
    }
    return num * factorial(num-1)
}

func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

func safeDiv(num1, num2 int) int {
    defer func() {
        //fmt.Println(recover())
        recover()
    }()
    solution := num1 / num2
    return solution
}

func demPanic() {
    defer func() {
        fmt.Println(recover())
    }()
    panic("PANIC")
}
func changeXVal(x *int) {
    *x = 2
}
func changeYValNow(yPtr *int) {
    *yPtr = 100
}

type RectangleX struct {
    leftX  float64
    topY   float64
    height float64
    width  float64
}

func (rect *RectangleX) area() float64 {
    return rect.width * rect.height
}

type Shape interface {
    area() float64
}

type Rectangle struct {
    height float64
    width  float64
}

type Circle struct {
    radius float64
}

func (r Rectangle) area() float64 {
    return r.height * r.width
}

func (c Circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
    return shape.area()
}
