package main

import "github.com/fsbano/DragonFly/src/config"
import "crypto/tls"
import "fmt"
import "io"
import "net"
import "os"
import "log"

func main() {

    c := config.LoadConfiguration()
    cert, err := tls.LoadX509KeyPair(c.SSLCertificateFile,
                                     c.SSLCertificateKeyFile)
    if err != nil {
        log.Fatal("Error loading certificate. ", err)
    }

    tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}}

    listener, err := tls.Listen("tcp4", "0.0.0.0:8001", tlsCfg)
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    for {
       log.Println("Waiting for clients")
       conn, err := listener.Accept()
       if err != nil {
           log.Fatal(err)
       }
       go func(c net.Conn) {
         io.Copy(os.Stdout, c)
         fmt.Println()
         c.Close()
       }(conn)
    }
}
