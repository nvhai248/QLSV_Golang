package studentbiz

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/component/tokenprovider"
	"studyGoApp/modules/student/studentmodel"
)

type LoginStorage interface {
	FindByStudentID(ctx context.Context,
		studentID string) (*studentmodel.StudentDetail, error)
}

type TokenConfig interface {
	GetAtExp() int
}

type loginBiz struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(
	storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBiz {
	return &loginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// Process login
//1. Find Username, password
//2. Hash pass from input and compare with pass in db
//3. Provider: issue JWT token for client
//3.1. Access token and refresh token
//4. Return token(s)

func (biz *loginBiz) Login(ctx context.Context, data *studentmodel.StudentLogin) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindByStudentID(ctx, data.StudentId)

	if err != nil {
		return nil, common.ErrCannotGetEntity(studentmodel.EntityName, err)
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, common.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	/* refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := studentmodel.NewAccount(accessToken) */

	return accessToken, nil
}
