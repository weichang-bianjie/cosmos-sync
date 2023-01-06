package ibc

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	ics721transfer "github.com/bianjieai/nft-transfer"
	ibctransfer "github.com/cosmos/ibc-go/v5/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v5/modules/core"
)

func init() {
	codec.RegisterAppModules(
		ibc.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
		ics721transfer.AppModuleBasic{},
	)
}
