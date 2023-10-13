package subscriber

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/student/studentstorage"
	"studyGoApp/pubsub"
)

func IncreaseClassCountAfterStudentRegisterToTheClass(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicStudentRegisterToTheClass)

	store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			mgs := <-c
			registerData := mgs.Data().(HasStudentId)
			_ = store.IncreaseClassCount(ctx, registerData.GetStudentId())
		}
	}()
}

func RunIncreaseClassCountAfterStudentRegisterToTheClass(appCtx component.AppContext) consumerJob {
	store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Increase StudentCount after student register to the class",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			registerData := message.Data().(HasStudentId)
			return store.IncreaseClassCount(ctx, registerData.GetStudentId())
		},
	}
}
