package handlers

import (
	"fmt"
	"goinaction/k8sapp/version"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/takama/bit"
)

type Status struct {
	Host     string   `json:"host"`
	Version  string   `json:"version"`
	Commit   string   `json:"commit"`
	Repo     string   `json:"repo"`
	Compiler string   `json:"compiler"`
	Runtime  Runtime  `json:"runtime"`
	State    State    `json:"state"`
	Requests Requests `json:"requests"`
}

type Runtime struct {
	CPU        int    `json:"cpu"`
	Memory     string `json:"memory"`
	Goroutines int    `json:"goroutines"`
}

type State struct {
	Maintenance bool   `json:"maintenance"`
	Uptime      string `json:"uptime"`
}

type Requests struct {
	Duration Duration `json:"duration"`
	Codes    Codes    `json:"codes"`
}

type Duration struct {
	Average string `json:"average"`
	Max     string `json:"max"`
}

type Codes struct {
	C2xx int `json:"2xx"`
	C4xx int `json:"4xx"`
	C5xx int `json:"5xx"`
}

func (h *Handler) Info(c bit.Control) {
	host, _ := os.Hostname()
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)

	c.Code(http.StatusOK)
	c.Body(Status{
		Host:     host,
		Version:  version.RELEASE,
		Commit:   version.COMMIT,
		Repo:     version.REPO,
		Compiler: runtime.Version(),
		Runtime: Runtime{
			CPU:        runtime.NumCPU(),
			Memory:     fmt.Sprintf("%.2fMB", float64(m.Sys)/(1<<(10*2))),
			Goroutines: runtime.NumGoroutine(),
		},
		State: State{
			Maintenance: h.maintenance,
			Uptime:      time.Now().Sub(h.stats.startTime).String(),
		},
		Requests: Requests{
			Duration: Duration{
				Average: h.stats.requests.Duration.Average,
				Max:     h.stats.requests.Duration.Max,
			},
			Codes: Codes{
				C2xx: h.stats.requests.Codes.C2xx,
				C4xx: h.stats.requests.Codes.C4xx,
				C5xx: h.stats.requests.Codes.C5xx,
			},
		},
	})
}
