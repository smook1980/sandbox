package target

import (
	"context"
	"net/http"
	"os/exec"
	"sync"

	"log"
)

type TargetRule struct {
	Host string
	Path string
}

type TargetConfig struct {
	PassEnv    bool
	CustomEnvs map[string]string
	TargetURI  string
	Cwd        string
	Cmd        string
	Args       []string
	Rules      []TargetRule
}

type Target struct {
	ctx    context.Context
	config *TargetConfig

	lock   sync.Mutex
	worker *exec.Cmd
}

func New(tc *TargetConfig, ctx context.Context) *Target {
	return &Target{ctx: ctx, config: tc}
}

func (t *Target) Spawn() {
	log.Println("Spawn()")
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.worker != nil {
		err := t.worker.Process.Kill()
		if err != nil {
			log.Printf("ERROR: %s", err)
		}
	}

	t.worker = exec.CommandContext(t.ctx, t.config.Cmd, t.config.Args...)
	t.worker.Start()
}

func (t *Target) Handles(r *http.Request) bool {
	return false
}

func (t *Target) Process(res http.ResponseWriter, req *http.Request) {

}
