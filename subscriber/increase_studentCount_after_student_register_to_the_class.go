package subscriber

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/class/classstorage"
	"studyGoApp/pubsub"
)

func IncreaseStudentCountAfterStudentRegisterToTheClass(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicStudentRegisterToTheClass)

	store := classstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			mgs := <-c
			registerData := mgs.Data().(HasClassId)
			_ = store.IncreaseStudentCount(ctx, registerData.GetClassId())
		}
	}()
}

func RunIncreaseStudentCountAfterStudentRegisterToTheClass(appCtx component.AppContext) consumerJob {
	store := classstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Increase StudentCount after student register to the class!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			registerData := message.Data().(HasClassId)
			return store.IncreaseStudentCount(ctx, registerData.GetClassId())
		},
	}
}
