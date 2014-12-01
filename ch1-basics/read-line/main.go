//package declarate
package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){

  read := bufio.NewReader(os.Stdin)

  fmt.Print("Name : ")
  name, _ := read.ReadString('\n')
  fmt.Printf("Hello %s",name)

}
