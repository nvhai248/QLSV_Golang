package studentbiz

import (
	"context"
	"log"
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

type RegisterClass interface {
	GetStudentRegister(ctx context.Context, ids []int) (map[int]int, error)
}

type listStudentBiz struct {
	store         ListStudentStore
	registerClass RegisterClass
}

func NewListStudentBiz(store ListStudentStore, registerClass RegisterClass) *listStudentBiz {
	return &listStudentBiz{
		store:         store,
		registerClass: registerClass,
	}
}

func (b *listStudentBiz) ListStudent(ctx context.Context,
	filter *studentmodel.Filter,
	paging *common.Paging,
) ([]studentmodel.Student, error) {
	result, err := b.store.ListDataByCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(studentmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapRegisterClass, err := b.registerClass.GetStudentRegister(ctx, ids)

	if err != nil {
		log.Println("Cannot get class registration!", err)
	}

	/* if err != nil {
		return nil, common.ErrCannotListEntity(studentmodel.EntityName, err)
	} */

	if v := mapRegisterClass; v != nil {
		for i, item := range result {
			result[i].ClassCount = mapRegisterClass[item.Id]
		}
	}

	return result, nil
}
