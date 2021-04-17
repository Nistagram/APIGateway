package globals

import "os"

const(
	usersMicroserviceHostDefault = "http://localhost"
	usersMicroservicePortDefault = "9095"
	serverPortDefault = "9090"
)

var (
	UsersMicroserviceHost string = loadEnvValue("UsersMicroserviceHost", usersMicroserviceHostDefault)
	UsersMicroservicePort string = loadEnvValue("UsersMicroservicePort", usersMicroservicePortDefault)
	ServerPort string = loadEnvValue("ServerPort", serverPortDefault)
	)

func loadEnvValue(envName string, defaultValue string) string {
	if val, present := os.LookupEnv(envName); present {
		return val
	} else {
		return defaultValue
	}
}

func GetUsersMicroserviceUrl() string{
	if len(UsersMicroservicePort) == 0 {
		return UsersMicroserviceHost
	} else{
		return UsersMicroserviceHost + ":" + UsersMicroservicePort
	}
}