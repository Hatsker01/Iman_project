package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct{
	Environment string
	
	DataServiceHost string
	DataServicePort int
	PostSericeHost string
	PostSericePort int

	//context time in second
	CtxTimeout int

	LogLevel string
	HTTPPort string
}

func Load() Config{
	c:=Config{}

	c.Environment=cast.ToString(getOrReturnDefault("ENVIRONMENT","develop"))


	c.LogLevel=cast.ToString(getOrReturnDefault("LOG_LEVEL","debug"))
	c.HTTPPort=cast.ToString(getOrReturnDefault("HTTP_PORT",":8090"))
	c.DataServiceHost=cast.ToString(getOrReturnDefault("GET_SERVICE_HOST","localhost"))
	c.DataServicePort=cast.ToInt(getOrReturnDefault("GET_SERVICE_PORT",9000))

	c.PostSericeHost = cast.ToString(getOrReturnDefault("DATA_SERVICE_HOST", "localhost"))
	c.PostSericePort = cast.ToInt(getOrReturnDefault("DATA_SERVICE_PORT", 8000))
	

	c.CtxTimeout=cast.ToInt(getOrReturnDefault("CTX_TIMEOUT",7))
	return c

}

func getOrReturnDefault(key string, defaultValue interface{}) interface{}{
	_,exists:=os.LookupEnv(key)
	if exists{
		return os.Getenv(key)
	}
	return defaultValue
}