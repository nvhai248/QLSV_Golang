package classbiz

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/modules/class/classmodel"
)

type CreateClassStore interface {
	CreateClass(ctx context.Context, data *classmodel.Class) error
}

type createClassBiz struct {
	store CreateClassStore
}

func NewCreateClassBiz(store CreateClassStore) *createClassBiz {
	return &createClassBiz{store: store}
}

func (biz *createClassBiz) Create(ctx context.Context, data *classmodel.Class) error {
	data.Status = 1

	if err := biz.store.CreateClass(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(classmodel.EntityName, err)
	}

	return nil
}
