package utl

import "os"

var (
	applicationName    = os.Getenv("APP_NAME")
	appicationVersion  = os.Getenv("APP_VERSION")
	applicationRelease = os.Getenv("APP_RELEASE")
)
