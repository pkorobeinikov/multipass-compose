package command

import (
	"github.com/urfave/cli/v2"

	"github.com/pkorobeinikov/multipass-compose/internal/cfg"
	"github.com/pkorobeinikov/multipass-compose/internal/multipass"
	"github.com/pkorobeinikov/multipass-compose/internal/spec"
)

func Status(ctx *cli.Context) error {
	s, err := spec.Load(cfg.DefaultMultipassComposeSpecFile)
	if err != nil {
		return err
	}

	for name := range s.Services {
		if err := multipass.Execute(ctx.Context, "info", name); err != nil {
			return err
		}
	}

	return nil
}
