package main

import("fmt"
       "bytes"
       "os"
       "io"
       "net"
       )

func main(){

    nil
    
}

func checkError(err error){

    if err != nil{
        fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
        os.Exit(1)
    }
}

func check16Sum(msg []byte/*workin on the math here...*/){

    sum := 0//its gone be pretty fixed right now
//this made for 2 bytes
    //lets kep it even bc pwer of 2 will save my life and think i gotta convert odd numbers
    for := 0; n < len(msg); n += 2{
        sum += int(msg[n]) * 128 + int(msg[n+1])
    }
    sum = (sum >> 16) + (sum & 0xffff)
    sum += (sum >> 16)
    var answer uint16 = uint16(^sum)//16**
    return answer
}

check32Sum(msg []byte){

    sum:=0
    for := 0; n < len(msg); n +=2{

        sum += int(msg[n]) * 256 + int(msg[n+1])        
       }
    sum = (sum >> 32) + (sum & 0xffff)
    sum += (sum >> 16)
           var answer uint32 = uint16(^sum) 
           return answer
}



/*

type Block interface {
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
}

type Stream interface {
    XORKeyStream(dst, src []byte)
}

been needing a skeleton*/

//keep this in mind...the page i be yea
Identification,err := hardwareid.ID()//think we coo
hardwareAddr net.HardwareAddr

*/
