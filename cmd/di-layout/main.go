package main

import (
	"log"

	"flamingo.me/dingo"
	"github.com/maltea/go-di-layout/pkg/application"
	"github.com/maltea/go-di-layout/pkg/paypal"
)

type stdloggerTransactionLog struct {
	prefix string
}

var _ application.TransactionLog = new(stdloggerTransactionLog)

func (s *stdloggerTransactionLog) Log(id, message string) {
	log.Print(s.prefix, id, message)
}

type defaultModule struct{}

func (*defaultModule) Configure(injector *dingo.Injector) {
	injector.Bind(new(application.TransactionLog)).ToInstance(&stdloggerTransactionLog{
		prefix: "foo_",
	})
}

func main() {
	injector, err := dingo.NewInjector(
		new(paypal.Module),
		new(defaultModule),
	)
	if err != nil {
		log.Fatal(err)
	}

	service, err := injector.GetInstance(application.Service{})
	if err != nil {
		log.Fatal(err)
	}

	if err := service.(*application.Service).MakeTransaction(99.95, "test transaction"); err != nil {
		log.Fatal(err)
	}
}
