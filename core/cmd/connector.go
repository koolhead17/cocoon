package cmd

import (
	"fmt"
	"os"

	"github.com/ncodes/cocoon/core/config"
	"github.com/ncodes/cocoon/core/connector/launcher"
	"github.com/ncodes/cocoon/core/connector/server"
	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
)

func init() {
	config.ConfigureLogger()
}

// creates a deployment request with argument
// fetched from the environment.
func getRequest() (*launcher.Request, error) {

	// get cocoon code github link and language
	ccID := os.Getenv("COCOON_ID")
	ccURL := os.Getenv("COCOON_CODE_URL")
	ccTag := os.Getenv("COCOON_CODE_TAG")
	ccLang := os.Getenv("COCOON_CODE_LANG")
	buildParam := os.Getenv("COCOON_BUILD_PARAMS")

	if ccID == "" {
		return nil, fmt.Errorf("Cocoon code id not set @ $COCOON_ID")
	} else if ccURL == "" {
		return nil, fmt.Errorf("Cocoon code url not set @ $COCOON_CODE_URL")
	} else if ccLang == "" {
		return nil, fmt.Errorf("Cocoon code url not set @ $COCOON_CODE_LANG")
	}

	return &launcher.Request{
		ID:          ccID,
		URL:         ccURL,
		Tag:         ccTag,
		Lang:        ccLang,
		BuildParams: buildParam,
	}, nil
}

// connectorCmd represents the connector command
var connectorCmd = &cobra.Command{
	Use:   "connector",
	Short: "Start the connector",
	Long:  `Starts the connector and launches a cocoon code.`,
	Run: func(cmd *cobra.Command, args []string) {

		launchFailedCh := make(chan bool, 1)

		var log = logging.MustGetLogger("connector")
		log.Info("Connector started. Initiating cocoon code launch procedure.")

		// get request
		req, err := getRequest()
		if err != nil {
			log.Error(err.Error())
			return
		}

		// install cooncode
		lchr := launcher.NewLauncher(launchFailedCh)
		lchr.AddLanguage(launcher.NewGo())
		go lchr.Launch(req)

		// start grpc API server
		apiServer := server.NewAPIServer()
		go apiServer.Start(fmt.Sprintf(":%s", "8002"), make(chan bool, 1))

		if <-launchFailedCh {
			lchr.Stop()
			apiServer.Stop(1)
			log.Fatal("aborting: cocoon code launch failed")
		}

		log.Info("cocoon code launch successfully exited")
	},
}

func init() {
	RootCmd.AddCommand(connectorCmd)
}
