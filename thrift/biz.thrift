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


struct Music {
    1: required string musicName;
    2: required string musicAuthor;
    3: required i32 id;
    4: required string musicPath;
    5: required i32 commentId;
}

struct GetMusicListRequest {
    1: string musicName;
    2: string musicAuthor;
}

struct GetMusicListResponse {
    1: Response resp;
    2: list<Music> musics;
}



service BizService {
    LoginReponse Login(1: LoginRequest request)
    LogoutResponse Logout(1: LogoutRequest request)
    SignUpResponse SignUp(1: SignUpRequest request)
    GetMusicListResponse searchMusics(1: GetMusicListRequest request)
}