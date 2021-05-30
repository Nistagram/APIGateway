package globals

import "os"

const (
	advertisementMicroserviceHostDefault = "http://localhost"
	advertisementMicroservicePortDefault = "9091"
	contentMicroserviceHostDefault       = "http://localhost"
	contentMicroservicePortDefault       = "9092"
	locationMicroserviceHostDefault      = "http://localhost"
	locationMicroservicePortDefault      = "9093"
	messagingMicroserviceHostDefault     = "http://localhost"
	messagingMicroservicePortDefault     = "9094"
	usersMicroserviceHostDefault         = "http://localhost"
	usersMicroservicePortDefault         = "9095"
	serverPortDefault                    = "9090"
)

var (
	AdvertisementMicroserviceHost string = loadEnvValue("ADVERTISEMENT_HOST", advertisementMicroserviceHostDefault)
	AdvertisementMicroservicePort string = loadEnvValue("ADVERTISEMENT_PORT", advertisementMicroservicePortDefault)
	ContentMicroserviceHost       string = loadEnvValue("CONTENT_HOST", contentMicroserviceHostDefault)
	ContentMicroservicePort       string = loadEnvValue("CONTENT_PORT", contentMicroservicePortDefault)
	LocationMicroserviceHost      string = loadEnvValue("LOCATION_HOST", locationMicroserviceHostDefault)
	LocationMicroservicePort      string = loadEnvValue("LOCATION_PORT", locationMicroservicePortDefault)
	MessagingMicroserviceHost     string = loadEnvValue("MESSAGING_HOST", messagingMicroserviceHostDefault)
	MessagingMicroservicePort     string = loadEnvValue("MESSAGING_PORT", messagingMicroservicePortDefault)
	UsersMicroserviceHost         string = loadEnvValue("USERS_HOST", usersMicroserviceHostDefault)
	UsersMicroservicePort         string = loadEnvValue("USERS_PORT", usersMicroservicePortDefault)
	Port                          string = loadEnvValue("PORT", serverPortDefault)
)

func loadEnvValue(envName string, defaultValue string) string {
	if val, present := os.LookupEnv(envName); present {
		return val
	} else {
		return defaultValue
	}
}

func GetUsersMicroserviceUrl() string {
	if len(UsersMicroservicePort) == 0 {
		return UsersMicroserviceHost
	} else {
		return UsersMicroserviceHost + ":" + UsersMicroservicePort
	}
}
