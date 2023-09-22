package classbiz

import (
	"context"
	"studyGoApp/common"
	classregistermodel "studyGoApp/modules/classregister/model"
	"studyGoApp/modules/student/studentmodel"
)

type ListRegisteredStuStore interface {
	GetListSimpleStudentByConditions(ctx context.Context,
		conditions map[string]interface{},
		filter *classregistermodel.Filter,
		paging *common.Paging,
	) ([]common.SimpleStudent, error)
}

type lisRegisteredStuBiz struct {
	store ListRegisteredStuStore
}

func NewListRegisteredStuBiz(store ListRegisteredStuStore) *lisRegisteredStuBiz {
	return &lisRegisteredStuBiz{store: store}
}

func (b *lisRegisteredStuBiz) GetListRegisteredStu(ctx context.Context,
	filter *classregistermodel.Filter,
	paging *common.Paging,
) ([]common.SimpleStudent, error) {
	result, err := b.store.GetListSimpleStudentByConditions(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(studentmodel.EntityName, err)
	}

	return result, nil
}
