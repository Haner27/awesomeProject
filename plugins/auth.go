package plugin

import (
	"awesomeProject/conf"
	"context"
	"errors"
	"github.com/smallnest/rpcx/protocol"
)

func AuthFunc(ctx context.Context, req *protocol.Message, token string) error {
	if token == conf.AuthToken {
		return nil
	}
	return errors.New("invalid token")
}
