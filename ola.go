package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

//Ola func....
func Ola(name string) string {
  if len(name) > 0 {
    return prefixoOlaPortugues + name;
  } 
   return prefixoOlaPortugues + "mundo";

}

func main() {
	fmt.Println(Ola("Chris"))
}
