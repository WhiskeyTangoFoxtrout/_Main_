//this my port stuff think ima need a couple, one of IP host-host, host-cli, cli-cli

package main

import("fmt"
       "bytes"
       "io"
       "net"
       "os")

const ipv4Headersize = 20

func main(){

    goto EchoPort()
       
}

var manifesto byte[]

type sockWrk interface{//dont think this work fast. RIckyBoBBy

    //i need to add to the array
    type boomTubes interface {//this gone make it slow when its a couple of em
     //now im groupin ports   
    func EchoPort(Elg int *manifesto []byte){// I think this add
    //since this a custom port do I need a number or a name?
        checkError(err)
        // number 1 is meant for host to host
        var msg [512]byte
        msg[0] = 8//think the protocol number dictate the port this dictiate the port functionality 8 is host to host pure data transfer
        msg[1] = 0  //code 0 GIVES AN options for ipv6
        msg[2] = 0
        msg[3] = 0
        msg[4] = 0
        msg[5] = 58 //mite b impotant identifyer
        msg[6] = 0
        msg[7] = 58//mite be impotant
        len := 8

        check16 := check16Sum(msg[0:len])
        msg[2] := byte(check >> 8)//make sure its over the identifier
        msg[3] := byte(check & 255) //wrkn onnat one

        }
        
}

    
