package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func foo() {
  for i := 0; i < 45; i++ {
    fmt.Println("Foo: ", i)
  }
  wg.Done()
}

func bar() {
  for i := 0; i < 45; i++ {
    fmt.Println("Bar: ", i)
  }
  wg.Done()
}

func main() {
  wg.Add(2)
  go foo()
  go bar()
  wg.Wait()
  //nothing happens -> the program will exit without waiting for the goroutines to return. So we add wait groups.
  //the result is still basically concurrent.. see concurrentwithwait.go

}
