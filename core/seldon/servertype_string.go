// Code generated by "stringer -type=ServerType"; DO NOT EDIT.

package seldon

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MLFLOW_SERVER-0]
	_ = x[SKLEARN_SERVER-1]
	_ = x[TENSORFLOW_SERVER-2]
	_ = x[PYTORCH_SERVER-3]
	_ = x[TRITON_SERVER-4]
	_ = x[XGBOOST_SERVER-5]
}

const _ServerType_name = "MLFLOW_SERVERSKLEARN_SERVERTENSORFLOW_SERVERPYTORCH_SERVERTRITON_SERVERXGBOOST_SERVER"

var _ServerType_index = [...]uint8{0, 13, 27, 44, 58, 71, 85}

func (i ServerType) String() string {
	if i < 0 || i >= ServerType(len(_ServerType_index)-1) {
		return "ServerType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ServerType_name[_ServerType_index[i]:_ServerType_index[i+1]]
}
