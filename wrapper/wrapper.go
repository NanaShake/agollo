package wrapper

import (
	"context"
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/context/metadata"
	"github.com/micro/micro/v3/service/server"
	"github.com/wolfplus2048/mcbeam-plugins/v3/session"
	cli "github.com/wolfplus2048/mcbeam-plugins/v3/session/client"
	"strconv"
)

func SessionHandler(client client.Client) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			uid, _ := metadata.Get(ctx, "mcb-session-uid")
			_id, ok := metadata.Get(ctx, "mcb-session-id")
			if !ok {
				return h(ctx, req, rsp)
			}
			fid, ok := metadata.Get(ctx, "mcb-session-fid")
			if !ok {
				return h(ctx, req, rsp)
			}
			sid, _ := strconv.ParseInt(_id, 10, 64)

			s := cli.NewSession(sid, fid, uid)
			ctx = context.WithValue(ctx, session.SessionCtxKey{}, s)

			return h(ctx, req, rsp)
		}
	}
}
