package ipgeo

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/oschwald/geoip2-golang"
)

var (
	db *geoip2.Reader
)

func Init(ctx context.Context) {
	v, err := g.Cfg().Get(ctx, "ipgeo.db")
	if err != nil {
		g.Log().Infof(ctx, "ipgeo database config error: %s", err.Error())
		return
	}
	if v.IsEmpty() {
		g.Log().Infof(ctx, "ipgeo database config is empty, skip init.")
		return
	}
	db, err = geoip2.Open(v.String())
	if err != nil {
		g.Log().Errorf(ctx, "open ipgeo database failed: %v", err)
		return
	}
}

func DB() *geoip2.Reader {
	return db
}
