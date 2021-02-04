package server

import (
	"context"
	pb "model-service/api/v1/proto"
	"model-service/server/e"
)

type modelService struct {
	pb.UnimplementedModelServiceServer
}

//NewModelService ///
func NewModelService() *modelService {
	return &modelService{}
}

func (m modelService) Publish(ctx context.Context, r *pb.PublishRequest) (*pb.Response, error) {
	code := PublishModel(r.GetModelId(), r.GetProtocol(), r.GetServiceName(), r.GetParams())
	return &pb.Response{
		Code:     code,
		Messsage: e.GetMessage(code),
	}, nil
}

func (m modelService) GetServiceProtocol(ctx context.Context, r *pb.Null) (*pb.ProtocolResponse, error) {
	code, data := GetSupportPushlishProtocol()
	return &pb.ProtocolResponse{
		Code:     code,
		Messsage: e.GetMessage(code),
		Data:     data,
	}, nil
}

func (m modelService) GetServerParameters(context.Context, *pb.Null) (*pb.ParamsResponse, error) {
	code, data := GetServerParameInfo()
	return &pb.ParamsResponse{
		Code:     code,
		Messsage: e.GetMessage(code),
		Data:     data,
	}, nil
}
