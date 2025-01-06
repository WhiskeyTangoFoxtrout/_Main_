//Hopin this telephone my mane

package main

import("fmt"
       "net")

 var globalDNSLookup = nil #still trynna figure a method where the user inputs a web address and I can look it up or convert it

func main(){

    if logn = 0 {
        
        goto lips    
    }else{searchBar(*usrNput)}
    
}

func lips(){

    conn, err := net.Dial("ipv4","Mainbackend.abc:80")
    if err != nil {

        os.Exit(1)//this should call 
      
    }

       conn, err := ln.Accept()
       if err != nil {
              
              os.Exit(1)
              
       }
       go handleConnection(conn)
}

func searchBar(r *Resolver){//**********************************WORK ON DIS*************************************

    //Sb,err := LookupAddr.UsrNputAddr
    sb,err := LookupCNAME(Search)([],err)//search is the userinput , and the C gile where its located
    
    
    
}
