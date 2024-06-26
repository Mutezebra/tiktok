namespace go api.user

include "base.thrift"

struct UserInfo {
    1: optional string ID (go.tag="json:\"id,omitempty\"")
    2: optional string UserName (go.tag="json:\"user_name,omitempty\"")
    3: optional string Email (go.tag="json:\"email,omitempty\"")
    4: optional i8 Gender (go.tag="json:\"gender,omitempty\"")
    5: optional string Avatar (go.tag="json:\"avatar,omitempty\"")
    6: optional i32 Fans (go.tag="json:\"fans,omitempty\"")
    7: optional i32 Follows (go.tag="json:\"follows,omitempty\"")
    8: optional i64 CreateAt (go.tag="json:\"create_at,omitempty\"")
    9: optional i64 UpdateAt (go.tag="json:\"update_at,omitempty\"")
    10: optional bool TotpStatus (go.tag="json:\"totp_status,omitempty\"")
}

struct RegisterReq {
    1: optional string UserName (go.tag="form:\"user_name,required\"")
    2: optional string Email    (go.tag="form:\"email,required\"")
    3: optional string Password (go.tag="form:\"password,required\"")
}

struct RegisterResp {
}


struct LoginReq {
    1: optional string UserName (go.tag="form:\"user_name,required\"")
    2: optional string password (go.tag="form:\"password,required\"")
    3: optional string OTPCode (go.tag="form:\"otp_code\"")
}

struct LoginResp {
    1: optional string AccessToken  (go.tag="json:\"access_token,omitempty\"")
    2: optional string RefreshToken (go.tag="json:\"refresh_token,omitempty\"")
}


struct InfoReq {
    1: optional i64 UID
    2: optional string UserName (go.tag="form:\"user_name\"")
}

struct InfoResp {
    1: optional UserInfo Data (go.tag="json:\"data\"")
}


struct UploadAvatarReq {
    1: optional i64 UID
    3: optional binary Avatar
    4: optional string FileName
}

struct UploadAvatarResp {
}


struct DownloadAvatarReq {
    1: optional i64 UID
}

struct DownloadAvatarResp {
    1: optional string URL (go.tag="json:\"url,omitempty\"")
}


struct TotpQrcodeReq {
    1: optional i64 UID
    2: optional string UserName (go.tag="json:\"user_name,omitempty\"")
}

struct TotpQrcodeResp {
    2: optional string Qrcode (go.tag="json:\"qrcode,omitempty\"")
}


struct EnableTotpReq {
    1: optional string Code (go.tag="form:\"code,required\"")
    2: optional i64 UID
}

struct EnableTotpResp {
}


service UserService {
    RegisterResp Register(1: RegisterReq req)
    LoginResp Login(1: LoginReq req)
    InfoResp Info(1: InfoReq req)
    UploadAvatarResp UploadAvatar(1: UploadAvatarReq req)
    DownloadAvatarResp DownloadAvatar(1: DownloadAvatarReq req)
    TotpQrcodeResp TotpQrcode(1: TotpQrcodeReq req)
    EnableTotpResp EnableTotp(1: EnableTotpReq req)
}
