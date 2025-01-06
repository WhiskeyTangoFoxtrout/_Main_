//ima try and make this a config file for __HostMane__\
//Felt like it would be better to import this like a package... Mane starting ta drag it
//still figuring how the line up and work off each other. This config file wit the client config file
package main

import("fmt"
       "net")

type earConfig struct{

    control func(ip string, p syscall.RawConn) error
    KeepAliveConfig KeepAliveConfig //this can be ignored if you setIt.Enable(false)//idk
    Nosy (Ear *ListenConfig) Listen(ctx context.Context, "IP","49,7,10,15")(Nosy, error)//idfk
    IP IP //set it to Ip might changeto TCP if I figure packets.
    Mask IPMask // think this mask it
       
}

type loConfig interface{//think this valid see what happens when I pipeline da yadig

    read (lo []byte)(ear int, err error)
    write(lo []byte)(ear int, err error)
    Close () error
    LocalAddr() lo
    RemoteAddr() locate
    Timeout (c time.Time) error
    readTimeout (c time.Time) error
       
}
