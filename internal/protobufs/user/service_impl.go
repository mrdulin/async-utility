package user

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	http "github.com/mrdulin/grpc-go-cnode/internal/utils/http"
)

var (
	ErrGetUserByLoginname  = errors.New("userServiceImpl: Get user by login name")
	ErrValidateAccessToken = errors.New("userServiceImpl: Validate accessToken")
)

type (
	validateAccessTokenRequestPayload struct {
		AccessToken string `json:"accesstoken"`
	}
)

type userServiceImpl struct {
	HttpClient http.Client
	BaseURL    string
	UnimplementedUserServiceServer
}

func NewUserServiceImpl(httpClient http.Client, baseurl string) *userServiceImpl {
	return &userServiceImpl{HttpClient: httpClient, BaseURL: baseurl}
}

func (svc *userServiceImpl) GetUserByLoginname(ctx context.Context, in *GetUserByLoginnameRequest) (*GetUserByLoginnameResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	endpoint := svc.BaseURL + "/user/" + in.Loginname
	var res GetUserByLoginnameResponse
	err = svc.HttpClient.Get(endpoint, &res)
	if err != nil {
		fmt.Println(err)
		return nil, ErrGetUserByLoginname
	}
	return &res, nil
}
func (svc *userServiceImpl) ValidateAccessToken(ctx context.Context, in *ValidateAccessTokenRequest) (*ValidateAccessTokenResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	url := svc.BaseURL + "/accesstoken"
	var res ValidateAccessTokenResponse
	err = svc.HttpClient.Post(url, &validateAccessTokenRequestPayload{AccessToken: in.Accesstoken}, &res)
	if err != nil {
		fmt.Println(err)
		return nil, ErrValidateAccessToken
	}
	return &res, nil
}
