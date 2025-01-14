//This is a server you would use to talk to my server
//ima eventually move it to another branch
//Note to self this is a http server
package main

import("fmt"
       "log"
       "net"
       "net/http"
       "io/buff"
       "Configs/__ManeHosCon__"//I think this valid
      )
//***********************************Global Variables*******************************************************************
//cgo only use 500 system threads which is nice
export GODEBUG=netdns=go + 1

//I need some sockets in the future
var nativeEdian binary.ByteOrder
//***********************************************************************SOCKETS*******************************************

//screw raw sockets! I gotta go read the docs but I got an interview brb

//**********************************************************************THE ACTUAL CONNECT PROGRAM*************************

func main(){

    loConfig := GPS
    earConfig := Ear

    goto ear()
    
}

//no idea how to display the URL... LMAO TYPE NUMBAS LMAOOOO

//****************************************************************Const and types******************************************

type IP []byte

/*type IPAddr struct{

    IP IP
    IPv4 string
       
} dont think this is needed still keeping it... congfig.go
*/
type Page Struct{

    Title string
    Body []byte
    
}

//*****************************************************************Server-to-Saver*************************************************

//ima need a save func but idk what ima save just yet. I just want it to echo for rn

//****************************************************ACtual Engine and raw connect****************************************

func ear(){//Ok so open port and have port listen for request and respond back with an "WE'RE LIVE FROM DA UNDAGROUND"
     
    ln, err := net.Listen("tcp",ipv4)
       if err != nil{

              fmt.Fprintf("ErrorCode: 404")
              os.Exit(1)
              
       }

       fmt.Fprintf(conn, "SHould be http /HTTP/1.0\r\n\r\n")
       status,err := bufio.NewReader(GPS).ReadString("\n")
       go handleConnection(gps)
       
}

func Lo(){//ima have to remember this location

    powerDialer := net.IPv6(49,7,10,15)
    ipv6 := net.IPv6(49,7,10,15)//I think this the real ip
    I6Ski:= GPS.IPMask(ipv6)//think this valid.
    func JoinHostPort([I6Ski], 52 string) string
       
}

