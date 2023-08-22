package studentbiz

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/modules/student/studentmodel"
)

type ListStudentStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *studentmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]studentmodel.Student, error)
}

type listStudentBiz struct {
	store ListStudentStore
}

func NewListStudentBiz(store ListStudentStore) *listStudentBiz {
	return &listStudentBiz{
		store: store,
	}
}

func (b *listStudentBiz) ListStudent(ctx context.Context,
	filter *studentmodel.Filter,
	paging *common.Paging,
) ([]studentmodel.Student, error) {
	result, err := b.store.ListDataByCondition(ctx, nil, filter, paging)

	return result, err
}
