package server

import (
	pb "model-service/api/v1/proto"
	"model-service/core"
	"model-service/core/seldon"
	"model-service/server/e"
)

//GetSupportPushlishProtocol ///
func GetSupportPushlishProtocol() (int32, map[string]int32) {
	return e.OK, core.GetProtocols()
}

//GetServerParameInfo //
func GetServerParameInfo() (int32, map[int32]*pb.Server) {
	result := make(map[int32]*pb.Server)
	serverParames := seldon.GetOptServerParams()
	for severType, serverOpt := range serverParames {
		param := map[int32]*pb.Parameter{}
		for k, v := range serverOpt.Params {
			param[k] = &pb.Parameter{
				Name:   v.Name,
				Type:   v.Type,
				Values: v.Values,
			}
		}
		temp := &pb.Server{
			Name:  serverOpt.Name,
			Param: param,
		}
		result[int32(severType)] = temp
	}
	return e.OK, result
}
