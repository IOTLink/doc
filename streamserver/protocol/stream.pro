syntax = "proto3";
package stream;

message RegisterReq {
    string user = 1;
    string pwd = 2;
}

message RegisterReply {
    string  message = 1;
    RegisterInfo info = 2;
}


message RegisterInfo {
    string appid = 1;
    string appkey = 2;
}


message Asset {
    string userid = 1;
    int32  value = 2;

}

message Transaction {
    string ownerid = 1;
    string receiverid = 2;
    int32  value = 3;  
}

message MsgReply{
    string message = 1;
}

service StreamServer{ 
    rpc RegisterClient(RegisterReq)  returns (RegisterReply) {}
    rpc InitAsset(Asset)  returns (MsgReply) {}
    rpc DealTransaction(Transaction) returns (MsgReply) {}
    rpc QueryAsset(Asset) returns(Asset) {}
}















