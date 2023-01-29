package ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
)

type DocNftMsgTransfer struct {
	PacketId         string   `bson:"packet_id"`
	SourcePort       string   `bson:"source_port"`
	SourceChannel    string   `bson:"source_channel"`
	ClassId          string   `bson:"class_id"`
	TokenIds         []string `bson:"token_ids"`
	Sender           string   `bson:"sender"`
	Receiver         string   `bson:"receiver"`
	TimeoutHeight    Height   `bson:"timeout_height"`
	TimeoutTimestamp int64    `bson:"timeout_timestamp"`
	Memo             string   `bson:"memo"`
}

func (m *DocNftMsgTransfer) GetType() string {
	return MsgTypeNftTransfer
}

func (m *DocNftMsgTransfer) BuildMsg(v interface{}) {
	msg := v.(*NftMsgTransfer)
	m.SourcePort = msg.SourcePort
	m.SourceChannel = msg.SourceChannel
	m.Sender = msg.Sender
	m.Receiver = msg.Receiver
	m.TimeoutTimestamp = int64(msg.TimeoutTimestamp)
	m.TimeoutHeight = loadHeight(msg.TimeoutHeight)
	m.ClassId = msg.ClassId
	m.TokenIds = msg.TokenIds
	m.Memo = msg.Memo
}

func (m *DocNftMsgTransfer) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   NftMsgTransfer
	)
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Sender, msg.Receiver)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
