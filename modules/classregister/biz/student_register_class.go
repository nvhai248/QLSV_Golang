package classregisterbiz

import (
	"context"
	"studyGoApp/common"
	classregistermodel "studyGoApp/modules/classregister/model"
	"studyGoApp/pubsub"
)

type RegisterStore interface {
	CreateClassRegister(ctx context.Context, data *classregistermodel.Register) error
	FindClassRegister(ctx context.Context, studentId, classId int) (*classregistermodel.Register, error)
}

/* type IncreaseStudentCountStore interface {
	IncreaseStudentCount(ctx context.Context, id int) error
}

type IncreaseClassCountStore interface {
	IncreaseClassCount(ctx context.Context, id int) error
} */

type registerBiz struct {
	store RegisterStore
	/* increaseStudentStore    IncreaseStudentCountStore
	increaseClassCountStore IncreaseClassCountStore */
	pubsub pubsub.Pubsub
}

func NewRegisterBiz(store RegisterStore,
	/* increaseStudentStore IncreaseStudentCountStore,
	increaseClassCountStore IncreaseClassCountStore, */
	pubsub pubsub.Pubsub) *registerBiz {
	return &registerBiz{
		store: store,
		/* increaseStudentStore:    increaseStudentStore,
		increaseClassCountStore: increaseClassCountStore, */
		pubsub: pubsub,
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
	// new solution: use pub/sub

	b.pubsub.Publish(ctx, common.TopicStudentRegisterToTheClass, pubsub.NewMessage(data))

	// async job
	/* go func() {
		defer common.AppRecover()
		job1 := asyncjob.NewJob(func(ctx context.Context) error {
			return b.increaseClassCountStore.IncreaseClassCount(ctx, data.StudentId)
		})

		job2 := asyncjob.NewJob(func(ctx context.Context) error {
			return b.increaseStudentStore.IncreaseStudentCount(ctx, data.ClassId)
		})

		_ = asyncjob.NewGroup(true, *job1, *job2).Run(ctx)
	}() */

	/* go func() {
		defer common.AppRecover()
		_ = b.increaseClassCountStore.IncreaseClassCount(ctx, data.StudentId)
		_ = b.increaseStudentStore.IncreaseStudentCount(ctx, data.ClassId)
	}() */

	return nil
}
