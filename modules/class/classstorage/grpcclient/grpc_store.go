package classgrpcclient

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/proto"
)

type gRPCClient struct {
	client proto.ClassRegistrationServiceClient
}

func NewgRPCClient(client proto.ClassRegistrationServiceClient) *gRPCClient {
	return &gRPCClient{client: client}
}

func (c *gRPCClient) GetNumberOfStudentRegisteredInClass(ctx context.Context, ids []int) (map[int]int, error) {
	classIds := make([]int32, len(ids))

	for i := range classIds {
		classIds[i] = int32(ids[i])
	}

	res, err := c.client.GetClassRegistrationStat(ctx, &proto.ClassRegistrationStatRequest{ClassIds: classIds})

	if err != nil {
		return nil, common.ErrDB(err)
	}

	mapRes := make(map[int]int)

	for k, v := range res.Result {
		mapRes[int(k)] = int(v)
	}

	return mapRes, nil
}
