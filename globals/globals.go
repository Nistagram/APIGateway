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
	jwtTokenSecretDefault                = "12345"
	jwtTokenLifeLength                   = "24"
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
	JwtTokenSecret                string = loadEnvValue("JWT_SECRET", jwtTokenSecretDefault)
	JwtTokenLifeLength            string = loadEnvValue("JWT_TOKEN_LIFE_LENGTH", jwtTokenLifeLength) //Hours
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
		if UsersMicroserviceHost == "http://localhost" {
			return UsersMicroserviceHost
		} else {
			return "http://" + UsersMicroserviceHost
		}
	} else {
		if UsersMicroserviceHost == "http://localhost" {
			return UsersMicroserviceHost + ":" + UsersMicroservicePort
		} else {
			return "http://" + UsersMicroserviceHost + ":" + UsersMicroservicePort
		}
	}
}

func GetContentMicroserviceUrl() string {
	if len(ContentMicroservicePort) == 0 || ContentMicroservicePort == "80" {
		if ContentMicroserviceHost == "http://localhost" {
			return ContentMicroserviceHost
		} else {
			return "http://" + ContentMicroserviceHost
		}
		// return ContentMicroserviceHost
	} else {
		if ContentMicroserviceHost == "http://localhost" {
			return ContentMicroserviceHost + ":" + ContentMicroservicePort
		} else {
			return "http://" + ContentMicroserviceHost + ":" + ContentMicroservicePort
		}
		// return ContentMicroserviceHost + ":" + ContentMicroservicePort
	}
}

func GetAdvertisementMicroserviceUrl() string {
	if len(AdvertisementMicroservicePort) == 0 || AdvertisementMicroservicePort == "80" {
		if AdvertisementMicroserviceHost == "http://localhost" {
			return AdvertisementMicroserviceHost
		} else {
			return "http://" + AdvertisementMicroserviceHost
		}
		// return ContentMicroserviceHost
	} else {
		if AdvertisementMicroserviceHost == "http://localhost" {
			return AdvertisementMicroserviceHost + ":" + AdvertisementMicroservicePort
		} else {
			return "http://" + AdvertisementMicroserviceHost + ":" + AdvertisementMicroservicePort
		}
		// return ContentMicroserviceHost + ":" + ContentMicroservicePort
	}
}
