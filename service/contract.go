package service

import (
	"context"
	"fmt"
	"gitee.com/chatcode/chainrouter/stub"
	"gitee.com/chatcode/simstub/commons"
	"gitee.com/chatcode/simstub/contracts"
	"gitee.com/chatcode/simstub/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/go-hclog"
	"os"
	"time"
)

var checkInterval time.Duration = time.Second * 2
var startInterval time.Duration = time.Second * 1

type ContractManager struct {
	dataBase      *models.DataBase
	proxyAddress  common.Address
	bridgeAddress common.Address
	auth          *bind.TransactOpts
	client        *ethclient.Client
	logger        hclog.Logger
	config        *commons.Config
	stop          chan struct{}
}

func (this *ContractManager) Stop() {
	close(this.stop)
}

func (this *ContractManager) GetDataBase() *models.DataBase {
	return this.dataBase
}
func (this *ContractManager) GetProxyAddress() common.Address {
	return this.proxyAddress
}
func (this *ContractManager) GetBridgeAddress() common.Address {
	return this.bridgeAddress
}
func NewContractManager(client *ethclient.Client, auth *bind.TransactOpts, config *commons.Config) (*ContractManager, error) {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: false,
		Name:       "ContractManager",
		TimeFormat: commons.TimeFormat,
	})
	db, err := models.GetDBConnection(config.DBConfig)
	if err != nil {
		return nil, err
	}
	dataBase := models.NewDataBase(db)
	return &ContractManager{
		dataBase: dataBase,
		config:   config,
		auth:     auth,
		logger:   logger,
		client:   client,
		stop:     make(chan struct{}),
	}, nil
}

func (this *ContractManager) Run(f func()) {
	this.CheckContractIsReady()
	go this.ListenCrossChainRequest()
	go this.ListenRegisterCallbackResult()
	f()
}

func (this *ContractManager) CheckContractIsReady() {
	start := time.Now()
	defer func() {
		this.logger.Info("CheckContractIsReady exited")
		this.logger.Info(fmt.Sprintf("CheckContractIsReady time elapsed:%+v", time.Now().Sub(start)))
	}()
	//应用的是同一个账户进行部署，所以同步部署合约较为稳妥
	this.CheckProxy()
	this.CheckBridge()
	this.CheckAnchor(this.config.Common.Resource)
	for {
		err := this.CheckAnchorInit(this.config.Common.Resource)
		if err == nil {
			break
		}
	}
	for {
		err := this.CheckProxyWithPath(this.config.Common.Zone, this.config.Common.Chain, this.config.Common.Resource)
		if err == nil {
			break
		}
	}
	for {
		err := this.CheckAnchorUser(this.config.Common.Resource)
		if err == nil {
			break
		}
	}
	for {
		err := this.CheckTargetChainPath(this.config.Common.Resource, this.config.DefaultChain.Path, this.config.DefaultChain.Method, this.config.DefaultChain.CallbackPath, this.config.DefaultChain.CallbackMethod,this.config.DefaultChain.CrossChainType)
		if err == nil {
			break
		}
	}
}

func (this *ContractManager) CheckProxy() {
	start := time.Now()
	defer func() {
		this.logger.Info("CheckProxy exited")
		this.logger.Info(fmt.Sprintf("CheckProxy time elapsed:%+v", time.Now().Sub(start)))
	}()
	proxyExists := this.dataBase.ContractExists(stub.PROXY_NAME)
	if !proxyExists {
		//部署合约
		if this.client != nil && this.auth != nil {
			address, tx, _, err := contracts.DeployProxy(this.auth, this.client)
			if err != nil {
				this.logger.Error(err.Error())
				return
			}

			timer := time.NewTimer(startInterval)
			defer timer.Stop()
		label:
			for {
				select {
				case <-timer.C:
					receipt, err := this.client.TransactionReceipt(context.Background(), tx.Hash())
					if err == nil {
						if receipt != nil {
							if receipt.Status == 1 {
								obj := &models.Contract{
									Name:    stub.PROXY_NAME,
									Address: address.String(),
								}
								this.proxyAddress = address
								this.logger.Info(fmt.Sprintf("Proxy deploy success;address=%s,tx=%s", address.String(), tx.Hash().String()))
								_, err = this.dataBase.AddContract(obj)
								if err != nil {
									this.logger.Error(err.Error())
									return
								}
								break label
							} else if receipt.Status == 0 {
								this.logger.Error(fmt.Sprintf("Proxy deploy fail;tx=%s", tx.Hash().String()))
								break label
							}
						} else {
							timer.Reset(checkInterval)
						}
					} else {
						timer.Reset(checkInterval)
					}
				}
			}

		}
	} else {
		addr := this.dataBase.GetContractAddressByName(stub.PROXY_NAME)
		if addr == "" {
			panic("system exception")
		}
		this.proxyAddress = common.HexToAddress(addr)
	}
}
func (this *ContractManager) CheckBridge() {
	start := time.Now()
	defer func() {
		this.logger.Info("CheckBridge exited")
		this.logger.Info(fmt.Sprintf("CheckBridge time elapsed:%+v", time.Now().Sub(start)))
	}()
	bridgeExists := this.dataBase.ContractExists(stub.BRIDGE_NAME)
	if !bridgeExists {
		//部署合约
		if this.client != nil && this.auth != nil {
			address, tx, _, err := contracts.DeployBridge(this.auth, this.client)
			if err != nil {
				this.logger.Error(err.Error())
				return
			}
			timer := time.NewTimer(startInterval)
			defer timer.Stop()
		label:
			for {
				select {
				case <-timer.C:
					receipt, err := this.client.TransactionReceipt(context.Background(), tx.Hash())
					if err == nil {
						if receipt != nil {
							if receipt.Status == 1 {
								obj := &models.Contract{
									Name:    stub.BRIDGE_NAME,
									Address: address.String(),
								}
								this.bridgeAddress = address
								this.logger.Info(fmt.Sprintf("bridge deploy success;address=%s,tx=%s", address.String(), tx.Hash().String()))
								_, err = this.dataBase.AddContract(obj)
								if err != nil {
									this.logger.Error(err.Error())
									return
								}
								break label
							} else if receipt.Status == 0 {
								this.logger.Error(fmt.Sprintf("bridge deploy fail;tx=%s", tx.Hash().String()))
								break label
							}
						} else {
							timer.Reset(checkInterval)
						}
					} else {
						timer.Reset(checkInterval)
					}
				}
			}

		}
	} else {
		addr := this.dataBase.GetContractAddressByName(stub.BRIDGE_NAME)
		if addr == "" {
			panic("system exception")
		}
		this.bridgeAddress = common.HexToAddress(addr)
	}
}
func (this *ContractManager) CheckAnchor(resource string) {
	start := time.Now()
	defer func() {
		this.logger.Info("CheckAnchor exited")
		this.logger.Info(fmt.Sprintf("CheckAnchor time elapsed:%+v", time.Now().Sub(start)))
	}()
	bridgeExists := this.dataBase.ContractExists(resource)
	if !bridgeExists {
		//部署合约
		if this.client != nil && this.auth != nil {
			address, tx, _, err := contracts.DeployAnchorChain(this.auth, this.client)
			if err != nil {
				this.logger.Error(err.Error())
				return
			}
			timer := time.NewTimer(startInterval)
			defer timer.Stop()
		label:
			for {
				select {
				case <-timer.C:
					receipt, err := this.client.TransactionReceipt(context.Background(), tx.Hash())
					if err == nil {
						if receipt != nil {
							if receipt.Status == 1 {
								obj := &models.Contract{
									Name:    resource,
									Address: address.String(),
								}
								this.logger.Info(fmt.Sprintf("anchor deploy success;address=%s,tx=%s", address.String(), tx.Hash().String()))
								_, err = this.dataBase.AddContract(obj)
								if err != nil {
									this.logger.Error(err.Error())
									return
								}
								break label
							} else if receipt.Status == 0 {
								this.logger.Error(fmt.Sprintf("anchor deploy fail;tx=%s", tx.Hash().String()))
								break label
							}
						} else {
							timer.Reset(checkInterval)
						}
					} else {
						timer.Reset(checkInterval)
					}
				}
			}
		}
	}
}
func (this *ContractManager) CheckAnchorInit(resource string) error {
	start := time.Now()
	defer func() {
		this.logger.Info("CheckAnchorInit exited")
		this.logger.Info(fmt.Sprintf("CheckAnchorInit time elapsed:%+v", time.Now().Sub(start)))
	}()
	bridgeExists := this.dataBase.ContractExists(stub.BRIDGE_NAME)

	anchorExists := this.dataBase.ContractExists(resource)

	if bridgeExists && anchorExists {

		bridgeAddress := this.dataBase.GetContractAddressByName(stub.BRIDGE_NAME)

		anchorAddress := this.dataBase.GetContractAddressByName(resource)

		if this.dataBase.ContractInitReady() {
			return nil
		}
		anchorContract, err := contracts.NewAnchorChain(common.HexToAddress(anchorAddress), this.client)

		if err != nil {
			return err
		}
		this.logger.Info("anchorAddress=" + anchorAddress)
		this.logger.Info("bridgeAddress=" + bridgeAddress)
		tx, err := anchorContract.Init(this.auth, common.HexToAddress(bridgeAddress))
		if err != nil {
			return err
		}
		timer := time.NewTimer(startInterval)
		defer timer.Stop()
	label:
		for {
			select {
			case <-timer.C:
				receipt, err := this.client.TransactionReceipt(context.Background(), tx.Hash())
				if err == nil {
					if receipt != nil {
						if receipt.Status == 1 {
							obj := &models.ContractInitialization{
								AnchorAddress: anchorAddress,
								BridgeAddress: bridgeAddress,
								TxHash:        tx.Hash().String(),
							}
							this.logger.Info(fmt.Sprintf("anchor init success;tx=%s", tx.Hash().String()))
							_, err = this.dataBase.AddContractInit(obj)
							if err != nil {
								return err
							}
							break label
						} else if receipt.Status == 0 {
							this.logger.Error(fmt.Sprintf("anchor init fail;tx=%s", tx.Hash().String()))
							break label
						}
					} else {
						timer.Reset(checkInterval)
					}
				} else {
					timer.Reset(checkInterval)
				}
			}
		}
	}
	return nil
}
func (this *ContractManager) CheckProxyWithPath(zone string, chain string, resource string) error {
	start := time.Now()
	defer func() {
		this.logger.Info("CheckProxyWithPath exited")
		this.logger.Info(fmt.Sprintf("CheckProxyWithPath time elapsed:%+v", time.Now().Sub(start)))
	}()
	bridgeExists := this.dataBase.ContractExists(stub.BRIDGE_NAME)
	anchorExists := this.dataBase.ContractExists(resource)
	proxyExists := this.dataBase.ContractExists(stub.PROXY_NAME)
	if bridgeExists && anchorExists && proxyExists {
		bridgeAddress := this.dataBase.GetContractAddressByName(stub.BRIDGE_NAME)
		anchorAddress := this.dataBase.GetContractAddressByName(resource)
		proxyAddress := this.dataBase.GetContractAddressByName(stub.PROXY_NAME)
		proxyAddr := common.HexToAddress(proxyAddress)
		anchorAddr := common.HexToAddress(anchorAddress)
		bridgeAddr := common.HexToAddress(bridgeAddress)

		proxy, err := contracts.NewProxy(proxyAddr, this.client)
		if err != nil {
			return err
		}

		resourceAnchorPath := fmt.Sprintf("%s.%s.%s", zone, chain, resource)

		if !this.dataBase.PathInitReady(proxyAddress, anchorAddress, resourceAnchorPath) {

			tx, err := proxy.AddPath(this.auth, resourceAnchorPath, anchorAddr, contracts.AnchorChainABI)
			if err != nil {
				return err
			}
			timer := time.NewTimer(startInterval)
			defer timer.Stop()
		first:
			for {
				select {
				case <-timer.C:
					receipt, err := this.client.TransactionReceipt(context.Background(), tx.Hash())
					if err == nil {
						if receipt != nil {
							if receipt.Status == 1 {
								obj := &models.PathInitialization{
									ProxyAddress: proxyAddress,
									PathAddress:  anchorAddress,
									Path:         resourceAnchorPath,
									TxHash:       tx.Hash().String(),
								}
								this.logger.Info(fmt.Sprintf("proxy AddPath Anchor success;tx=%s", tx.Hash().String()))
								_, err = this.dataBase.AddPathInit(obj)
								if err != nil {
									return err
								}
								break first
							} else if receipt.Status == 0 {
								this.logger.Error(fmt.Sprintf("proxy AddPath Anchor fail;tx=%s", tx.Hash().String()))
								break first
							}
						} else {
							timer.Reset(checkInterval)
						}
					} else {
						timer.Reset(checkInterval)
					}
				}
			}
		}
		resourceBridgePath := fmt.Sprintf("%s.%s.%s", zone, chain, stub.BRIDGE_NAME)

		if !this.dataBase.PathInitReady(proxyAddress, bridgeAddress, resourceBridgePath) {
			tx, err := proxy.AddPath(this.auth, resourceBridgePath, bridgeAddr, contracts.BridgeABI)
			if err != nil {
				return err
			}
			timer := time.NewTimer(startInterval)
			defer timer.Stop()
		label:
			for {
				select {
				case <-timer.C:
					receipt, err := this.client.TransactionReceipt(context.Background(), tx.Hash())
					if err == nil {
						if receipt != nil {
							if receipt.Status == 1 {
								this.logger.Info(fmt.Sprintf("proxy AddPath bridge success;tx=%s", tx.Hash().String()))
								obj := &models.PathInitialization{
									ProxyAddress: proxyAddress,
									PathAddress:  bridgeAddress,
									Path:         resourceBridgePath,
									TxHash:       tx.Hash().String(),
								}
								_, err := this.dataBase.AddPathInit(obj)
								if err != nil {
									return err
								}
								break label
							} else if receipt.Status == 0 {
								this.logger.Error(fmt.Sprintf("proxy AddPath bridge fail;tx=%s", tx.Hash().String()))
								break label
							}
						} else {
							timer.Reset(checkInterval)
						}
					} else {
						timer.Reset(checkInterval)
					}
				}
			}
		}
	}
	return nil
}
