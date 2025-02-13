package main

import("fmt"
       "io")
/*#################################################################################################################################################################################################################################################

				network														NETWORK

*/############################################################################################################################################################################################################################################################

type PnI interface{

    IP IP
    IPv4 string
    IPv6 string
    port string//trying to put the port and IP together
	func JoinHostPort([],port string) string//this makes addr that can be sent to other port'd machine

}

type earConfig struct{//I forgot wat this for
    //addr/something else,listen fo da lo to drop
    KeepAliveConfig KeepAliveConfig //this can be ignored if you setIt.Enable(false)//idk
    Nosy (Ear *ListenConfig) Listen(ctx context.Context, "IP","49,7,10,15")(Nosy, error)//idfk
    IP IP //set it to Ip might changeto TCP if I figure packets.
    Mask IPMask // think this mask it
    const(IPv4len = 4, IPv6 = 16)//syntax mite whoop me
       
}

type Operator struct{//basically Resolver

    preferGo bool
    strictErrors bool
    dial func(ctx context.Context, network, address string)(Conn,error)//wen called should make TCP/UDP calls
    
}
	
//this just figs ima need for the intranet
//these are lessons for me and who ever configing this
//*******************************BYTES**********************************************************************************BYTES****************************************************************************************************BYTES
//just diffrent bytes flows
type buffer interface{//inna a way I can make this a recursive formula... thinking this the bufferFlow skel

    read(bf [][]byte)(e int err error)//should treat buffer like args.([0:n])
    write(bf [][]byte)(e int err error)
  
  
}
//****************This is a more concrete way to actuall read and write the data in this method******************************************************************************************************************************************************************************
type Mybuffer struct{

    data [][]byte
    
}

func (b *MyBuffer) Read(p [][]byte) (n int, err error) {
    // Implementation for reading into p
    return Mybuffer.data()
}

func (b *MyBuffer) Write(p [][]byte) (n int, err error) {
    // Implementation for writing from p
    return Mybuffer.data()
}

type oilPipeln interface{//this supposed to be a data writeer

    
    
}


type OilPipelin struct{


    
}
//************************************std(standard inputoutputerr methods)*******************************************************************************************************************************
type pipes struct{

    stdinn io.Reader//shouts out to the state but im finna trick on it
    stdout io.Writer//I think interface would be quick
    stderr io.Error
    
}

type std interface{//this supposed to be an addative of the pipes struct.

    read(in []byte)(err error)
    write(out []byte)(err error)
    Error(Err []byte)(err error)
    
}

/************************************TECHIE***********************************************************************TECHIE*******************************************************************************************************************************************/

type hardwareAddr []byte
func (a hardwareAddr) String() string {
	if len(a) == 0 {//when 0 not nil return an sttring with no input well call not string
		return ""//this this nil/none we think of
	}
	buf := make([]byte, 0, len(a)*3-1)//the math mathical limit of whatever 2**64 is think its default limit of comps now-adays
	for i, b := range a {//i in b make b range of a
		if i > 0 {//i can not be less that 0 thats an over flow
			buf = append(buf, ':')//this is building a string out
		}
		buf = append(buf, hexDigit[b>>4]) //I think this breaks it down
		buf = append(buf, hexDigit[b&0xF])//this give the starting number shape best way to put it
	}
	return string(buf)//keeps a string
}



/*
//I been looking for this for an hour but there is a short hand
//there's a 
//var hardwareaddr := HwAdr u get the gist

type HardwareAddr []byte

func (a HardwareAddr) String() string {
	if len(a) == 0 {//when 0 not nil return an sttring with no input well call not string
		return ""//this this nil/none we think of
	}
	buf := make([]byte, 0, len(a)*3-1)//the math mathical limit of whatever 2**64 is think its default limit of comps now-adays
	for i, b := range a {//i in b make b range of a
		if i > 0 {//i can not be less that 0 thats an over flow
			buf = append(buf, ':')//this is building a string out
		}
		buf = append(buf, hexDigit[b>>4]) //I think this breaks it down
		buf = append(buf, hexDigit[b&0xF])//this give the starting number shape best way to put it
	}
	return string(buf)//keeps a string
}
