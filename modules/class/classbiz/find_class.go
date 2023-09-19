package classbiz

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/modules/class/classmodel"
)

type FindClassStore interface {
	FindClassById(ctx context.Context, id int, condition map[string]interface{}) (*classmodel.Class, error)
}

type findClassBiz struct {
	store FindClassStore
}

func NewFindClassBiz(store FindClassStore) *findClassBiz {
	return &findClassBiz{store: store}
}

func (b *findClassBiz) FindClassById(ctx context.Context, id int, condition map[string]interface{}) (*classmodel.Class, error) {
	result, err := b.store.FindClassById(ctx, id, condition)

	if err != nil {
		if err != common.ErrorNoRows {
			return nil, common.ErrCannotGetEntity(classmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(classmodel.EntityName, err)
	}

	if result.Status != 1 {
		return nil, common.ErrEntityDeleted(classmodel.EntityName, nil)
	}

	return result, err
}
