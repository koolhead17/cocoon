package languages

import (
	"os"
	"os/user"
	"path"
	"strings"

	"fmt"

	"github.com/ellcrys/cocoon/core/types"
	"github.com/ellcrys/util"
	"github.com/goware/urlx"
	"github.com/ncodes/modo"
)

// SupportedVendorTool includes a list of supported vendor packaging tool
var SupportedVendorTool = []string{
	"glide",
	"govendor",
}

// Go defines a deployment helper for go cocoon.
type Go struct {
	spec        *types.Spec
	name        string
	image       string
	userHome    string
	imgGoPath   string
	repoURL     string
	env         map[string]string
	buildParams map[string]interface{}
}

// NewGo returns a new instance a golang
// cocoon code deployment helper.
func NewGo(spec *types.Spec) *Go {
	g := &Go{
		name:      "go",
		image:     "ellcrys/launch-go",
		userHome:  "/home",
		imgGoPath: "/go",
		spec:      spec,
		env:       map[string]string{},
	}
	g.setUserHomeDir()
	return g
}

// GetName returns the name for the language
func (g *Go) GetName() string {
	return g.name
}

// GetImage returns the image name to deploy cocoon code on
func (g *Go) GetImage() string {
	return g.image
}

// setUserHomeDir sets the user home directory
func (g *Go) setUserHomeDir() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	g.userHome = usr.HomeDir
	return nil
}

// SetRunEnv adds environment variables to be added when the run script is constructed.
func (g *Go) SetRunEnv(env map[string]string) {
	for k, v := range env {
		g.env[k] = v
	}
}

// GetDownloadDir returns the location to save
// the downloaded go cocoon code source on the connector
func (g *Go) GetDownloadDir() string {
	u, _ := urlx.Parse(g.spec.URL)
	repoID := strings.Trim(u.Path, "/")
	return path.Join(g.userHome, "/ccode/sources/", g.spec.ID, repoID)
}

// GetCopyDir returns the location in
// the container where the source code will be copied to
func (g *Go) GetCopyDir() string {
	u, _ := urlx.Parse(g.spec.URL)
	repoID := strings.Trim(u.Path, "/")
	return path.Join(g.imgGoPath, "src/github.com/", strings.Split(repoID, "/")[0])
}

// GetSourceDir returns the root directory of the cocoon source code
// in container.
func (g *Go) GetSourceDir() string {
	u, _ := urlx.Parse(g.spec.URL)
	repoID := strings.Trim(u.Path, "/")
	return path.Join(g.imgGoPath, "src/github.com/", repoID)
}

// RequiresBuild returns true if cocoon codes written in
// go language spec a build process.
// During development, If DEV_RUN_ROOT_BIN env is set, it will return false as
// the run command will find and find the ccode binary in the repo root.
func (g *Go) RequiresBuild() bool {
	if runRootBin := os.Getenv("DEV_RUN_ROOT_BIN"); len(runRootBin) > 0 && runRootBin == "true" {
		return false
	}
	return true
}

// SetBuildParams sets and validates build parameters
func (g *Go) SetBuildParams(buildParams map[string]interface{}) error {
	if pkgMgr := buildParams["pkgMgr"]; pkgMgr != nil {
		if pkgMgrVal, ok := pkgMgr.(string); ok {
			if !util.InStringSlice(SupportedVendorTool, pkgMgrVal) {
				return fmt.Errorf("invalid `pkgMgr` value in build script")
			}
		} else {
			return fmt.Errorf("invalid type for `pkgMgr` paremeter, expected string")
		}
	}
	g.buildParams = buildParams
	return nil
}

// GetBuildCommand will return the command to build
// the source code and generate a binary.
func (g *Go) GetBuildCommand() *modo.Do {

	cmds := []string{
		"cd " + g.GetSourceDir(),
	}

	// add the command to run package manager vendor operation
	if len(g.buildParams) > 0 {
		if pkgMgr := g.buildParams["pkgMgr"]; pkgMgr != nil {
			switch pkgMgr {
			case "glide":
				cmds = append(cmds, "glide install")
			case "govendor":
				cmds = append(cmds, "govendor fetch -v +out")
			}
		}
	}

	// build source
	cmds = append(cmds, "go build -v -o /bin/ccode")

	return &modo.Do{
		Cmd: []string{
			"bash",
			"-c",
			strings.Join(cmds, "&&"),
		},
	}
}

// GetRunCommand returns the command to start the
// cocoon code for this language. If DEV_RUN_ROOT_BIN env is set, it will run any
// ccode binary located in the mount destination.
func (g *Go) GetRunCommand() *modo.Do {
	cmds := []string{}

	// add environment variables
	for k, v := range g.env {
		cmds = append(cmds, fmt.Sprintf("export %s='%s'", k, v))
	}

	// if DEV_RUN_ROOT_BIN is set (for development only), run 'ccode' binary that is expected to be in source root
	if runRootBin := os.Getenv("DEV_RUN_ROOT_BIN"); len(runRootBin) > 0 && runRootBin == "true" {
		cmds = append(cmds, "cd "+g.GetSourceDir()) // move into source root director
		cmds = append(cmds, "./ccode")
		return &modo.Do{
			Cmd: []string{
				"bash",
				"-c",
				strings.Join(cmds, "&&"),
			},
		}
	}
	cmds = append(cmds, "ccode")
	return &modo.Do{
		Cmd: []string{
			"bash",
			"-c",
			strings.Join(cmds, "&&"),
		},
	}
}
