package main

import "fmt"

func helloWorld(){
    fmt.Println("Hello, World!")
}

func helloName(){
    var myname string
    fmt.Printf("What is your name?\n")
    fmt.Scanf("%s", &myname)
    fmt.Printf("Hello, %s\n",myname)
}

func fizzBuzz(){
    for i:=0;i<=100;i++{
        if i%15==0{
	    fmt.Println("FizzBuzz")
	} else if i%3==0{
	    fmt.Println("Fizz")
	} else if i%5==0{
	    fmt.Println("Buzz")
	} else{
	    fmt.Println(i)
	}
    }
}

func main(){
    fmt.Println("Select a program\n1.Hello World\n2.Hello Name\n3.FizzBuzz")
    var input int
    fmt.Scanf("%d",&input)
    if input==1{
        helloWorld()
    } else if input==2{
        helloName()
    } else if input==3{
        fizzBuzz()
    } else{
        fmt.Print("Error, incorrect input")
    }
}
