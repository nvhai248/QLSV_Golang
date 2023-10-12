package subscriber

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/class/classstorage"
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