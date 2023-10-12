package subscriber

import (
	"context"
	"studyGoApp/component"
)

func Setup(ctx component.AppContext) {
	IncreaseStudentCountAfterStudentRegisterToTheClass(ctx, context.Background())
	IncreaseClassCountAfterStudentRegisterToTheClass(ctx, context.Background())
}
