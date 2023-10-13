package subscriber

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/student/studentstorage"
	"studyGoApp/pubsub"
)

func DecreaseClassCountAfterStudentRegisterToTheClass(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicStudentCancelRegistration)

	store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			mgs := <-c
			registerData := mgs.Data().(HasStudentId)
			_ = store.DecreaseClassCount(ctx, registerData.GetStudentId())
		}
	}()
}

func RunDecreaseClassCountAfterStudentRegisterToTheClass(appCtx component.AppContext) consumerJob {
	store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Decrease classCount after student cancel registration!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			registerData := message.Data().(HasStudentId)
			return store.DecreaseClassCount(ctx, registerData.GetStudentId())
		},
	}
}
