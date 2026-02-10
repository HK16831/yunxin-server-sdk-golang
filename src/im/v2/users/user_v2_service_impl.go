package users

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// UpdateUser 更新用户信息
func (s *UserV2ServiceImpl) UpdateUser(req *UpdateUserRequestV2) (*core.Result[*UpdateUserResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.PATCH, UpdateUser, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	resultT, err := util.ConvertWithClass[UserInfo](apiResponse)
	if err == nil {
		return core.NewResult[*UpdateUserResponseV2](resultT.GetEndpoint(), resultT.GetCode(), resultT.GetTraceId(), resultT.GetMsg(), &UpdateUserResponseV2{UserInfo:resultT.Response}), nil
	}
	return nil, err
}

// GetUser 获取用户信息
func (s *UserV2ServiceImpl) GetUser(req *GetUserRequestV2) (*core.Result[*GetUserResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.GET, GetUser, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	resultT, err := util.ConvertWithClass[UserInfo](apiResponse)
	if err == nil {
		return core.NewResult[*GetUserResponseV2](resultT.GetEndpoint(), resultT.GetCode(), resultT.GetTraceId(), resultT.GetMsg(), &GetUserResponseV2{UserInfo:resultT.Response}), nil
	}
	return nil, err
}

// BatchGetUsers 批量获取用户信息
func (s *UserV2ServiceImpl) BatchGetUsers(req *BatchGetUsersRequestV2) (*core.Result[*BatchGetUsersResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, BatchGetUsers, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchGetUsersResponseV2](apiResponse)
}

// GetUsersOnlineStatus 获取用户在线状态
func (s *UserV2ServiceImpl) GetUsersOnlineStatus(req *GetUserOnlineStatusRequestV2) (*core.Result[*GetUserOnlineStatusResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, GetUsersOnlineStatus, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*GetUserOnlineStatusResponseV2](apiResponse)
}
