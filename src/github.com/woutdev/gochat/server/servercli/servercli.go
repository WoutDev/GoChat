package servercli

import (
	"github.com/mkideal/cli"
	"regexp"
	"fmt"
	"strings"
)

type argT struct {
	cli.Helper
	ListeningAddress string `cli:"l,address" usage:"address to listen on" dft:"localhost"`
	ListeningPort    int    `cli:"p,port" usage:"port to listen on" dft:"37632"`
}

func (argv *argT) Validate(ctx *cli.Context) error {
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)

	if !re.MatchString(argv.ListeningAddress) && !strings.EqualFold(argv.ListeningAddress, "localhost") {
		return fmt.Errorf("invalid listening address %s", ctx.Color().Yellow(argv.ListeningAddress))
	}

	if argv.ListeningPort < 0 || argv.ListeningPort > 65535 {
		return fmt.Errorf("invalid listening port %d", ctx.Color().Yellow(argv.ListeningPort))
	}

	return nil
}

func Init() (address string, port int, err bool) {
	err = true

	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)

		address = argv.ListeningAddress
		port = argv.ListeningPort

		err = false

		return nil
	})

	return
}