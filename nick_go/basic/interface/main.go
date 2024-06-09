package main

import (
	_case "basic/interface/case"
	"context"
	"os"
	"os/signal"
)

func main() {
	_case.SimpleCase()
	_case.CusNumCase()
	_case.BuiltInCase()
	_case.TTypeCase()
	_case.TTypeCase1()
	_case.InterfaceCase()
	_case.ReceiverCase()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
