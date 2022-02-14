package demo

import "github.com/bitwormhole/starter/application"

func ExportConfigForStarterMailDemo(cb application.ConfigBuilder) error {

	return autoGenConfig(cb)
}
