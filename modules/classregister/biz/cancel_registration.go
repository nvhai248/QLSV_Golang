package classregisterbiz

import (
	"context"
	"studyGoApp/common"
	classregistermodel "studyGoApp/modules/classregister/model"
	"studyGoApp/pubsub"
)

type CancelRegistrationStore interface {
	DeleteClassRegister(ctx context.Context, data *classregistermodel.Register) error
}

/* type DecreaseClassCountStore interface {
	DecreaseClassCount(ctx context.Context, id int) error
}

type DecreaseStudentCountStore interface {
	DecreaseStudentCount(ctx context.Context, id int) error
} */

type cancelRegistrationBiz struct {
	store CancelRegistrationStore
	/* decreaseClassCountStore   DecreaseClassCountStore
	decreaseStudentCountStore DecreaseStudentCountStore */
	pubsub pubsub.Pubsub
}

func NewCancelRegistrationBiz(
	store CancelRegistrationStore,
	/* decreaseClassCountStore DecreaseClassCountStore,
	decreaseStudentCountStore DecreaseStudentCountStore */
	pubsub pubsub.Pubsub,
) *cancelRegistrationBiz {
	return &cancelRegistrationBiz{
		store: store,
		/* decreaseClassCountStore:   decreaseClassCountStore,
		decreaseStudentCountStore: decreaseStudentCountStore, */
		pubsub: pubsub,
	}
}

func (b *cancelRegistrationBiz) CancelRegistration(ctx context.Context, data *classregistermodel.Register) error {
	if err := b.store.DeleteClassRegister(ctx, data); err != nil {
		return classregistermodel.ErrorCannotCancelRegistration(err)
	}

	// side effect

	b.pubsub.Publish(ctx, common.TopicStudentCancelRegistration, pubsub.NewMessage(data))

	/* go func() {
		defer common.AppRecover()
		job1 := asyncjob.NewJob(func(ctx context.Context) error {
			return b.decreaseClassCountStore.DecreaseClassCount(ctx, studentId)
		})

		job2 := asyncjob.NewJob(func(ctx context.Context) error {
			return b.decreaseStudentCountStore.DecreaseStudentCount(ctx, classId)
		})

		_ = asyncjob.NewGroup(true, *job1, *job2).Run(ctx)
	}() */

	/* go func() {
		defer common.AppRecover()
		_ = b.decreaseClassCountStore.DecreaseClassCount(ctx, studentId)
		_ = b.decreaseStudentCountStore.DecreaseStudentCount(ctx, classId)
	}() */

	return nil
}
