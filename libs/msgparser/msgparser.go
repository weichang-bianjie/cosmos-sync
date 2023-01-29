package msgparser

import (
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/client"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/types"
	"strings"
)

type MsgParser interface {
	HandleTxMsg(v types.SdkMsg) MsgDocInfo
	MsgType(v types.SdkMsg) string
}

var (
	_client client.MsgClient
)

func NewMsgParser(router Router) MsgParser {
	return &msgParser{
		rh: router,
	}
}

type msgParser struct {
	rh Router
}

// Handler returns the MsgServiceHandler for a given msg or nil if not found.
func (parser msgParser) getModule(v types.SdkMsg) string {
	var (
		route string
	)

	data := types.MsgTypeURL(v)
	if strings.HasPrefix(data, "/ibc.core.") {
		route = IbcRouteKey
	} else if strings.HasPrefix(data, "/ibc.applications.") {
		route = IbcTransferRouteKey
	} else {
		route = data
	}
	return route
}

func (parser msgParser) HandleTxMsg(v types.SdkMsg) MsgDocInfo {
	module := parser.getModule(v)
	handleFunc, err := parser.rh.GetRoute(module)
	if err != nil {
		logger.Warn(err.Error(),
			logger.String("route", module),
			logger.String("type", module))
		return MsgDocInfo{}
	}
	return handleFunc(v)
}

func (parser msgParser) MsgType(v types.SdkMsg) string {
	switch v.(type) {
	case *MsgRecvPacket:
		return MsgTypeRecvPacket
	case *MsgTransfer:
		return MsgTypeIBCTransfer
	case *NftMsgTransfer:
		return MsgTypeNftTransfer
	case *MsgUpdateClient:
		return MsgTypeUpdateClient
	case *MsgChannelOpenConfirm:
		return MsgTypeChannelOpenConfirm
	case *MsgTimeout:
		return MsgTypeTimeout
	case *MsgAcknowledgement:
		return MsgTypeAcknowledgement
	case *MsgCreateClient:
		return MsgTypeCreateClient
	case *MsgUpgradeClient:
		return MsgTypeUpgradeClient
	case *MsgSubmitMisbehaviour:
		return MsgTypeSubmitMisbehaviourClient
	case *MsgConnectionOpenInit:
		return MsgTypeConnectionOpenInit
	case *MsgConnectionOpenTry:
		return MsgTypeConnectionOpenTry
	case *MsgConnectionOpenAck:
		return MsgTypeConnectionOpenAck
	case *MsgConnectionOpenConfirm:
		return MsgTypeConnectionOpenConfirm
	case *MsgChannelOpenInit:
		return MsgTypeChannelOpenInit
	case *MsgChannelOpenTry:
		return MsgTypeChannelOpenTry
	case *MsgChannelOpenAck:
		return MsgTypeChannelOpenAck
	case *MsgChannelCloseInit:
		return MsgTypeChannelCloseInit
	case *MsgChannelCloseConfirm:
		return MsgTypeChannelCloseConfirm
	case *MsgTimeoutOnClose:
		return MsgTypeTimeoutOnClose
	}
	return ""
}

func init() {
	codec.MakeEncodingConfig()
	_client = client.NewMsgClient()
}
func handleIbc(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Ibc.HandleTxMsg(v)
	return docInfo
}
