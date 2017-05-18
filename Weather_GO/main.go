package main

import (
  "flag"
  "fmt"
)

var zipcode int

func init() {
  flag.IntVar(&zipcode ,"z", 80203, "Zipcode flag to store into var")
  flag.Parse()
}

func main(){
   fmt.Println("Initiliazing Weather via:", zipcode )
}
