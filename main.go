package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

func main() {
	app := cli.NewApp()
	app.Name = "etcd-lock-port"
	app.Usage = "establish a lock on a port using etcd"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "key, k",
			EnvVar: "ETCD_LOCK_PORT_KEY",
			Usage:  "etcd key to put the registered port for reverse lookup",
		},
		cli.StringFlag{
			Name:   "name, n",
			EnvVar: "ETCD_LOCK_PORT_NAME",
			Usage:  "Name of the service to register the port to",
		},
		cli.StringFlag{
			Name:   "registry, r",
			EnvVar: "ETCD_LOCK_PORT_REGISTRY",
			Usage:  "Directory where all locked ports are registered in etcd",
		},
	}

	app.Run(os.Args)
}

func run(context *cli.Context) {
	key := context.String("key")
	name := context.String("name")
	registry := context.String("registry")

	if key == "" || name == "" || registry == "" {
		cli.ShowAppHelp(context)
		color.Red("  --key, --name, and --registry are all required")
		os.Exit(1)
	}

	etcdLockPort, err := New(name, registry, key)
	if err != nil {
		log.Panicf("Error connecting to etcd: %v", err.Error)
	}

	port, err := etcdLockPort.LockPort()
	if err != nil {
		log.Panicf("Error establishing lock: %v", err.Error)
	}
	fmt.Printf("Locked: %v", port)
}
