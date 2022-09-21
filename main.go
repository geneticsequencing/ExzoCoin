package main

import (
	_ "embed"

	"github.com/ExzoNetwork/ExzoCoin/command/root"
	"github.com/ExzoNetwork/ExzoCoin/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
