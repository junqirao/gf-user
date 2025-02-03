package token

import (
	"context"
	"net"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gipv4"

	"gf-user/internal/model"
	"gf-user/internal/packed/ipgeo"
)

func (t sToken) ListUserRefreshTokenDetails(ctx context.Context, locale string) (dts []*model.RefreshTokenDetail, err error) {
	var (
		info   = t.GetTokenInfoFromCtx(ctx)
		rtsMap = make(map[string]*refreshToken)
		keys   []string
	)

	if locale == "" {
		locale = "en"
	}

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

	ipgeoDB := ipgeo.DB()

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
			if gipv4.IsIntranet(detail.IP) {
				detail.City = "-"
				detail.Country = "-"
			} else if ipgeoDB == nil {
				detail.City = "Unknown"
				detail.Country = "Unknown"
			} else {
				city, err := ipgeoDB.City(net.ParseIP(detail.IP))
				if err != nil {
					g.Log().Warningf(ctx, "ipgeo get [ip:%s] city failed: %v", detail.IP, err)
					continue
				}
				detail.City = city.City.Names[locale]
				detail.Country = city.Country.Names[locale]
				detail.CountryISO = city.Country.IsoCode

				// mask ip
				ipParts := strings.Split(detail.IP, ".")
				for i, part := range ipParts {
					if i == 0 || i == len(ipParts)-1 {
						continue
					}
					ipParts[i] = strings.Repeat("*", len(part))
				}

				detail.IP = strings.Join(ipParts, ".")
			}
		}
		dts = append(dts, detail)
	}

	sort.Slice(dts, func(i, j int) bool {
		return dts[i].ExpireAt > dts[j].ExpireAt
	})
	return
}

func (t sToken) ClearRefreshTokens(ctx context.Context) (cnt int64, err error) {
	tokenInfo := t.GetTokenInfoFromCtx(ctx)
	rts, err := t.getUserRefreshTokens(ctx, t.getUserRefreshTokenKey(tokenInfo.AccountId))
	if err != nil {
		return
	}
	var keys []string
	for _, rt := range rts {
		keys = append(keys, t.getUserRefreshTokenExtraDataKey(tokenInfo.AccountId, rt.Key))
	}
	keys = append(keys, t.getUserRefreshTokenKey(tokenInfo.AccountId))
	return g.Redis().Unlink(ctx, keys...)
}
