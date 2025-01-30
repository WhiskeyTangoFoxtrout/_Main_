package main

import("fmt"
       "os")

//feel like ima need something for writing file stuff

//USE THIS WHEN DIGGIN OIL

type File struct{//borrowing from go but i wanna beef it

    *file
  
}

//random thought what does raw data look like when u send it back and forth

type oil interface{

    read(oF [][]byte)(e int err error)
    write(oF [][]byte)(e int err error)
    oil [][]byte *Oil
    
}

//*********************************************************************this controls flow/

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

    stdinn io.Reader//shouts out to the state but im finna trick on it
    stdout io.Writer
    stderr io.Error
    written int    
}

//************************************************************

type scrambler interface{

    read(scr [][]byte)(e int err error)
    write(scr [][]byte)(e int err error)
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
    scrambler() *src []byte
    
    
}
    
