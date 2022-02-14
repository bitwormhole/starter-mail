package lib

import "github.com/bitwormhole/starter/application"

func ExportConfigForStarterMailLib(cb application.ConfigBuilder) error {

	return autoGenConfig(cb)
}
