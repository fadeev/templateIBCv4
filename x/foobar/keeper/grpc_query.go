package keeper

import (
	"github.com/fadeev/templateIBCv4/x/foobar/types"
)

var _ types.QueryServer = Keeper{}
