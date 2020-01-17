package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/config/source/flag"

	micro "github.com/micro/go-micro"
	proto "github.com/chat/greeter"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

type Host struct {
	Address string `json:"address"`
	Port int `json:"port"`
}

type Config struct {
	Hosts map[string]Host `json:"hosts"`
}

var (

	host Host
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)
	//conf:=config.NewConfig()

	//config.LoadFile("../config/database.json")

	//config.Scan(&conf)

	config.Load(
		env.NewSource(),
		flag.NewSource(),
		file.NewSource(
			file.WithPath("../config/database.json"),
		),
	)


	config.Get("hosts","database").Scan(&host)

	host11:=config.Get("hosts","database","address").String("localhost")
	port11:=config.Get("hosts","database","port").Int(3000)

	fmt.Println("=====================")
	fmt.Printf("%#v\n",host)
	fmt.Println(host11)
	fmt.Println(port11)
	//fmt.Printf("%#v\n",conf)
	fmt.Println("======================")





	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}