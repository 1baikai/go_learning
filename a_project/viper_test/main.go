package main

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	InstallPkg map[string]interface{} `yaml:"install_pkg" mapstructure:"install_pkg"`

}

func main() {
	//t := map[string]interface{}{}
	//buffer, err := ioutil.ReadFile("/Users/baikai/GoProject/src/baikai/leetcode/a_project/6/config.yaml")
	//err = yaml.Unmarshal(buffer, &t)
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//fmt.Printf("%v",t)
	InitConfig()

	a:=GetPkgName()
	fmt.Printf("%v",a)



	omPkg := strings.Replace("openGauss-2.0.0-openEuler_arm-64bit-all.tar.gz", "all", "om", -1)
	publisherAndPlatform:=strings.Split(omPkg,"-")[2]
	publisher := strings.Split(publisherAndPlatform,"_")[0]
	//platform :=strings.Split(publisherAndPlatform,"_")[1]
	omPkg = strings.Replace(omPkg,publisherAndPlatform,publisher,-1)
	logrus.Println(omPkg)
}
var cfg *Config
func GetPkgName() map[string]interface{} { return cfg.GetPkgName() }
func (c *Config) GetPkgName() map[string]interface{} {
	return c.InstallPkg
}

func InitConfig() error {
	conf, err := LoadConfigFile("config")
	if err != nil {
		logrus.Errorf("LoadConfigFile err %v", err)
		return err
	}
	logrus.Info("load config success")
	cfg = conf
	return nil
}

func LoadConfigFile(configFileName string) (*Config, error) {
	var config Config
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(".")
	viper.NewWithOptions(viper.KeyDelimiter("::"))
	viper.SetDefault("chart::values", map[string]interface{}{
		"ingress": map[string]interface{}{
			"annotations": map[string]interface{}{
				"traefik.frontend.rule.type":                 "PathPrefix",
				"traefik.ingress.kubernetes.io/ssl-redirect": "true",
			},
		},
	})

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("read config error: %s", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config error: %s", err)
	}
	return &config, nil
}