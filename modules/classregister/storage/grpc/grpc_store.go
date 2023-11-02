package classregistrationgrpc

import (
	"context"
	classregisterstorage "studyGoApp/modules/classregister/storage"
	"studyGoApp/proto"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCServer struct {
	db *sqlx.DB
	proto.UnimplementedClassRegistrationServiceServer
}

func NewgRPCServer(db *sqlx.DB) *gRPCServer {
	return &gRPCServer{db: db}
}

func (r *gRPCServer) GetClassRegistrationStat(ctx context.Context, request *proto.ClassRegistrationStatRequest) (*proto.ClassRegistrationStatResponse, error) {
	storage := classregisterstorage.NewSQLStore(r.db)
	ids := make([]int, len(request.ClassIds))

	for i := range ids {
		ids[i] = int(request.ClassIds[i])
	}
	result, err := storage.GetNumberOfStudentRegisteredInClass(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GetClassRegistrationStat has something error %s", err)
	}

	mapRs := make(map[int32]int32)

	for k, v := range result {
		mapRs[int32(k)] = int32(v)
	}

	return &proto.ClassRegistrationStatResponse{Result: mapRs}, nil
}
