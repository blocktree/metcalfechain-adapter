# metcalfechain-adapter

本项目适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建M.ini文件，编辑如下内容：

```ini
# ws - wsAPI  rpc - nodeAPI
apiChoose = "rpc"
# node api url
nodeAPI = "http://ip:port"
# websocket api url
wsAPI = "ws://ip:port"
# fixed Fee in smallest unit
fixedFee = 100
# reserve amount in smallest unit
reserveAmount = 100
# ignore reserve amount
ignoreReserve = true
# register fee in sawi
registerFee = 100
# last ledger sequence number
lastLedgerSequenceNumber = 20
# memo type
memoType = "withdraw"
# memo format
memoFormat = "text/plain"
# which feild of memo to scan
memoScan = "MemoData"
# Cache data file directory, default = "", current directory: ./data
dataDir = "/home/golang/data"
```
