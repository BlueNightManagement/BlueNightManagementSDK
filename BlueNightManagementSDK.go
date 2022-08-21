package BlueNightManagementSDK

import (
	"github.com/gorilla/mux"
)

type PlugInMetadataVer1 struct {
	Name string
	Version string
	License string
	Author string
	Date string
	Description string
	URL string
	RouterPrefix string
}

type PlugInVer1 interface {
	Init(router *mux.Router)
	Info() *PlugInMetadataVer1
}
