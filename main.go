package main

import (
	"flag"
	"fmt"
	"gitee.com/chatcode/chainrouter/stub"
	"gitee.com/chatcode/simstub/commons"
	"gitee.com/chatcode/simstub/service"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"os"
	"time"
)

var config *string = flag.String("config", "./stub.yaml", "")

func main() {
	usePlugin()
	//client()
	//select {
	//
	//}
}
func usePlugin() {
	//一定要记得加
	flag.Parse()
	c, err := commons.LoadConfig(*config)
	if err != nil {
		fmt.Println("LoadConfig error:", err)
		return
	}
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
		Name:       c.Common.Name,
	})
	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BASIC_PLUGIN",
		MagicCookieValue: "hello",
	}

	manager, err := service.NewManager(c)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	stubServer := service.NewStubServer(manager)

	var pluginMap = map[string]plugin.Plugin{
		"stub": &stub.StubPlugin{Impl: stubServer},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}

func client() {
	//一定要记得加
	flag.Parse()
	c, err := commons.LoadConfig(*config)
	if err != nil {
		fmt.Println("LoadConfig error:", err)
		return
	}
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: false,
		TimeFormat: commons.TimeFormat,
	})
	manager, err := service.NewManager(c)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	stubServer := service.NewStubServer(manager)

	//sources := stubServer.GetResources()

	//for _, v := range sources {
	//	logger.Debug(fmt.Sprintf("%+v", v))
	//}
	//request := stub.TransactionRequest{}
	//request.Method = "connect"
	//args := []string{"gaogao", "meimei"}
	//request.Args = []interface{}{args}
	//ctx := stub.TransactionContext{}
	//path := stub.Path{}
	//path.Resource = "evidence"
	//ctx.Path = path
	//request.TransactionContext = ctx
	//request.Options=make(map[string]interface{})
	//request.Options[stub.TRANSACTION_UNIQUE_ID]=fmt.Sprintf("%d",time.Now().UnixNano())
	//
	//stubServer.SendTransaction(request)

	request := stub.TransactionRequest{}
	request.Method = "scanBlocks"
	request.Args = []interface{}{}
	ctx := stub.TransactionContext{}
	path := stub.Path{}
	path.Resource = "evidence"
	ctx.Path = path
	request.TransactionContext = ctx
	request.Options = make(map[string]interface{})
	request.OutputType = "[]string"
	request.Options[stub.TRANSACTION_UNIQUE_ID] = fmt.Sprintf("%d", time.Now().UnixNano())

	res := stubServer.Call(request)
	fmt.Printf("res=%+v\n", res)
}

//flag.Parse()
//c, err := commons.LoadConfig(*config)
//if err != nil {
//	fmt.Println("LoadConfig error:",err)
//	return
//}
//manager, err := service.NewManager(c)
//if err != nil {
//	fmt.Println(err)
//	return
//}
//stubServer := service.NewStubServer(manager)
//resources:=stubServer.GetResources()
//fmt.Println("resources=",len(resources))
//for _,v:=range resources{
//	fmt.Printf("resources=%+v",v)
//}
//flag.Parse()
//c, err := commons.LoadConfig(*config)
//if err != nil {
//	fmt.Println(err)
//	return
//}
//fmt.Println("host:", c.Node.Host)
//ethclient, err := ethclient.Dial("http://192.168.4.248:8545")
//
//conn := service.NewSimpleConnection(ethclient)
//
//conn.SetProxyAddress(common.HexToAddress("0xE9E8Dc18E8D5de2BCc15604B406dD678124355ea"))
//
//server := service.NewStubServer(conn)
//
////server.ListPath()
////server.GetAbiByName("evidence")
//
//request := stub.TransactionRequest{}
//
//request.Method = "getIncrement"
//
//request.Args = nil
//
//context1 := stub.TransactionContext{}
//
//path := &stub.Path{}
//
////judicial.sim.bridge
//path.Zone = "judicial"
//path.Chain = "sim"
//path.Resource = "bridge"
//
//context1.Path = path
//
//request.TransactionContext = context1
//
//response := server.AsyncCall(request)
//
//fmt.Printf("getIncrement=%+v", response)
//
//request.Method = "getInterchainRequests"
//
//hello := big.NewInt(2)
//
//request.Args = []interface{}{hello}
//
//response = server.AsyncCall(request)
//
//fmt.Printf("getInterchainRequests=%+v", response)
