package token

import (
	"context"
	"net"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/model"
	"gf-user/internal/packed/ipgeo"
)

func (t sToken) ListUserRefreshTokenDetails(ctx context.Context) (dts []*model.RefreshTokenDetail, err error) {
	var (
		info   = t.GetTokenInfoFromCtx(ctx)
		rtsMap = make(map[string]*refreshToken)
		keys   []string
	)

	rts, err := t.getUserRefreshTokens(ctx, t.getUserRefreshTokenKey(info.AccountId))
	if err != nil {
		return
	}

	for _, rt := range rts {
		keys = append(keys, t.getUserRefreshTokenExtraDataKey(info.AccountId, rt.Key))
		rtsMap[rt.Key] = rt
	}
	res, err := g.Redis().MGet(ctx, keys...)
	if err != nil {
		return
	}

	for k, v := range res {
		bs, err := gbase64.DecodeString(v.String())
		if err != nil {
			continue
		}

		var extra model.RefreshTokenExtraData
		if err = gjson.Unmarshal(bs, &extra); err != nil {
			continue
		}
		parts := strings.Split(k, ":")
		key := parts[len(parts)-1]
		detail := &model.RefreshTokenDetail{
			Key:       key,
			IP:        extra.ClientIP,
			UserAgent: extra.UA,
			From:      extra.From,
			ExpireAt:  rtsMap[key].ExpireAt,
		}
		if detail.IP != "" {
			city, err := ipgeo.DB().City(net.ParseIP(detail.IP))
			if err != nil {
				g.Log().Warningf(ctx, "ipgeo get [ip:%s] city failed: %v", detail.IP, err)
				continue
			}
			detail.City = city.City.Names
			detail.Country = city.Country.Names
			detail.CountryISO = city.Country.IsoCode
		}
		dts = append(dts, detail)
	}

	sort.Slice(dts, func(i, j int) bool {
		return dts[i].ExpireAt > dts[j].ExpireAt
	})
	return
}
