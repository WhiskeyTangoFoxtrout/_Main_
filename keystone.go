package main
//pipeln really for the socket its so its can be more automative and opt filled. still need to set up an GUI
import (
        "fmt"
        "sys"
        "io"
        "os"
        "string")


//*****************************I need to make a parser to read the input and look for port numbers
type roughneck interface{

    nil// still not good at parser
    
}

//**********************************actual I/o stream that moves the data around
type keyStone interface{

    func flow(oil byte[]){

        return Meter( oil.stdin() byte[])
        
    }
    
}

//*********************************************************************this controls flow

type Vale interface{//want this to be more like vale.stdio... sys.stdwrite...sys.stdflush()
    // thx go!
    func Copy(rcr io.Writer(), src io.Reader()) (written int32, err error){
    
        return copyBuffer(rcr, src, nil)

  }

    func copyBuffer(rcr io.Writer(), src io.Reader() twin []byte)(written int32, err error){

        return twin byte[0:rcr]
    
}
    
//************************************************************I just need an oil stucture

type pipes struct{

    stdin io.Reader//shouts out to the state but im finna trick on it
    stdout io.Writer
    stderr io.Error
    written int    
}
