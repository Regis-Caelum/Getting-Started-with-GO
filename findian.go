package main
import "fmt"
func main(){
  var str string;
  fmt.Scan(&str);
  //fmt.Println(string(str[0]));
  /*if(string(str[0]) == "i"){
    fmt.Println("hi");
  }*/
  n := len(str) - 1;
  if(string(str[0]) != "i"  && string(str[0]) != "I"){
    fmt.Println("Not Found!");
  }else if (string(str[n]) != "n" && string(str[n]) != "N"){
    fmt.Println("Not Found!");
  }else{
    for i:= 0; i < len(str); i++ {
      if(string(str[i]) == "a" || string(str[i]) == "A"){
        fmt.Println("Found!");
        break;
      }
    }
 }
}
