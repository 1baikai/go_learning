package endpoint

import (
	"context"

	"bproject/entity"
	service "bproject/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// AddUserRequest collects the request parameters for the AddUser method.
type AddUserRequest struct {
	Username string `json:"username"`
	Fraction string `json:"fraction"`
}

// AddUserResponse collects the response parameters for the AddUser method.
type AddUserResponse struct {
	Err error `json:"err"`
}

// MakeAddUserEndpoint returns an endpoint that invokes AddUser on the service.
func MakeAddUserEndpoint(s service.BProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddUserRequest)
		err := s.AddUser(ctx, req.Username, req.Fraction)
		return AddUserResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r AddUserResponse) Failed() error {
	return r.Err
}

// GetUserRequest collects the request parameters for the GetUser method.
type GetUserRequest struct {
	UserID int64 `json:"user_id"`
}

// GetUserResponse collects the response parameters for the GetUser method.
type GetUserResponse struct {
	Res *entity.User `json:"res"`
	Err error        `json:"err"`
}

// MakeGetUserEndpoint returns an endpoint that invokes GetUser on the service.
func MakeGetUserEndpoint(s service.BProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		res, err := s.GetUser(ctx, req.UserID)
		return GetUserResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// AddUser implements Service. Primarily useful in a client.
func (e Endpoints) AddUser(ctx context.Context, username string, fraction string) (err error) {
	request := AddUserRequest{
		Fraction: fraction,
		Username: username,
	}
	response, err := e.AddUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddUserResponse).Err
}

// GetUser implements Service. Primarily useful in a client.
func (e Endpoints) GetUser(ctx context.Context, userID int64) (res *entity.User, err error) {
	request := GetUserRequest{UserID: userID}
	response, err := e.GetUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserResponse).Res, response.(GetUserResponse).Err
}
