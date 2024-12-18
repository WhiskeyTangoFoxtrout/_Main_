//This is a server you would use to talk to my server
//ima eventually move it to another branch
//Note to self this is a http server
package main

import("fmt"
       "log"
       "net"
       "net/http"
       "io/buff")

func main(){

    goto ear()
    http.HandleFunc("/", handler)//http://localhost:52
    http.HandleFunc("/view/",viewHandler) This should display a website and the title of the website from the URL
    log.Fatal(http.ListenAndServer(":52",nil))//need to work on ports
    
}

type IP []byte

type IPAddr struct{

    IP IP
    IPv4 string
       
}

type Page Struct{

    Title string
    Body []byte
    
}

//ima need a save func but idk what ima save just yet. I just want it to echo for rn

func handler( r *http.Request){

    fmt.Fprintf(w "This is the Users BackEnd")//I think this will write out the URL
    
}

func ViewHandler(w http.ResponseWriter, t*http.Request){

    title := r.URL.Path[len("/view/"):] //this the title is the URL Path you can see on a website
    p, _loadPage(title)
    fmt.Fprintf(w, "<h1>s%</h1><div>%s</div>", p.Title, p.Body)//crude this helps make a website and displays the title
    
}

func ear(){//Ok so open port and have port listen for request and respond back with an "WE'RE LIVE FROM DA UNDAGROUND"
     
    ln, err := net.Listen("ip",ipv4)
       if err != nil{

              fmt.Fprintf("ErrorCode: 404")
              os.Exit(1)
              
       }

       fmt.Fprintf(conn, "SHould be http /HTTP/1.0\r\n\r\n")
       status,err := bufio.NewReader(conn).ReadString("\n")
       
}

func Lo(){//ima have to remember this location

    ipv4 := net.IPv4(49,7,10,15)//I think this the real ip
    I4Ski:= net.CIDRMask(16,32)//keep it multiples of 16 feel like subnets based on 16 or 8
       
}

