type Johnreach interface{//****add DNS look up and a way to communicate to Login Page, this a conn

    Read(IP []byte)(n int, err error)//transfer
    Write(IP []byte)(n int, err error)//transfer
    Close() error//ends conn
    Lo() Addr //mo litteral
    remoteLo() Addr //this b mo of reference
    Timeout (c time.Time) error 
    readTimeout (c time.Time) error
    
}

//I need to figure a way to get the IpandSocket sent to Dewalt

//need a config for the addr
type PnI interface{

    type PoIp struct{

    IP IP
    IPv4 string
    IPv6 string
    port string//trying to put the port and IP together
    func JoinHostPort([],port string) string//this makes addr that can be sent to other port'd machine
        
    }
    
}
