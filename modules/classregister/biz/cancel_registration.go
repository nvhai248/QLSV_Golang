package classregisterbiz

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/component/asyncjob"
	classregistermodel "studyGoApp/modules/classregister/model"
)

type CancelRegistrationStore interface {
	DeleteClassRegister(ctx context.Context, studentId, classId int) error
	FindClassRegister(ctx context.Context, studentId, classId int) (*classregistermodel.Register, error)
}

type DecreaseClassCountStore interface {
	DecreaseClassCount(ctx context.Context, id int) error
}

type DecreaseStudentCountStore interface {
	DecreaseStudentCount(ctx context.Context, id int) error
}

type cancelRegistrationBiz struct {
	store                     CancelRegistrationStore
	decreaseClassCountStore   DecreaseClassCountStore
	decreaseStudentCountStore DecreaseStudentCountStore
}

func NewCancelRegistrationBiz(
	store CancelRegistrationStore,
	decreaseClassCountStore DecreaseClassCountStore,
	decreaseStudentCountStore DecreaseStudentCountStore) *cancelRegistrationBiz {
	return &cancelRegistrationBiz{
		store:                     store,
		decreaseClassCountStore:   decreaseClassCountStore,
		decreaseStudentCountStore: decreaseStudentCountStore,
	}
}

func (b *cancelRegistrationBiz) CancelRegistration(ctx context.Context, studentId, classId int) error {
	if err := b.store.DeleteClassRegister(ctx, studentId, classId); err != nil {
		return classregistermodel.ErrorCannotCancelRegistration(err)
	}

	// side effect

	go func() {
		defer common.AppRecover()
		job1 := asyncjob.NewJob(func(ctx context.Context) error {
			return b.decreaseClassCountStore.DecreaseClassCount(ctx, studentId)
		})

		job2 := asyncjob.NewJob(func(ctx context.Context) error {
			return b.decreaseStudentCountStore.DecreaseStudentCount(ctx, classId)
		})

		_ = asyncjob.NewGroup(true, *job1, *job2).Run(ctx)
	}()

	/* go func() {
		defer common.AppRecover()
		_ = b.decreaseClassCountStore.DecreaseClassCount(ctx, studentId)
		_ = b.decreaseStudentCountStore.DecreaseStudentCount(ctx, classId)
	}() */

	return nil
}
