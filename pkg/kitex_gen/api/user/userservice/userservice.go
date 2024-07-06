// Code generated by Kitex v0.10.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"

	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"

	user "github.com/mutezebra/tiktok/pkg/kitex_gen/api/user"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newUserServiceRegisterArgs,
		newUserServiceRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newUserServiceLoginArgs,
		newUserServiceLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Info": kitex.NewMethodInfo(
		infoHandler,
		newUserServiceInfoArgs,
		newUserServiceInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UploadAvatar": kitex.NewMethodInfo(
		uploadAvatarHandler,
		newUserServiceUploadAvatarArgs,
		newUserServiceUploadAvatarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DownloadAvatar": kitex.NewMethodInfo(
		downloadAvatarHandler,
		newUserServiceDownloadAvatarArgs,
		newUserServiceDownloadAvatarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"TotpQrcode": kitex.NewMethodInfo(
		totpQrcodeHandler,
		newUserServiceTotpQrcodeArgs,
		newUserServiceTotpQrcodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"EnableTotp": kitex.NewMethodInfo(
		enableTotpHandler,
		newUserServiceEnableTotpArgs,
		newUserServiceEnableTotpResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.10.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func infoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceInfoArgs)
	realResult := result.(*user.UserServiceInfoResult)
	success, err := handler.(user.UserService).Info(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceInfoArgs() interface{} {
	return user.NewUserServiceInfoArgs()
}

func newUserServiceInfoResult() interface{} {
	return user.NewUserServiceInfoResult()
}

func uploadAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUploadAvatarArgs)
	realResult := result.(*user.UserServiceUploadAvatarResult)
	success, err := handler.(user.UserService).UploadAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUploadAvatarArgs() interface{} {
	return user.NewUserServiceUploadAvatarArgs()
}

func newUserServiceUploadAvatarResult() interface{} {
	return user.NewUserServiceUploadAvatarResult()
}

func downloadAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceDownloadAvatarArgs)
	realResult := result.(*user.UserServiceDownloadAvatarResult)
	success, err := handler.(user.UserService).DownloadAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceDownloadAvatarArgs() interface{} {
	return user.NewUserServiceDownloadAvatarArgs()
}

func newUserServiceDownloadAvatarResult() interface{} {
	return user.NewUserServiceDownloadAvatarResult()
}

func totpQrcodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceTotpQrcodeArgs)
	realResult := result.(*user.UserServiceTotpQrcodeResult)
	success, err := handler.(user.UserService).TotpQrcode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceTotpQrcodeArgs() interface{} {
	return user.NewUserServiceTotpQrcodeArgs()
}

func newUserServiceTotpQrcodeResult() interface{} {
	return user.NewUserServiceTotpQrcodeResult()
}

func enableTotpHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceEnableTotpArgs)
	realResult := result.(*user.UserServiceEnableTotpResult)
	success, err := handler.(user.UserService).EnableTotp(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceEnableTotpArgs() interface{} {
	return user.NewUserServiceEnableTotpArgs()
}

func newUserServiceEnableTotpResult() interface{} {
	return user.NewUserServiceEnableTotpResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *user.RegisterReq) (r *user.RegisterResp, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Req = req
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *user.LoginReq) (r *user.LoginResp, err error) {
	var _args user.UserServiceLoginArgs
	_args.Req = req
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Info(ctx context.Context, req *user.InfoReq) (r *user.InfoResp, err error) {
	var _args user.UserServiceInfoArgs
	_args.Req = req
	var _result user.UserServiceInfoResult
	if err = p.c.Call(ctx, "Info", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadAvatar(ctx context.Context, req *user.UploadAvatarReq) (r *user.UploadAvatarResp, err error) {
	var _args user.UserServiceUploadAvatarArgs
	_args.Req = req
	var _result user.UserServiceUploadAvatarResult
	if err = p.c.Call(ctx, "UploadAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DownloadAvatar(ctx context.Context, req *user.DownloadAvatarReq) (r *user.DownloadAvatarResp, err error) {
	var _args user.UserServiceDownloadAvatarArgs
	_args.Req = req
	var _result user.UserServiceDownloadAvatarResult
	if err = p.c.Call(ctx, "DownloadAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) TotpQrcode(ctx context.Context, req *user.TotpQrcodeReq) (r *user.TotpQrcodeResp, err error) {
	var _args user.UserServiceTotpQrcodeArgs
	_args.Req = req
	var _result user.UserServiceTotpQrcodeResult
	if err = p.c.Call(ctx, "TotpQrcode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EnableTotp(ctx context.Context, req *user.EnableTotpReq) (r *user.EnableTotpResp, err error) {
	var _args user.UserServiceEnableTotpArgs
	_args.Req = req
	var _result user.UserServiceEnableTotpResult
	if err = p.c.Call(ctx, "EnableTotp", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
