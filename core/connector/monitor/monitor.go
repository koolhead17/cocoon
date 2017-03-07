package monitor

import (
	"fmt"
	"time"

	docker "github.com/ncodes/go-dockerclient"
	"github.com/olebedev/emitter"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("launcher.monitor")

// HandleFunc is the expected function signature
type HandleFunc func(map[string]interface{})

// Report represents the result of the monitors checks
type Report struct {
	DiskUsage int64
}

// Monitor defines a launcher monitor module checking resource
// useage of a cocoon code. This module provides a pubsub feature that allows
// other external modules to subscribe to events from it and to also emit events to
// the module.
type Monitor struct {
	emitter     *emitter.Emitter
	containerID string
	stop        bool
	dckClient   *docker.Client
}

// NewMonitor creates a new monitor instance.
func NewMonitor() *Monitor {
	e := emitter.New(10)
	return &Monitor{
		emitter: e,
	}
}

// SetContainerID sets the id of the container to monitor
func (m *Monitor) SetContainerID(cID string) {
	m.containerID = cID
}

// SetDockerClient sets the docker client
func (m *Monitor) SetDockerClient(dckClient *docker.Client) {
	m.dckClient = dckClient
}

// GetEmitter returns the monitor's emitter
func (m *Monitor) GetEmitter() *emitter.Emitter {
	return m.emitter
}

// Stop the monitor
func (m *Monitor) Stop() {
	m.stop = true
	m.emitter.Off("*")
}

// getContainerRootSize fetches the total
// size of all the files in the container.
func (m *Monitor) getContainerRootSize() (int64, error) {
	containers, err := m.dckClient.ListContainers(docker.ListContainersOptions{
		All:   true,
		Size:  true,
		Limit: 1,
		Filters: map[string][]string{
			"id": []string{m.containerID},
		},
	})
	if err != nil {
		return 0, fmt.Errorf("failed to list containers. %s", err)
	}

	return containers[0].SizeRw, nil
}

// Monitor starts the monitor
func (m *Monitor) Monitor() {
	for !m.stop {

		size, err := m.getContainerRootSize()
		if err != nil {
			log.Error(err.Error())
		}

		report := Report{
			DiskUsage: size,
		}

		<-m.emitter.Emit("monitor.report", report)
		time.Sleep(1 * time.Second)
	}
}
