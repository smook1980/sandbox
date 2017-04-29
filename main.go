package main

import (
	"log"
	"os/exec"

	"github.com/smook1980/sandbox/app"
	"github.com/smook1980/sandbox/target"
)

func main() {
	app.Boot(func(c *app.SandboxConfig) {
		cmd, err := exec.LookPath("say")
		if err != nil {
			log.Printf("say command not found? %s", err)
		}

		t := &target.TargetConfig{
			Cmd:  cmd,
			Args: []string{"Hot", "Dog"},
		}

		c.Targets = append(c.Targets, t)
	})
}
