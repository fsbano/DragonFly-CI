package main

import "config"
import "crypto/x509"
import "crypto/tls"
import "io"
import "io/ioutil"
import "fmt"
import "log"

func main() {

    c := config.LoadConfiguration()

    roots := x509.NewCertPool()
    data, err := ioutil.ReadFile(c.SSLCACertificateFile)
    if err != nil {
        log.Fatal("failed to parse root certificate")
    }

    roots.AppendCertsFromPEM(data)
    configRootCA := &tls.Config{RootCAs: roots}

    conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", c.Server, c.Port), configRootCA)
    if err != nil {
        log.Fatal(err)
    }

    io.WriteString(conn, "Hello")
    conn.Close()
}
