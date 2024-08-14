package main

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

type Subscription struct {
	sSub       ethereum.Subscription
	sContract  *SContract
	sEventName string
	sLogs      *chan types.Log
}
