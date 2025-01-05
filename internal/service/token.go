// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-user/internal/model"
)

type (
	IToken interface {
		GenerateAccessToken(ctx context.Context, user *model.UserAccount) (accessToken string, refreshToken string, err error)
		ValidAccessToken(ctx context.Context, accessToken string) (tokenInfo *model.TokenInfo, err error)
		// GetTokenInfoFromCtx only work under middleware.AuthToken
		GetTokenInfoFromCtx(ctx context.Context) (tokenInfo model.TokenInfo)
		RefreshToken(ctx context.Context, user *model.UserAccount, claims *model.RefreshTokenClaims) (newAccessToken string, err error)
		ParseRefreshToken(_ context.Context, refreshToken string) (claims *model.RefreshTokenClaims, err error)
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
