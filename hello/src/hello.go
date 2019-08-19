//hello go program
package main

import "fmt"

func main() {
  fmt.Println("Hello, 世界")

  var num1 int = 1 
  var num2 int = 2

  fmt.Printf("%v + %v = %v\n",num1, num2, num1 + num2)

}

func addingfloat(num1 float64,num2 float64) float64{
  sum := num1 + num2
  return sum
}
