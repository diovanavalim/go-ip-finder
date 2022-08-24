package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func GenerateApp() *cli.App {
	app := cli.NewApp()

	app.Name = "Command Line Application"
	app.Usage = "Find IPs and Server Names over the Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find an IP by its DNS",
			Flags:  flags,
			Action: findIpByDns,
		},
		{
			Name:   "server",
			Usage:  "Find an IP servers by its DNS",
			Flags:  flags,
			Action: findServerByDns,
		},
	}

	return app
}

func findIpByDns(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServerByDns(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
