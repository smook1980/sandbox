package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/smook1980/sandbox/target"
)

type Configuration func(*SandboxConfig)

type SandboxConfig struct {
	Listen  string
	Targets []*target.TargetConfig
}

type Sandbox struct {
	config   *SandboxConfig
	targets  []*target.Target
	context  context.Context
	cancleFn context.CancelFunc
}

func NewSandbox() *Sandbox {
	ctx, cfn := context.WithCancel(context.Background())
	return &Sandbox{
		config:   defaultConfig(),
		context:  ctx,
		cancleFn: cfn,
	}
}

func Boot(cfs ...Configuration) {
	sbx := NewSandbox()
	for _, cf := range cfs {
		cf(sbx.config)
	}

	for _, tc := range sbx.config.Targets {
		sbx.makeTarget(tc)
	}

	sbx.Start()
}

func (s *Sandbox) makeTarget(tc *target.TargetConfig) {
	s.targets = append(s.targets, target.New(tc, s.context))
}

func (s *Sandbox) ServeHTTP(rep http.ResponseWriter, req *http.Request) {
	fmt.Println("Req: %+v", req)
	fmt.Println("Hello world.")
	for _, t := range s.targets {
		t.Spawn()
	}
}

func (s *Sandbox) Start() {
	http.ListenAndServe(":8080", s)
}

func defaultConfig() *SandboxConfig {
	return &SandboxConfig{}
}
