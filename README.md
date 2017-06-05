# fabric-ca
ca提供用户身份注册，发放ECert证书，证书的更新和撤销 ，目前Tcert证书没有看到，应该是在
Fabric内部做交易的时候申请的证书。
######fabric-ca 初始化时，执行fabric-ca-server init -b admin:adminpw
其中admin是client链接时的用户名，adminw是对应的密码
在使用sdk接口时：
首先第一步调用：
key, cert, err := caClient.Enroll("admin", "adminpw")
申请一个admin用户的证书和私钥

接着申请注册其他用户：
registerRequest := fabricCAClient.RegistrationRequest{Name: userName, Type: "user", Affiliation: "org1.department1"}
enrolmentSecret, err := caClient.Register(adminUser, &registerRequest)
会使用admin用户的私钥签名

接着enroll 登记用户：
ekey, ecert, err := caClient.Enroll(userName, enrolmentSecret)
申请用户的ECert证书。


思考：
在第一步时，使用admin账户链接应该是申请一个ca的子证书，使用这个子证书签发用户register信息，保证用户注册时信息的安全，用户register信息发送给CA，CA会判断用户名是否已经存在，
如果没有则返回用户的密码
其中用户在发送register信息时，Affiliation信息必须在ca的配置文件Affiliation已经配置，否则申请会被拒绝。
注意：如果用户向ca申请register信息，在发送的信息中含有用户密码，则ca向用户发送密码，ca仅保存用户register信息，并返回OK

用户register之后，就开始enroll阶段，次阶段是用户真正的申请ECert证书
ekey, ecert, err := caClient.Enroll(用户名, 密码（自己申请的或者ca返回的密码）)
此时返回私钥和Ecert证书信息。

其他细节信息查看：
http://8btc.com/article-4515-1.html

fabric 1.0
目前镜像打包：
https://github.com/hyperledger/fabric
 Release 1.0.0-alpha 版本，
一次执行：
make gotools 下载go的代码配置工具
make linter 执行代码检查，并下载chaintool工具
make peer 编译peer节点 
make orderer 编译共识节点
make 制作peer orderer镜像
其中gotools \linter过程中会遇到一些链接外国google代码的问题，需要翻墙出去，或者直接
下载google的依赖包，编译后放到fabric对应的目录下


root@peer0:/go/src/github.com/hyperledger/fabric# peer chaincode deploy -n test_cc -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02 -c '{"Args":["init","a","100","b","200"]}'
...
http://blog.csdn.net/yeasy/article/details/54928343
网上的镜像install智能合约时，应为不支持deploy 替换为install ，因为有go imnport代码
依赖问题，怀疑的网上的镜像环境有问题，所以重新找一个官方的镜像试一试；






