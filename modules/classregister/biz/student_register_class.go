package classregisterbiz

import (
	"context"
	"studyGoApp/common"
	classregistermodel "studyGoApp/modules/classregister/model"
)

type RegisterStore interface {
	CreateClassRegister(ctx context.Context, data *classregistermodel.Register) error
	FindClassRegister(ctx context.Context, studentId, classId int) (*classregistermodel.Register, error)
}

type IncreaseStudentCountStore interface {
	IncreaseStudentCount(ctx context.Context, id int) error
}

type IncreaseClassCountStore interface {
	IncreaseClassCount(ctx context.Context, id int) error
}

type registerBiz struct {
	store                   RegisterStore
	increaseStudentStore    IncreaseStudentCountStore
	increaseClassCountStore IncreaseClassCountStore
}

func NewRegisterBiz(store RegisterStore,
	increaseStudentStore IncreaseStudentCountStore,
	increaseClassCountStore IncreaseClassCountStore) *registerBiz {
	return &registerBiz{
		store:                   store,
		increaseStudentStore:    increaseStudentStore,
		increaseClassCountStore: increaseClassCountStore,
	}
}

func (b *registerBiz) Register(ctx context.Context, data *classregistermodel.Register) error {

	registration, err := b.store.FindClassRegister(ctx, data.StudentId, data.ClassId)

	if err != nil && err != common.ErrorNoRows {
		return err
	}

	if registration != nil {
		return classregistermodel.ErrorIsRegistered(err)
	}

	if err := b.store.CreateClassRegister(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(classregistermodel.EntityName, err)
	}

	// side effect
	_ = b.increaseClassCountStore.IncreaseClassCount(ctx, data.StudentId)
	_ = b.increaseStudentStore.IncreaseStudentCount(ctx, data.ClassId)

	return nil
}
