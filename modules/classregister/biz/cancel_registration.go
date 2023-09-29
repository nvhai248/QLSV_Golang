package classregisterbiz

import (
	"context"
	classregistermodel "studyGoApp/modules/classregister/model"
)

type CancelRegistrationStore interface {
	DeleteClassRegister(ctx context.Context, studentId, classId int) error
	FindClassRegister(ctx context.Context, studentId, classId int) (*classregistermodel.Register, error)
}

type cancelRegistrationBiz struct {
	store CancelRegistrationStore
}

func NewCancelRegistrationBiz(store CancelRegistrationStore) *cancelRegistrationBiz {
	return &cancelRegistrationBiz{store: store}
}

func (b *cancelRegistrationBiz) CancelRegistration(ctx context.Context, studentId, classId int) error {
	if err := b.store.DeleteClassRegister(ctx, studentId, classId); err != nil {
		return classregistermodel.ErrorCannotCancelRegistration(err)
	}

	return nil
}
