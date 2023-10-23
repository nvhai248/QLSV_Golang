package subscriber

import (
	"context"
	"log"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/component/asyncjob"
	"studyGoApp/pubsub"
	"studyGoApp/skio"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx   component.AppContext
	rtEngine skio.RealtimeEngine
}

func NewEngine(appContext component.AppContext, rtEngine skio.RealtimeEngine) *consumerEngine {
	return &consumerEngine{appCtx: appContext, rtEngine: rtEngine}
}

func (engine *consumerEngine) Start() error {
	engine.startSubTopic(
		common.TopicStudentRegisterToTheClass,
		true,
		RunIncreaseStudentCountAfterStudentRegisterToTheClass(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicStudentRegisterToTheClass,
		true,
		RunIncreaseClassCountAfterStudentRegisterToTheClass(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicStudentCancelRegistration,
		true,
		RunDecreaseClassCountAfterStudentRegisterToTheClass(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicStudentCancelRegistration,
		true,
		RunDecreaseStudentCountAfterStudentRegisterToTheClass(engine.appCtx),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("Setup consumer for: ", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			mgs := <-c

			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], mgs)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
