package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    firstArg:=os.Args[1]
    if (firstArg == "add") || (firstArg == "Add") || (firstArg == "ADD"){
        num1, err1 := strconv.Atoi(os.Args[2])
        num2, err2 := strconv.Atoi(os.Args[3])
        if (err1==nil) && (err2 ==nil){
            sum:=num1+num2
            fmt.Printf("Hey! The sum of number is ")
            fmt.Printf("%v", sum)
        }
    }
    if (firstArg == "sub") || (firstArg == "Sub") || (firstArg == "SUB"){
        num1, err1 := strconv.Atoi(os.Args[2])
        num2, err2 := strconv.Atoi(os.Args[3])
        var sub int
        if (err1 == nil) && (err2 == nil){
            if (num1 >= num2){
                sub =num1-num2
            }
            if (num2 > num1){
                sub =num2-num1
            }
            fmt.Printf("Hey! The subtraction of numbers is %v", sub)
        }
    }
    if (firstArg == "mul") || (firstArg == "Mul") || (firstArg == "MUL"){
        num1, err1 := strconv.Atoi(os.Args[2])
        num2, err2 := strconv.Atoi(os.Args[3])
        if (err1 == nil) && (err2 == nil){
            mul := num1 * num2
            fmt.Printf("Hey! The Multiplication of numbers is %v", mul)
        }
    }
    if (firstArg == "div") || (firstArg == "Div") || (firstArg == "DIV"){
        num1, err1 := strconv.Atoi(os.Args[2])
        num2, err2 := strconv.Atoi(os.Args[3])
        if (err1 == nil) && (err2 == nil){
            div := num1/num2
            fmt.Printf("Hey! The division of numbers is %v", div)
        }
    }
}
