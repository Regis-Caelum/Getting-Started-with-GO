package main
import (
  "fmt"
  "encoding/json"
)
type person struct{
  name string
  address string
}
func main(){
  var tmp person
  m := make(map[string]string)
  fmt.Println("Enter name :")
  fmt.Scan(&tmp.name)
  fmt.Println("Enter the address :")
  fmt.Scan(&tmp.address)
  m[tmp.name] = tmp.address
  obj,_ := json.Marshal(m)
  fmt.Println(string(obj))
}
