package config

import "encoding/json"
import "os"
import "fmt"

type CfgDragonfly struct {
  SSLCACertificateFile string `json:"SSLCACertificateFile"`
  SSLCertificateFile string `json:"SSLCertificateFile"`
  SSLCertificateKeyFile string `json:"SSLCertificateKeyFile"`
  Server string `json:"Server"`
  Port int `json:"Port"`
}

func LoadConfiguration() CfgDragonfly {
  var config CfgDragonfly
  cfgFile, err := os.Open("/etc/dragonfly/config.json")
  if err != nil {
    fmt.Println(err)
  }
  defer cfgFile.Close()

  jsonParser := json.NewDecoder(cfgFile)
  jsonParser.Decode(&config)

  return config
}
