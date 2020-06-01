package main
import(
  "fmt"
  "strconv"
  "sort"
  )
func main(){
  sli := make([]int, 3)
  i := 0
  for true{
    var x string;
    fmt.Println("Enter the No.            (type 'X' or 'x' to quit)");
    fmt.Scan(&x);
    n,err := strconv.ParseInt(x, 10, 64)
    if x != "x" && x != "X"{
      if err == nil{
        c := int(n);
        if i <2{
          //fmt.Printf("i = %d c = %d\n", i, c)
          sli[i] = c;
          //fmt.Printf("sli[%d] = %d\n", i, sli[i])
          i++;
          sort.Ints(sli)
          fmt.Println(sli)
        }else{
          sli = append(sli, c);
          sort.Ints(sli)
          fmt.Println(sli)
        }
      }else{
        fmt.Println("Please Enter appropiate values");
        continue;
      }
    }else{
      break;
    }
  }
}
