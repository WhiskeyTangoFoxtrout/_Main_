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

    http.HandleFunc("/", handler)//http://localhost:52
    http.HandleFunc("/view/",viewHandler) This should display a website and the title of the website from the URL
    log.Fatal(http.ListenAndServer(":52",nil))//need to work on ports
    
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

    
    
}
