package util

import (
	libconstants "github.com/filswan/go-swan-lib/constants"
)

type BasicResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func CreateSuccessResponse(_data interface{}) BasicResponse {
	return BasicResponse{
		Status: libconstants.SWAN_API_STATUS_SUCCESS,
		Data:   _data,
		Code:   SuccessCode,
	}
}

func CreateErrorResponse(code int, errMsg ...string) BasicResponse {
	var msg string
	if len(errMsg) == 0 {
		msg = codeMsg[code]
	} else {
		msg = errMsg[0]
	}
	return BasicResponse{
		Status:  libconstants.SWAN_API_STATUS_FAIL,
		Code:    code,
		Message: msg,
	}
}

const (
	SuccessCode = 200
	ServerError = 500

	JsonError                = 4000
	BadParamError            = 4001
	SpaceSignatureError      = 4002
	SpaceCheckResourcesError = 4003
	SpaceCheckWhiteListError = 4004

	SpaceDeployStatusError  = 6002
	ProofParamError         = 7001
	ProofReadLogError       = 7002
	ProofError              = 7003
	UbiTaskParamError       = 8001
	UbiTaskReadLogError     = 8002
	UbiTaskError            = 8003
	CheckResourcesError     = 9001
	CheckAvailableResources = 9002
)

var codeMsg = map[int]string{
	BadParamError:          "The request parameter is not valid",
	JsonError:              "An error occurred while converting to json",
	ServerError:            "Service failed",
	SpaceSignatureError:    "Verify signature failed",
	SpaceDeployStatusError: "Check deploy status failed",

	ProofReadLogError: "An error occurred while read the log of proof",
	ProofError:        "An error occurred while executing the calculation task",

	CheckResourcesError:      "An error occurred while check resources available",
	CheckAvailableResources:  "No resources available",
	SpaceCheckWhiteListError: "This cp does not accept tasks from wallet addresses outside the whitelist",
}
