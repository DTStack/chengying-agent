package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"

	"easyagent/internal"
	"easyagent/internal/proto"
	"easyagent/internal/sidecar/base"
	"easyagent/internal/sidecar/client"
	"easyagent/internal/sidecar/config"
	"easyagent/internal/sidecar/controller"
	"easyagent/internal/sidecar/event"
	"easyagent/internal/sidecar/monitor"
	"easyagent/internal/sidecar/register"
	"github.com/kardianos/service"
	"github.com/urfave/cli"
)

var winStop = make(chan struct{}, 1)

type WinSvc struct {
	ctl *controller.Controller
}

func (svc WinSvc) Start(s service.Service) error {
	base.Infof("windows easyAgent start")
	return nil
}
func (svc WinSvc) Stop(s service.Service) error {
	base.Infof("windows easyAgent stop")

	exitGracefully(svc.ctl)
	winStop <- struct{}{}

	return nil
}

func main() {
	fmt.Println("EasyAgent Client " + internal.VERSION)
	fmt.Println("Copyright (c) 2017 DTStack Inc.")

	app := cli.NewApp()
	app.Version = internal.VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Usage: "config path",
		},
		cli.StringFlag{
			Name:  "agents,ags",
			Usage: "agents path to load/store",
			Value: "agents-file.yml",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "debug info",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		base.SetDebug(ctx.Bool("debug"))

		cfg, err := config.ParseConfig(ctx.String("config"))
		if err != nil {
			return err
		}
		log := cfg.Log
		if err := base.ConfigureLogger(log.Dir, log.MaxSize, log.MaxBackups, log.MaxAge); err != nil {
			return err
		} else {
			fmt.Printf("saving logs at %s\n", log.Dir)
		}

		rpc := cfg.Rpc
		serverAddr := net.JoinHostPort(rpc.Server, strconv.Itoa(rpc.Port))
		eaClient, err := client.NewEasyAgentClient(cfg.EasyAgent.Uuid.UUID, serverAddr, rpc.Tls, rpc.TlsSkipVerify, rpc.CertFile)
		if err != nil {
			base.Errorf("%v", err)
			return err
		}

		agents, err := config.ParseAgents(ctx.String("agents"))
		if err != nil {
			return err
		}
		ctl := controller.NewController(eaClient, agents, ctx.String("agents"))

		if runtime.GOOS == "windows" {
			svc, err := service.New(WinSvc{ctl}, &service.Config{Name: "easyAgent"})
			if err != nil {
				base.Errorf("init windows service error: %v", err)
				return err
			}
			go svc.Run()
		}

		event.SetEventDefaultClient(eaClient)
		monitor.SetMonitorInterval(cfg.EasyAgent.MonitorInterval)
		register.RegisterSidecar(eaClient, ctl, cfg.CallBack)
		if err = base.MountCgroup(); err != nil {
			event.ReportEvent(&proto.Event_AgentError{Errstr: err.Error()})
		}
		ctl.Run()
		monitor.StartMonitSystem()

		if runtime.GOOS == "windows" {
			<-winStop
			return nil
		}

		signalCapture()
		exitGracefully(ctl)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exit with failure: %v\n", err)
	}
}

func signalCapture() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case sig := <-signals:
		err := fmt.Errorf("quit according to signal '%s'\n", sig.String())
		base.Errorf("%v", err)
		event.ReportEvent(&proto.Event_AgentError{Errstr: err.Error()})
	}
}

func exitGracefully(ctl *controller.Controller) {
	base.DeleteTcDev()
	ctl.KillAgents()
}
