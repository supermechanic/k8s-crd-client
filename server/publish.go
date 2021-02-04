package server

import (
	"fmt"
	"model-service/core"
	"model-service/core/seldon"
	"model-service/models"
	"model-service/server/e"
	"strings"
)

//PublishModel ///
func PublishModel(modelID int32, protocol int32, serviceName string, params map[int32]int32) (code int32) {
	if modelID <= 0 {
		return e.ERROR_BAD_PARAM
	}
	modelInfo := &models.Model{
		ModelID: modelID,
	}
	err := modelInfo.GetModelByID()
	if err != nil {
		fmt.Println(err.Error())
		return e.ERROR_DB_NOT_FIND
	}
	deployName := modelInfo.ModelName + "_" + modelInfo.ModelVer
	modelURL := modelInfo.ModelPath
	var serverType seldon.ServerType
	if strings.Contains(modelInfo.Cti.FullName, "sklearn") {
		serverType = seldon.SKLEARN_SERVER
	} else if strings.Contains(modelInfo.Cti.FullName, "tf") {
		serverType = seldon.TENSORFLOW_SERVER
	}
	predict := &models.Predict{
		ModelID: modelID,
	}
	switch protocol {
	case int32(core.SeldonCoreV1):
		if err := seldon.CreateSeldonDeployV1(deployName, modelURL, params, serverType); err != nil {
			fmt.Println(err.Error())
			return e.ERROR_INTERNAL
		}
		if err := predict.Add(); err != nil {
			return e.ERROR_INTERNAL
		}
		return e.OK
	default:
		return e.ERROR_NOT_SUPPORT
	}
}
