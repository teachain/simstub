package main

import (
	"fmt"
	"gitee.com/chatcode/simstub/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func test1() {
	ethclient, err := ethclient.Dial("ws://192.168.3.92:8546")
	if err != nil {
		fmt.Println(err)
		return
	}
	address := common.HexToAddress("0x789E1f7e4A776Cf0479C9703aa05420fB2169Af9")

	proxyContract, err := contracts.NewProxy(address, ethclient)
	if err != nil {
		fmt.Println(err)
		return
	}
	sink := make(chan *contracts.ProxyOnSendTransaction, 1)
	filter, err := proxyContract.WatchOnSendTransaction(&bind.WatchOpts{}, sink)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filter.Unsubscribe()
	for {
		select {
		case e := <-sink:
			fmt.Println(e.Name)
			fmt.Println(e.ContractAddress)
			fmt.Println(e.ArgsWithMethodId)
		case err := <-filter.Err():
			fmt.Println("filter", err)

		}
	}
}

func main() {

	go listenAnchor()

	ethclient, err := ethclient.Dial("ws://192.168.3.92:8546")
	if err != nil {
		fmt.Println(err)
		return
	}
	address := common.HexToAddress("0xcb661AAC16cA1173D4F20eb9D995877B1d371EBD")

	bridgeContract, err := contracts.NewBridge(address, ethclient)
	if err != nil {
		fmt.Println(err)
		return
	}
	//opts *bind.WatchOpts, sink chan<- *AnchorChainOncallContract
	//proxyContract.WatchOncallContract()
	sink := make(chan *contracts.BridgeOnHandleRequest, 1)
	var start uint64 = 40000
	opts := &bind.WatchOpts{
		Start: &start,
	}
	filter, err := bridgeContract.WatchOnHandleRequest(opts, sink)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filter.Unsubscribe()
	for {
		select {
		case e := <-sink:
			fmt.Println("bridge e.Raw.TxHash=", e.Raw.TxHash.String())
			fmt.Println("bridge  e.Raw.address=", e.Raw.Address.String())
			fmt.Println("bridge  e.Raw.BlockHash=", e.Raw.BlockHash.String())
			fmt.Println("bridge  e.Raw.BlockNumber=", e.Raw.BlockNumber)
			fmt.Println(e.Request)
		case err := <-filter.Err():
			fmt.Println("filter", err)

		}
	}
}
func listenAnchor() {
	ethclient, err := ethclient.Dial("ws://192.168.3.92:8546")
	if err != nil {
		fmt.Println(err)
		return
	}
	address := common.HexToAddress("0x718Cea95F036B365Fd7fF07e800E51e29F79199f")

	proxyContract, err := contracts.NewAnchorChain(address, ethclient)
	if err != nil {
		fmt.Println(err)
		return
	}
	//opts *bind.WatchOpts, sink chan<- *AnchorChainOncallContract
	//proxyContract.WatchOncallContract()
	sink := make(chan *contracts.AnchorChainOncallContract, 1)
	var start uint64 = 40000
	opts := &bind.WatchOpts{
		Start: &start,
	}
	filter, err := proxyContract.WatchOncallContract(opts, sink)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filter.Unsubscribe()
	for {
		select {
		case e := <-sink:
			fmt.Println("anchor e.Raw.TxHash=", e.Raw.TxHash.String())
			fmt.Println("anchor  e.Raw.address=", e.Raw.Address.String())
			fmt.Println("anchor  e.Raw.BlockHash=", e.Raw.BlockHash.String())
			fmt.Println("anchor  e.Raw.BlockNumber=", e.Raw.BlockNumber)
			fmt.Println(e.Result)
		case err := <-filter.Err():
			fmt.Println("filter", err)

		}
	}
}
