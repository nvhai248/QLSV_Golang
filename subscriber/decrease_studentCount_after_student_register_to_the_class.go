package subscriber

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/class/classstorage"
	"studyGoApp/pubsub"
)

func DecreaseStudentCountAfterStudentRegisterToTheClass(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicStudentCancelRegistration)

	store := classstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			mgs := <-c
			registerData := mgs.Data().(HasClassId)
			_ = store.DecreaseStudentCount(ctx, registerData.GetClassId())
		}
	}()
}

func RunDecreaseStudentCountAfterStudentRegisterToTheClass(appCtx component.AppContext) consumerJob {
	store := classstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return consumerJob{
		Title: "Increase StudentCount after student register to the class",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			registerData := message.Data().(HasClassId)
			return store.DecreaseStudentCount(ctx, registerData.GetClassId())
		},
	}
}
