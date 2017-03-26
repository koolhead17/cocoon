package server

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/ncodes/cocoon/core/common"
	"github.com/ncodes/cocoon/core/connector/launcher"
	"github.com/ncodes/cocoon/core/connector/server/proto"
	stub_proto "github.com/ncodes/cocoon/core/runtime/golang/proto"
	logging "github.com/op/go-logging"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var log = logging.MustGetLogger("connector.api")

// APIServer defines a grpc server for
// invoking operations against cocoon code
type APIServer struct {
	server   *grpc.Server
	endedCh  chan bool
	launcher *launcher.Launcher
}

// NewAPIServer creates a new grpc API server
func NewAPIServer(launcher *launcher.Launcher) *APIServer {
	server := new(APIServer)
	server.launcher = launcher
	return server
}

// Start starts the API service
func (api *APIServer) Start(addr string, endedCh chan bool) {

	api.endedCh = endedCh

	lis, err := net.Listen("tcp", fmt.Sprintf("%s", addr))
	if err != nil {
		log.Fatalf("failed to listen on port=%s. Err: %s", strings.Split(addr, ":")[1], err)
	}

	time.AfterFunc(2*time.Second, func() {
		log.Infof("Started GRPC API server on port %s", strings.Split(addr, ":")[1])
	})

	api.server = grpc.NewServer()
	proto.RegisterAPIServer(api.server, api)
	api.server.Serve(lis)
}

// Stop stops the orderer and returns an exit capie.
func (api *APIServer) Stop(exitCode int) int {
	api.server.Stop()
	close(api.endedCh)
	return exitCode
}

// Invoke calls a function in the cocoon code.
func (api *APIServer) Invoke(ctx context.Context, req *proto.InvokeRequest) (*proto.InvokeResponse, error) {
	log.Infof("New invoke transaction (%s)", req.GetId())

	var respCh = make(chan *stub_proto.Tx)
	var txID = req.GetId()
	err := api.launcher.GetClient().SendTx(&stub_proto.Tx{
		Id:     txID,
		Invoke: true,
		Name:   "function",
		Params: append([]string{req.GetFunction()}, req.GetParams()...),
	}, respCh)

	if err != nil {
		log.Debugf("Failed to send transaction [%s] to cocoon code. %s", txID, err)
		return nil, err
	}

	resp, err := common.AwaitTxChan(respCh)
	if err != nil {
		return nil, err
	}

	return &proto.InvokeResponse{
		Id:       txID,
		Function: req.GetFunction(),
		Body:     resp.GetBody(),
	}, nil
}
