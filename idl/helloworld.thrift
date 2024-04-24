// helloworld.thrift
namespace go buyandship.marketplace.helloworld

struct BaseResp {
    1: string code
    2: string msg
}

struct GreetReq {
    1: string key
}

struct GreetResp {
    1: string val
    255: BaseResp baseResp
}

service GreetService {
    GreetResp Greet(1: GreetReq req)
} 