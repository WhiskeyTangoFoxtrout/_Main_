type PnI interface{//bascially how I make an dns for the network

    type PoIp struct{

    IP IP
    IPv4 string
    IPv6 string
    port string//trying to put the port and IP together
    func JoinHostPort([],port string) string//this makes addr that can be sent to other port'd machine
        
    }

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



type pipes struct{

    stdinn io.Reader//shouts out to the state but im finna trick on it
    stdout io.Writer//I think interface would be quick
    stderr io.Error
    
}

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
