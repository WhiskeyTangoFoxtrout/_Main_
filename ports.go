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
        var msg [128]byte//65
        msg[0] = 8//think the protocol number dictate the port this dictiate the port functionality 8 is host to host pure data transfer
        msg[1] = 0  //code 0 GIVES AN options for ipv6
        msg[2] = 0
        msg[3] = 0
        msg[4] = 0
        msg[5] = 4 //mite b impotant identifyer
        msg[6] = 0
        msg[7] = 4//mite be impotant//this is ipv4 if get it right
        len := 8

        check16 := check16Sum(msg[0:len])
        msg[2] := byte(check >> 8)//make sure its over the identifier
        msg[3] := byte(check && 64) //wrkn onnat one feel like this a assec check

        }

       func SwitchBoard( swG int *manifesto []byte){ //this mo host to hose

              checkError(err)
              var keystnPipeln [256]byte//42
              keystnPipeln[0] = 28//this is made for biggerdata transfer host ot host
              keystnPipeln[1] = 0
              keystnPipeln[2] = 0
              keystnPipeln[3] = 0
              keystnPipeln[4] = 0
              keystnPipeln[5] = 4//ip i think
              keystnPipeln[6] = 0
              keystnPipeln[7] = 4
              len := 8

              check32 := check32Sum(keystnPipeln[0:8])
              keystnPipeln[2] := byte(check >> 28)
              keystnPipeln[3] := byte(check && 168)
              
       }

       func ampBoard(ampG int *mainifesto []byte){

              checkError(err)
              var KeystnPipeln [128]byte//16
              KeystnPipeln[0] = 27//if this proto 27 cant use TCP
              KeystnPipeln[1] = 0
              KeystnPipeln[2] = 0
              KeystnPipeln[3] = 0
              KeystnPipeln[4] = 0
              KeystnPipeln[5] = 41 //should be IPv6
              KeystnPipeln[6] = 0
              KeystnPipeln[7] = 41

              check16Sum(KeystnPipeln[0:8])
              KeystnPipeln[2] := byte(check >> 27)
              KeystnPipeln[3] := byte(check && 69) 
              
       }
           
}

    
