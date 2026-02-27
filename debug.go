package sio

import eio "github.com/zulfirman/socket.io-go/engine.io"

type Debugger = eio.Debugger

var (
	NewPrintDebugger = eio.NewPrintDebugger
	newNoopDebugger  = eio.NewNoopDebugger
)
