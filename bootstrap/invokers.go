package bootstrap

import (
	"github.com/saxenashivang/api-gateway/servers/rest"
	"go.uber.org/fx"
)

// GetInvokersOptions GetInvokersOptions: Please do not change the sequence because it invoker is lifecycle based method .
// So changing the sequence will be harmful.
func GetInvokersOptions() fx.Option {
	return fx.Invoke(
		rest.Run,
	)
}
