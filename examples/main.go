package main

import (
	"os"

	msalgo "github.com/AzureAD/microsoft-authentication-library-for-go/src/msal"
)

const port = "3000"

var config = createConfig("config.json")
var publicClientApp *msalgo.PublicClientApplication
var err error
var authCodeParams *msalgo.AcquireTokenAuthCodeParameters
var cacheAccessor = &SampleCacheAccessor{"examples/serialized_cache.json"}

func main() {
	exampleType := os.Args[1]
	if exampleType == "1" {
		acquireTokenDeviceCode()
	} else if exampleType == "2" {
		acquireByAuthorizationCodePublic()
	} else if exampleType == "3" {
		acquireByUsernamePasswordPublic()
	} else if exampleType == "4" {
		acquireByAuthorizationCodeConfidential()
	}
}
