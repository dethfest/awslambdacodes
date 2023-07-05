package plugin

import (
	"github.com/securezeron/cq-source-hudsonrock/client"
	"github.com/securezeron/cq-source-hudsonrock/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"securezeron-hudsonrock",
		Version,
		schema.Tables{
			resources.HudsonRockTable(),
		},
		client.New,
	)
}
