package dopplergoruntime

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type DopplerRuntime struct {
  Token string
  Project string
  Config string
  EnableDebug bool

  Result map[string]string
  CommonResponse *CommonResponse
}

func (dr *DopplerRuntime) DownloadSecret() ([]byte, error) {
  var (
    body []byte
    err error
  )

  url := fmt.Sprintf(URL, dr.Token, dr.Project , dr.Config)
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return body, err;
  }

  req.Header.Add("Accept", "application/json")

  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return body, err;
  }

  defer res.Body.Close()
  body, err = ioutil.ReadAll(res.Body)
  if err != nil {
    return body, err;
  }

  return body, err;

}

func (dr *DopplerRuntime) Parse(jsonByte []byte) error {
  return json.Unmarshal(jsonByte, &dr.Result)
}

func (dr *DopplerRuntime) ParseCommon(jsonByte []byte) error {
  return json.Unmarshal(jsonByte, &dr.CommonResponse)
}

func (dr *DopplerRuntime) SetEnv() {
  for key, val := range(dr.Result) {
    os.Setenv(key, val);
  }
}

func (dr *DopplerRuntime) Load() error {
  res, err := dr.DownloadSecret()
  if err != nil {
    return err
  }

  if err = dr.Parse(res); err != nil {
    if err = dr.ParseCommon(res); err != nil {

    }

    if dr.CommonResponse.Success == false {
      return errors.New(dr.CommonResponse.Messages[0])
    }
  }

  dr.SetEnv()


  if dr.EnableDebug {
    fmt.Println(fmt.Sprintf("%v Env Loaded; Project %s; Config %s", len(dr.Result), dr.Project, dr.Config))
  }

  return nil
}

func loadDopplerYaml (filename string)  (dopplerConfig *DopplerConfigYaml) {
  yamlFile, err := ioutil.ReadFile(filename)
  if err == nil {
    err = yaml.Unmarshal(yamlFile, &dopplerConfig)
  }

  return dopplerConfig
}


func NewDopplerRuntime(opt DopplerRuntimeConfig) *DopplerRuntime {

  var (
    token string
    defaultConfig string = "doppler.yaml"
    project string
    config string
  )


  if opt.Token == "" {
    token = os.Getenv("DOPPLER_TOKEN")
  } else {
    token = opt.Token
  }


  dopplerConfigYaml := loadDopplerYaml(defaultConfig)

  if opt.Project == "" {
    project = dopplerConfigYaml.Setup.Project
  } else {
    project = opt.Project
  }

  if opt.Config == "" {
    config = dopplerConfigYaml.Setup.Config 
  } else {
    config = opt.Config
  }

  return &DopplerRuntime{
    Token: token,
    Project: project,
    Config: config,
    EnableDebug: opt.EnableDebug,
  }
}


