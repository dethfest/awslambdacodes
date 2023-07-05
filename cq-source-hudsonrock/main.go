package main

import (
	"github.com/securezeron/cq-source-hudsonrock/plugin"

	"github.com/cloudquery/plugin-sdk/v3/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
