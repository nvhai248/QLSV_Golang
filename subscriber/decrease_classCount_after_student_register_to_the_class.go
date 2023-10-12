package subscriber

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/student/studentstorage"
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
