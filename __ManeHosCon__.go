//ima try and make this a config file for __HostMane__\
//Felt like it would be better to import this like a package... Mane starting ta drag it
//still figuring how the line up and work off each other. This config file wit the client config file
//where ever I see Resl
package main

import("fmt"
       "net"
       "time")

type earConfig struct{//I forgot wat this for
    //addr/something else,listen fo da lo to drop
    KeepAliveConfig KeepAliveConfig //this can be ignored if you setIt.Enable(false)//idk
    Nosy (Ear *ListenConfig) Listen(ctx context.Context, "IP","49,7,10,15")(Nosy, error)//idfk
    IP IP //set it to Ip might changeto TCP if I figure packets.
    Mask IPMask // think this mask it
    const(IPv4len = 4, IPv6 = 16)//syntax mite whoop me
       
}

type loConfig interface{//think this valid see what happens when I pipeline da yadig
//this is for conn/ kinda wanna see if ican loConfig.control()
    control func(ip string, p syscall.RawConn) error
    read (lo []byte)(ear int, err error)//I need these arguements ima try to have ear listen for a port/ lo be da IP
    write(lo []byte)(ear int, err error) 
    Close () error
    LocalAddr() lo
    RemoteAddr() locate
    Timeout (c time.Time) error 
    readTimeout (c time.Time) error
    hardwareAddr net.HardwareAddr
       
}

type Operator struct{//basically Resolver

    preferGo bool
    strictErrors bool
    dial func(ctx context.Context, network, address string)(Conn,error)//wen called should make TCP/UDP calls
    
}

type SwitchBoard struct{//part dialer meant to complement the Op.

    timeout time.ParseDuration("1m30s")//if Operator.Timeout ever get called
    deadline time.Time
    localAddr Addr
    dualStack bool //helps if IPv4 or the other path not working
    fallBackDelay time.ParseDuration("3m")//actually what checks dual stack path
    keepAliveConfig KeepAliveConfig
    operator *Operator
    Cancel <-chan struct{}//if you could cancel a channel this would
    Control func(network, address string, c syscall.RawConn) error //
    ControlContext func(ctx context.Context, network, address string, c syscall.RawConn) error
    
}

type KeepAliveConfig struct{

    enable bool
    idle time.ParseDuration("15s")//how long it would sit in idel
    interval time.ParseDuration("20s")//btw each call
    count int//how many probes
    
}

