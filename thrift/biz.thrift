namespace go tCloudMisc.biz.api


struct Response {
    1: required i32 code;
    2: required string msg;
    3: required i32 timestamp;
}

struct LoginRequest {
    1: required string username;
    2: required string password;
}

struct LoginReponse {
    1: required string token;
    2: required Response resp;
}

struct LogoutRequest {
    1: required string username;
}

struct LogoutResponse {
    1: required bool result;
    2: required Response resp;
}

struct SignUpRequest {
    1: required string username;
    2: required string password;
    3: required string email;
}


struct SignUpResponse {
    1: required bool result;
}


service BizService {
    LoginReponse Login(1: LoginRequest request)
    LogoutResponse Logout(1: LogoutRequest request)
    SignUpResponse SignUp(1: SignUpRequest request)
}