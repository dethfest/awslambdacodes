package resources

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/securezeron/hudsonrock/cq-source-hudsonrock/internal/hudsonrock"
)

func HudsonRockTable() *schema.Table {
	return &schema.Table{
		Name:     "hudsonrock_table",
		Resolver: fetchHudsonTable,
		Transform: transformers.TransformWithStruct(&hudsonrock.HudsonRock{}),
	}
}

func fetchHudsonTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {

	c := meta.(*client.Client)

	// hudsonrock, err := c.HUDSONROCK.GetHudsonRock (num: 0)
	// if err != nil {
	// 	return nil, err
	// }
	// res <- hudsonrock
	
	for i := 1; i <hudsonrock.num; i++ {
		hudsonrock, err := c.HUDSONROCK.GetHudsonRock(i)
		if err != nil {
			return err
		}
		res <- comic
	}
	return nil
}
