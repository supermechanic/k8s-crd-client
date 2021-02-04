package seldon

import "errors"

//ParameterOpt ///
type ParameterOpt struct {
	Name   string
	Type   string
	Values map[int32]string
}

//Parameter ///
type Parameter struct {
	Name  string
	Type  string
	Value string
}

//ServerOpt ///
type ServerOpt struct {
	Name   string
	Params map[int32]ParameterOpt
}

var innerServerInfo = map[ServerType]ServerOpt{
	SKLEARN_SERVER: {
		Name: SKLEARN_SERVER.String(),
		Params: map[int32]ParameterOpt{
			1: {
				Name: "method",
				Type: "STRING",
				Values: map[int32]string{
					1: "predict",
					2: "predict_prob:default",
					3: "dicision_function",
				},
			},
		},
	},
	TENSORFLOW_SERVER: {
		Name:   TENSORFLOW_SERVER.String(),
		Params: map[int32]ParameterOpt{},
	},
}

//GetOptServerParams ///
func GetOptServerParams() map[ServerType]ServerOpt {
	return innerServerInfo
}

//GetServerParams ///
func GetServerParams(sType ServerType, paramValues map[int32]int32) ([]Parameter, error) {
	var params []Parameter
	serverInfo, ok := innerServerInfo[sType]
	if !ok {
		return nil, errors.New("request not supported server type")
	}
	for p, v := range paramValues {
		var parameter Parameter
		if param, ok := serverInfo.Params[p]; !ok {
			return nil, errors.New("request not supported parameter type")
		} else {
			parameter.Name = param.Name
			parameter.Type = param.Type
			if v, ok := param.Values[v]; ok {
				parameter.Value = v
			} else {
				return nil, errors.New("request not supported parameter value type")
			}
		}
		params = append(params, parameter)
	}
	return params, nil
}
