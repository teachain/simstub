package service

import (
	"gitee.com/chatcode/chainrouter/stub"
)

type StubServer struct {
	manager *Manager
}

func NewStubServer(manager *Manager) *StubServer {
	return &StubServer{
		manager: manager,
	}
}
func (this *StubServer) GetResources() []stub.ResourceInfo {
	return this.manager.GetResources()
}
func (this *StubServer) Call(request stub.TransactionRequest) stub.TransactionResponse {
	return this.manager.CallByProxy(request)
}
func (this *StubServer) SendTransaction(request stub.TransactionRequest) stub.TransactionResponse {
	return this.manager.SendTransactionByProxy(request)
}
func (this *StubServer) GetCrossTransactionInformationPage(request stub.CrossTransactionInformationPageRequest) stub.CrossTransactionInformationPage {
	return this.manager.GetCrossTransactionInformationPage(request)
}
func (this *StubServer) GetCrossChainTransaction(request stub.CrossChainTransactionRequest) stub.CrossChainTransaction {
	return this.manager.GetCrossChainTransaction(request)
}



