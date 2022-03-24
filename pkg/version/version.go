package version

import (
	"net/http"

	"github.com/pumphouse-p/peek-go/pkg/apiutils"
)

var VERSION = "JENKINS"

type VersionStatus struct {
	Version string `json:"version"`
}

type Version struct{}

func New() *Version {
	return &Version{}
}

func (v *Version) APIGet(w http.ResponseWriter, r *http.Request) {
	status := VersionStatus{}

	status.Version = VERSION

	apiutils.ServeJSON(w, status)
}
