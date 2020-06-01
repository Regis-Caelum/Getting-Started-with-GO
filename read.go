package main
import(
    "fmt"
    "os"
    "strings"
    "bufio"
)
type name struct {
	fname string
	lname string
}

func main() {
  tmp := bufio.NewReader(os.Stdin)
	names := make([]name, 0, 3)
	fmt.Println("Enter the name of the text file with destination ( '/tmp/file.txt' ):")

  f, err := tmp.ReadString('\n')

  if err == nil{
    f = strings.Trim(f, "\n")
    //fmt.Println(f)

  	file, err := os.Open(f)
  	if err == nil{
      scanner := bufio.NewScanner(file)
    	for scanner.Scan() {
    		tmpname := strings.Split(scanner.Text(), " ")
    		var ame name
    		ame.fname, ame.lname = tmpname[0], tmpname[1]
    		names = append(names, ame)
    	}

    	file.Close()
    	for _, v := range names {
    		fmt.Println(v.fname, v.lname)
    	}
    }else{
      fmt.Println(err)
    }
    }
}
