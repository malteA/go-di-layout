package paypal

import (
	"flamingo.me/dingo"

	"github.com/maltea/go-di-layout/pkg/application"
)

type Module struct{}

func (m *Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(application.CreditCardProcessor)).To(new(paypalCCProcessor))
}
