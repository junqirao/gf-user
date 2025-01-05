package space

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/junqirao/gocomponents/response"

	"gf-user/internal/dao"
	"gf-user/internal/model"
	"gf-user/internal/model/do"
	"gf-user/internal/model/entity"
	"gf-user/internal/service"
)

func init() {
	service.RegisterSpace(&sSpace{})
}

type sSpace struct {
}

func (s sSpace) GetSpaceInfo(ctx context.Context, spaceId int64) (space *model.Space, err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	v, err := dao.Space.Ctx(ctx).Where(dao.Space.Columns().Id, spaceId).One()
	if err != nil {
		return
	}
	if v.IsEmpty() {
		err = response.CodeNotFound.WithDetail(spaceId)
		return
	}

	sd := new(do.Space)
	if err = v.Struct(sd); err != nil {
		return
	}
	space = &model.Space{
		Id:          gconv.Int64(sd.Id),
		Name:        gconv.String(sd.Name),
		IsOwner:     tokenInfo.AccountId == gconv.String(sd.Owner),
		Description: gconv.String(sd.Description),
		Profile:     gconv.Map(sd.Profile),
		UpdateAt:    sd.UpdateAt,
		CreatedAt:   sd.CreatedAt,
	}
	return
}

func (s sSpace) CreateSpace(ctx context.Context, in model.CreateSpaceInput) (space *model.Space, err error) {
	token := service.Token().GetTokenInfoFromCtx(ctx)
	account, err := service.Account().IsValid(ctx, token.AccountId)
	if err != nil {
		return
	}
	res, err := dao.Space.Ctx(ctx).Insert(entity.Space{
		Name:        in.Name,
		Logo:        in.Logo,
		Owner:       token.AccountId,
		Description: in.Description,
		Profile:     "{}",
		CreatedAt:   gtime.Now(),
	})
	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		return
	}
	_, err = service.User().CreateSpaceUser(ctx, account, id)
	if err != nil {
		return
	}
	return s.GetSpaceInfo(ctx, id)
}
