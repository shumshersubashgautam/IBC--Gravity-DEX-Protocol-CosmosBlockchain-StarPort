package keeper

import (
	"github.com/toschdev/pluto/x/pluto/types"
)

var _ types.QueryServer = Keeper{}
