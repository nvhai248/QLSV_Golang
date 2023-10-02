package classbiz

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/modules/class/classmodel"
)

type ListClassStore interface {
	ListClassByCondition(ctx context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]classmodel.Class, error)
}

type RegisterClass interface {
	GetNumberOfStudentRegisteredInClass(ctx context.Context, ids []int) (map[int]int, error)
}

type listClassBiz struct {
	store         ListClassStore
	registerClass RegisterClass
}

func NewListClassBiz(store ListClassStore, registerClass RegisterClass) *listClassBiz {
	return &listClassBiz{
		store:         store,
		registerClass: registerClass,
	}
}

func (b *listClassBiz) ListClass(ctx context.Context,
	paging *common.Paging,
) ([]classmodel.Class, error) {
	result, err := b.store.ListClassByCondition(ctx, nil, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(classmodel.EntityName, err)
	}

	/* ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapRegisterClass, err := b.registerClass.GetNumberOfStudentRegisteredInClass(ctx, ids)

	if err != nil {
		log.Println("Cannot get class registration!", err)
	}

	if err != nil {
		return nil, common.ErrCannotListEntity(studentmodel.EntityName, err)
	}

	if v := mapRegisterClass; v != nil {
		for i, item := range result {
			result[i].StudentCount = mapRegisterClass[item.Id]
		}
	} */

	return result, nil
}
