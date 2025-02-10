package space

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/junqirao/gocomponents/response"

	"gf-user/internal/consts"
	"gf-user/internal/dao"
	"gf-user/internal/model"
	"gf-user/internal/model/code"
	"gf-user/internal/model/do"
	"gf-user/internal/model/entity"
	"gf-user/internal/service"
)

func (s sSpace) CreateInvitation(ctx context.Context, in model.CreateSpaceInvitationInput) (err error) {
	token := service.Token().GetTokenInfoFromCtx(ctx)
	sp, err := service.Space().GetSpaceInfo(ctx, token.SpaceId)
	if err != nil {
		return
	}
	usr, err := service.User().GetUserByAccountId(ctx, token.AccountId, token.SpaceId)
	if err != nil {
		return
	}
	if !sp.IsOwner && gconv.Int(usr.Type) != consts.UserTypeManager {
		err = response.CodePermissionDeny
		return
	}
	account, err := service.Account().GetAccount(ctx, in.TargetAccount)
	if err != nil {
		return
	}
	if gconv.Int(account.Status) != consts.AccountStatusNormal {
		err = code.ErrAccountLocked.WithMessage("target account frozen")
		return
	}

	_, err = dao.SpaceInvitation.Ctx(ctx).Insert(entity.SpaceInvitation{
		Space:     int(token.SpaceId),
		From:      token.AccountId,
		Status:    consts.SpaceInvitationStatusCreate,
		Target:    gconv.String(account.Id),
		Comment:   in.Comment,
		CreatedAt: gtime.Now(),
	})

	return
}

func (s sSpace) HandleInvitation(ctx context.Context, id, op int) (err error) {
	token := service.Token().GetTokenInfoFromCtx(ctx)
	account, err := service.Account().IsValid(ctx, token.AccountId)
	if err != nil {
		return
	}
	v, err := dao.SpaceInvitation.Ctx(ctx).Where(dao.SpaceInvitation.Columns().Id, id).One()
	if err != nil {
		return
	}
	if v.IsEmpty() {
		err = response.CodeNotFound.WithDetail(id)
		return
	}
	si := new(do.SpaceInvitation)
	if err = v.Struct(si); err != nil {
		return
	}

	if op == consts.SpaceInvitationStatusCancel && gconv.String(si.Target) == token.AccountId {
		err = response.CodeInvalidParameter.WithDetail("invalid operation")
		return
	}

	if op == consts.SpaceInvitationStatusCancel || op == consts.SpaceInvitationStatusReject {
		_, err = dao.SpaceInvitation.Ctx(ctx).Where(dao.SpaceInvitation.Columns().Id, id).Update(g.Map{
			dao.SpaceInvitation.Columns().Status: op,
		})
		return
	}

	if gconv.String(si.From) == token.AccountId {
		err = response.CodeFromHttpStatus(http.StatusBadRequest).WithDetail("cannot accept your invitation")
		return
	}

	exist, err := service.User().Exist(ctx, gconv.String(si.Target), gconv.Int64(si.Space))
	if err != nil {
		return
	}
	if exist {
		err = response.CodeConflict.WithDetail("account already in space")
		return
	}
	err = s.acceptInvitation(ctx, account, si)
	return
}

func (s sSpace) MyInvitations(ctx context.Context) (target, source []*model.SpaceInvitation, err error) {
	token := service.Token().GetTokenInfoFromCtx(ctx)
	list := make([]*do.SpaceInvitation, 0)
	err = dao.SpaceInvitation.Ctx(ctx).
		Where(dao.SpaceInvitation.Columns().From, token.AccountId).
		WhereOr(dao.SpaceInvitation.Columns().Target, token.AccountId).Scan(&list)
	if err != nil {
		return
	}
	target, source = make([]*model.SpaceInvitation, 0), make([]*model.SpaceInvitation, 0)

	var (
		accountIdsMap = make(map[any]struct{})
		accountIds    []string
		accountMap    = make(map[string]*model.Account)
		spaceIdsMap   = make(map[any]struct{})
		spaceIds      []int64
		spaceMap      = make(map[int64]*model.Space)
	)
	for _, v := range list {
		accountIdsMap[v.Target] = struct{}{}
		accountIdsMap[v.From] = struct{}{}
		spaceIdsMap[v.Space] = struct{}{}
	}
	for id, _ := range accountIdsMap {
		accountIds = append(accountIds, gconv.String(id))
	}
	for id, _ := range spaceIdsMap {
		spaceIds = append(spaceIds, gconv.Int64(id))
	}

	accounts, err := service.Account().GetAccountByIds(ctx, accountIds)
	if err != nil {
		return
	}
	for _, account := range accounts {
		accountMap[gconv.String(account.Id)] = account
	}

	spaces, err := s.GetSpaceListByIds(ctx, spaceIds)
	if err != nil {
		return
	}
	for _, space := range spaces {
		spaceMap[gconv.Int64(space.Id)] = space
	}

	for _, v := range list {
		inv := &model.SpaceInvitation{
			Id:        v.Id,
			Space:     spaceMap[gconv.Int64(v.Space)],
			From:      accountMap[gconv.String(v.From)].AccountBrief,
			To:        accountMap[gconv.String(v.Target)].AccountBrief,
			Status:    v.Status,
			Comment:   v.Comment,
			CreatedAt: v.CreatedAt,
		}
		if gconv.String(v.Target) == token.AccountId {
			target = append(target, inv)
		} else {
			if token.SpaceId != consts.DefaultSpaceId && token.SpaceId != gconv.Int64(v.Space) {
				continue
			}
			source = append(source, inv)
		}
	}
	return
}

func (s sSpace) acceptInvitation(ctx context.Context, account *do.Account, invitation *do.SpaceInvitation) (err error) {
	return dao.SpaceInvitation.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.SpaceInvitation.Ctx(ctx).Where(dao.SpaceInvitation.Columns().Id, invitation.Id).Update(g.Map{
			dao.SpaceInvitation.Columns().Status: consts.SpaceInvitationStatusAccept,
		})
		if err != nil {
			return
		}

		_, err = service.User().CreateSpaceUser(ctx, account, gconv.Int64(invitation.Space))
		return
	})
}
