type JohnLo interface{

    Read(IP []byte)(n int, err error)
    Write(IP []byte)(n int, err error)
    Close() error
    Lo() Addr
    remoteLo() Addr
    setDeadline(t time.Time) error
    setWriteDeadline(t time.Time) error
  
}
