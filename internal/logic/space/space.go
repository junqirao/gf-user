package space

import (
	"context"
	"fmt"
	"sort"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/junqirao/gocomponents/response"

	"gf-user/internal/dao"
	"gf-user/internal/model"
	"gf-user/internal/model/code"
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
	space = s.convert2SpaceInfo(ctx, sd)
	return
}

func (s sSpace) GetSpaceList(ctx context.Context, accountId string) (spaceList []*model.Space, err error) {
	ul := make([]*do.User, 0)
	if err = dao.User.Ctx(ctx).Where(dao.User.Columns().Account, accountId).Scan(&ul); err != nil {
		return
	}
	ids := make([]int64, 0)
	for _, v := range ul {
		ids = append(ids, gconv.Int64(v.Space))
	}
	if len(ids) == 0 {
		err = code.ErrUserNotExist
		return
	}

	sps := make([]*do.Space, 0)
	if err = dao.Space.Ctx(ctx).WhereIn(dao.Space.Columns().Id, ids).Scan(&sps); err != nil {
		return
	}

	for _, v := range sps {
		spaceList = append(spaceList, s.convert2SpaceInfo(ctx, v))
	}
	sort.Slice(spaceList, func(i, j int) bool {
		return spaceList[i].Id < spaceList[j].Id
	})
	return
}

func (s sSpace) convert2SpaceInfo(ctx context.Context, v *do.Space) *model.Space {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	space := &model.Space{
		Id:          gconv.Int64(v.Id),
		Name:        gconv.String(v.Name),
		LogoKey:     gconv.String(v.Logo),
		IsOwner:     tokenInfo.AccountId == gconv.String(v.Owner),
		Description: gconv.String(v.Description),
		Profile:     gconv.Map(v.Profile),
		UpdateAt:    v.UpdateAt,
		CreatedAt:   v.CreatedAt,
	}
	if space.LogoKey != "" {
		token := service.Token().GetTokenInfoFromCtx(ctx)
		space.Logo = fmt.Sprintf("/v1/storage/space/logo?key=%s&space_id=%v&access_token=%s", space.LogoKey, space.Id, token.AccessToken)
	}
	return space
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
