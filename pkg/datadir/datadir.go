package datadir

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/rancher/wrangler/pkg/resolvehome"
)

var (
	DefaultDataDir     = "/var/lib/coname/" + "projname"
	DefaultHomeDataDir = "${HOME}/.coname/" + "projname"
	HomeConfig         = "${HOME}/.kube/" + "projname" + ".yaml"
	GlobalConfig       = "/etc/coname/" + "projname" + "/" + "projname" + ".yaml"
)

func Resolve(dataDir string) (string, error) {
	return LocalHome(dataDir, false)
}

func LocalHome(dataDir string, forceLocal bool) (string, error) {
	if dataDir == "" {
		if os.Getuid() == 0 && !forceLocal {
			dataDir = DefaultDataDir
		} else {
			dataDir = DefaultHomeDataDir
		}
	}

	dataDir, err := resolvehome.Resolve(dataDir)
	if err != nil {
		return "", errors.Wrapf(err, "resolving %s", dataDir)
	}

	return filepath.Abs(dataDir)
}
