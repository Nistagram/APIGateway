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

//func GetUsersMicroserviceUrl() string {
//	if len(UsersMicroservicePort) == 0 || UsersMicroservicePort == "80" {
//		return "http://" + UsersMicroserviceHost
//		// return UsersMicroserviceHost
//	} else {
//		return "http://" + UsersMicroserviceHost + ":" + UsersMicroservicePort
//		// return UsersMicroserviceHost + ":" + UsersMicroservicePort
//	}
//}

func GetUsersMicroserviceUrl() string {
	if len(UsersMicroservicePort) == 0 || UsersMicroservicePort == "80" {
		if UsersMicroserviceHost == "http://localhost"{
			return UsersMicroserviceHost
		}else {
			return "http://" + UsersMicroserviceHost
		}
	} else {
		if UsersMicroserviceHost == "http://localhost"{
			return UsersMicroserviceHost + ":" + UsersMicroservicePort
		}else{
			return "http://" + UsersMicroserviceHost + ":" + UsersMicroservicePort
		}
	}
}

func GetContentMicroserviceUrl() string {
	if len(ContentMicroservicePort) == 0 || ContentMicroservicePort == "80" {
		if ContentMicroserviceHost == "http://localhost"{
			return ContentMicroserviceHost
		}else{
			return "http://" + ContentMicroserviceHost
		}
		// return ContentMicroserviceHost
	} else {
		if ContentMicroserviceHost == "http://localhost"{
			return ContentMicroserviceHost + ":" + ContentMicroservicePort
		}else{
			return "http://" + ContentMicroserviceHost + ":" + ContentMicroservicePort
		}
		// return ContentMicroserviceHost + ":" + ContentMicroservicePort
	}
}
