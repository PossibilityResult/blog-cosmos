package keeper

import (
	"github.com/username/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
