package bootstrap

import "go.uber.org/fx"

func GetApp() *fx.App {
	opts := GetProviderOptions()
	opts = append(opts, GetInvokersOptions())
	return fx.New(
		opts...,
	)
}
