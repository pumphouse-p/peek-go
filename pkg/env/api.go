package env

import (
	"net/http"
	"os"
	"strings"

	"github.com/pumphouse-p/peek-go/pkg/apiutils"
)

type EnvStatus struct {
	CommandLine []string          `json:"commandLine"`
	Env         map[string]string `json:"env"`
}

type Env struct{}

func New() *Env {
	return &Env{}
}

func (e *Env) APIGet(w http.ResponseWriter, r *http.Request) {
	status := EnvStatus{}

	status.CommandLine = os.Args

	status.Env = map[string]string{}
	for _, e := range os.Environ() {
		splits := strings.SplitN(e, "=", 2)
		k, v := splits[0], splits[1]
		status.Env[k] = v
	}

	apiutils.ServeJSON(w, status)
}
