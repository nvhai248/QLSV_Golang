package subscriber

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/student/studentstorage"
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
