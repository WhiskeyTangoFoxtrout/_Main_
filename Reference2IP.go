//this is a refrence Book wrk also
type IP []byte

package main 
import("fmt"
       "net"
       "os")

func main(){

    if len(os.Args) !=2{

       fmt.Fprintf(os.Stderr,"this supposed to make a homemade IP address")
       os.Exit(1)

    }

    name  := os.Args[1]

    addr := net.ParseIP(name)
    if addr == nil{

         fmt.Println("NOT VALID")

    }else{

         fmt.Println("VALID ADDY", addr.String())

    }
    os.Exit(0)
}
