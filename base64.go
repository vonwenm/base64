package main

import (
  "gopkg.in/alecthomas/kingpin.v2"
  "encoding/base64"
  "io/ioutil"
  "fmt"
  "os"
)

var (
  config      = kingpin.New("b64", "Base64 Encode/Decoder  v 1.0")
  encode     = config.Command("encode", "Encode an input file.")
  encodeFile = encode.Arg("srcfile", "Name of file to encode").Required().String()
  decode     = config.Command("decode", "Decode an input file.")
  decodeFile  = decode.Arg("srcfile", "Name of file to decode").Required().String()
  targetFile  = decode.Arg("tgtfile", "Name of the target file").Required().String()
)

func doEncode (srcFile string){
  srcDat, err1 := ioutil.ReadFile( srcFile )
  if err1 != nil {
    fmt.Println("ERROR: ",err1)
  }
  xEnc := base64.StdEncoding.EncodeToString(srcDat)
  err2 := ioutil.WriteFile(srcFile + ".b64",  []byte(xEnc), 0644)
  if err2 != nil {
    fmt.Println("ERROR: ",err2)
  }

}

func doDecode (srcFile string, tgtFile string,){
  srcDat, err1 := ioutil.ReadFile( srcFile )
  if err1 != nil {
    fmt.Println("ERROR: ",err1)
  }
  srcStr := string(srcDat)
  xDec, _ := base64.StdEncoding.DecodeString(srcStr)
  err2 := ioutil.WriteFile(tgtFile,  []byte(xDec), 0644)
  if err2 != nil {
    fmt.Println("ERROR: ",err2)
  }

}


func main() {
  fmt.Println( "** START RUN *************************************" )
  switch kingpin.MustParse( config.Parse(os.Args[1:]) ) {
    case encode.FullCommand():
      fmt.Println( "File to encode: " + *encodeFile )
      doEncode ( *encodeFile )
    case decode.FullCommand():
      fmt.Println( "File to decode: " + *decodeFile )
            doDecode ( *decodeFile, *targetFile )
  }
  fmt.Println( "** END RUN *************************************" )







/*
    outf,_ := os.Create("verify.xlsx")

    xDec, _ := base64.StdEncoding.DecodeString(xEnc)

    for i, v := range xDec {

      if v == 0 {
        fmt.Println("Pos: ",i)
        fmt.Println("Found Zero: ",v)
        err := binary.Write(outf, binary.LittleEndian, v)
        fmt.Print(v)
        if err != nil {
          fmt.Println("err!",err)
        }
      } else
      {
        err := binary.Write(outf, binary.LittleEndian, v)
        fmt.Print(v)
        if err != nil {
          fmt.Println("err!",err)
        }
      }

    }


    //err2 := ioutil.WriteFile("verify.xlsx", xDec, 0644)
    //if err2 != nil {
    //}

    // Go supports both standard and URL-compatible
    // base64. Here's how to encode using the standard
    // encoder. The encoder requires a `[]byte` so we
    // cast our `string` to that type.
    //sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    //fmt.Println(sEnc)

    // Decoding may return an error, which you can check
    // if you don't already know the input to be
    // well-formed.
    //sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    //fmt.Println(string(sDec))
    */
}
