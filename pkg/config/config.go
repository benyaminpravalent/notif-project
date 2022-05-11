package config

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultType  = "json"
	configURLEnv = "CONFIG_URL"
)

var config *viper.Viper

//Load load configuration from config url, by default will load environment variable
func Load(def map[string]interface{}, path string) error {

	// first lets load .env file
	econf := viper.New()
	config = viper.New()
	aconf := viper.New()

	for k, v := range def {
		econf.BindEnv(k)
		config.SetDefault(k, v)
		if econf.IsSet(k) {
			config.Set(k, econf.Get(k))
		}
	}

	name := filepath.Base(path)
	t := strings.TrimPrefix(filepath.Ext(path), ".")
	path = filepath.Dir(path)
	aconf.SetConfigName(name)
	aconf.SetConfigType(t)
	aconf.AddConfigPath(path)
	if err := aconf.ReadInConfig(); err != nil {
		return err
	}

	for k := range def {
		if aconf.IsSet(k) {
			config.Set(k, aconf.Get(k))
		}
	}

	return nil
}

//Get get interface{}
func Get(k string) interface{} {
	return config.Get(k)
}

//GetString get string
func GetString(k string) string {
	return config.GetString(k)
}

//GetBool get bool
func GetBool(k string) bool {
	return config.GetBool(k)
}

//GetInt get int
func GetInt(k string) int {
	return config.GetInt(k)
}

//GetFloat64 get float64
func GetFloat64(k string) float64 {
	return config.GetFloat64(k)
}

//GetStringSlice get []string
func GetStringSlice(k string) []string {
	return config.GetStringSlice(k)
}

//GetStringMapString get map[string]string
func GetStringMapString(k string) map[string]string {
	return config.GetStringMapString(k)
}

// GetStringMap get map[string]interface{}
func GetStringMap(k string) map[string]interface{} {
	return config.GetStringMap(k)
}
