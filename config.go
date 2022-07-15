package dopplergoruntime

const URL = "https://%s@api.doppler.com/v3/configs/config/secrets/download?format=json&project=%s&config=%s"

type DopplerConfigYaml struct {
  Setup struct {
    Project string `yaml:"project"`
    Config string `yaml:"config"`
  } `yaml:"setup"`
}

type DopplerRuntimeConfig struct {
  Token string
  Project string
  Config string
  EnableDebug bool
}
