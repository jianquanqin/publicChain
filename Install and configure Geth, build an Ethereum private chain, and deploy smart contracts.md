# 以太坊私有链

The article introduced how to install and use the client of ethereum - Geth firstly. then talking about building a private chain, at the end, took an example to illustrate how to deploy a smart contract and connect all nodes

## 1 以太坊私有链

**为什么要用以太坊的私有链？**

在以太坊的公链上部署智能合约、发起交易需要花费以太币，这种门槛对项目的前期测试不友好，通过修改配置，可以在本季搭建一套以太坊私有链。私有链和以太坊并没有关系，所以不用同步公有链庞大的数据和购买以太币，这能够很好的满足项目的开发和测试需求，并且开发好的智能合约可以很容易的切换接口部署到以太坊公有链上

**安装homebrew**（以mac为例）

在终端中运行以下命令：

```
 /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

但是此时可能发生443连接错误，请修改以下端口为自己本机的端口，然后继续输入上一步骤命令，科学上网

```
export https_proxy=http://127.0.0.1:1087 http_proxy=http://127.0.0.1:1087 all_proxy=socks5://127.0.0.1:1087
```

输入密码，静待一会儿即可完成安装

**安装ethereum源码**

```
git clone https://github.com/ethereum/go-ethereum.git
```

**从源码构建Geth**

```
 $ cd go-ethereum
 $ make geth
```

安装好了，请注意不要随意运行同步，因为数据量实在太大了

通过以下命令查看客户端版本

```
./build/bin/geth version

Geth
Version: 1.10.17-unstable
Git Commit: e0e8bf31c5d44f7de33ce774b221debf2c42256c
Git Commit Date: 20220323
Architecture: amd64
Go Version: go1.17.5
Operating System: darwin
GOPATH=
GOROOT=go
```

**安装Solidity语言**

```
brew install solidity
```

会有点耗时，请耐心等待

### 1.1 **搭建私有节点**

因为公链区块数量较多，同步耗时久，我们现在搭建一条只属于我们自己的私有链，用于测试我们实现的各项功能

**1.创建文件夹存储数据**

```
mkdir privatechain
cd privatechain
```

**2.创建节点**

在创建节点之前我们需要使用一个json文件来配置一些初始参数，json文件内容如下，并将其保存在privatechain文件夹下

```
{
  "config": {
    "chainId": 111,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "ethash": {}
  },
  "nonce": "0x0",
  "timestamp": "0x5ddf8f3e",
  "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "gasLimit": "0x47b760",
  "difficulty": "0x00002",
  "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "coinbase": "0x0000000000000000000000000000000000000000",
  "alloc": { },
  "number": "0x0",
  "gasUsed": "0x0",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
}
```

**3.执行初始化命令**

```
geth --datadir data1 init genesis.json
```

**4.启用节点**

注意：在这里我们使用了一系列的命令，命令的具体含义请读者参考官方教程定义

https://geth.ethereum.org/docs/interface/command-line-options

现在的目的是通过一台机器创建多个节点，并将其连接起来成为网络，在这里使用同一个本机地址：192.168.31.144 和不同的服务器监听端口和网络监听端口（--http.port  --port 3366），也就是说只要每个节点的这两个端口不同，节点之间就不会冲突，也就能够实现互联，另外，在这里我们使用了--nodiscover参数，这意味着我们只能通过手动实现节点连接，下文我们也会实现这一点

```
geth --datadir "/Users/qinjianquan/privatechain/data1" --http --http.api "eth,web3,miner,admin,personal,net" --http.corsdomain "*" --nodiscover --networkid 111111 --http.addr 192.168.31.144  --http.port 9049 --port 3366 --allow-insecure-unlock
```

此时data1文件夹会生成以下文件，keystore用来存放私钥，geth.ipc启动节点，geth存放了一些链上数据

```
% cd data1
% ls
geth		geth.ipc	keystore
```

**5.连接节点**

重新打开一个终端窗口，进入privatechain文件夹的data1文件中执行以下命令

```
geth attach ipc:./geth.ipc
```

如下是连接成功的信息

```
Welcome to the Geth JavaScript console!

instance: Geth/v1.10.16-stable/darwin-amd64/go1.17.6
at block: 0 (Thu Nov 28 2019 17:11:26 GMT+0800 (CST))
 datadir: /Users/qinjianquan/privatechain/data2
 modules: admin:1.0 debug:1.0 eth:1.0 ethash:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0

To exit, press ctrl-d or type exit
```

非常棒，以上就是一个私有链节点搭建的全部过程

我们现在就可以创建钱包、部署合约等

### 1.2 钱包和挖矿

现在我们来在geth console中执行一些基本操作

**1. 创建钱包地址**

```
personal.newAccount()
Passphrase: 
Repeat passphrase: 
"0xf980455ca20d7b9a05d72e5439cc52f5d25f0f12"
```

**2.查看所有钱包地址**

此处我们又新建了一个地址，所以现在有两个

```
personal.listAccounts
["0xf980455ca20d7b9a05d72e5439cc52f5d25f0f12", "0x4a4eacfb8cda324418d4afc515c67c266ebd4df1"]
```

**3.重命名地址**

3.1 将某个地址定义为coinbase地址

```
miner.setEtherbase("0xf980455ca20d7b9a05d72e5439cc52f5d25f0f12")
true
```

3.2 将某个地址重命名,Accountlists 其实是一个数组，可以通过下标访问

```
receiver = web3.eth.accounts[1]
"0x4a4eacfb8cda324418d4afc515c67c266ebd4df1"
web3.eth.getBalance(receiver) 
0
```

**4.查询余额**

```
web3.eth.getBalance(eth.coinbase)
0
```

```
web3.eth.getBalance("0xf980455ca20d7b9a05d72e5439cc52f5d25f0f12")
0
```

**5.查看节点信息**

```
admin.nodeInfo
{
  enode: "enode://6aff3fb4d6d5be9d9898312baa1a5668c8fb3a5f50cb68f93523b543ea491f73296d8f648806cb72044115aa806b455772bc36f7c976d0b3da99b05b2e6131de@127.0.0.1:3000?discport=0",
  enr: "enr:-Jy4QOeCeWb5CaKeUST32uZKPrrK-OmGfK4-laAhZsAV4W6DYsEMbBbCKSedlJkH8_dd0vSs3VWfdsLa4aDPv4C49GqGAYAddmj8g2V0aMfGhJPR68WAgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQJq_z-01tW-nZiYMSuqGlZoyPs6X1DLaPk1I7VD6kkfc4RzbmFwwIN0Y3CCC7g",
  id: "a2c158b19b4e5400fec0de28b9d0066653ab209abc7d5bd04e6413fa1fe9a028",
  ip: "127.0.0.1",
  listenAddr: "[::]:3000",
  name: "Geth/v1.10.16-stable/darwin-amd64/go1.17.6",
  ports: {
    discovery: 0,
    listener: 3000
  },
  protocols: {
    eth: {
      config: {
        byzantiumBlock: 0,
        chainId: 666,
        constantinopleBlock: 0,
        eip150Block: 0,
        eip150Hash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        eip155Block: 0,
        eip158Block: 0,
        ethash: {},
        homesteadBlock: 0,
        istanbulBlock: 0,
        petersburgBlock: 0
      },
      difficulty: 2,
      genesis: "0xd3d6bb893a6e274cab241245d5df1274c58d664fbb1bfd6e59141c2e0bc5304a",
      head: "0xd3d6bb893a6e274cab241245d5df1274c58d664fbb1bfd6e59141c2e0bc5304a",
      network: 111111
    },
    snap: {}
  }
}
```

**6.查看节点连接情况**

```
web3.net.peerCount
0
```

**7.挖矿**

所挖的奖励默认会打入coinbase账户

```
miner.start()
null
```

```
INFO [04-12|21:45:58.331] 🔗 block reached canonical chain          number=36 hash=20e69e..b09a2b
INFO [04-12|21:45:58.331] 🔨 mined potential block                  number=43 hash=07b4e0..029551
INFO [04-12|21:45:58.331] Commit new sealing work                  number=44 sealhash=054209..b97179 uncles=0 txs=0 gas=0 fees=0 elapsed="143.977µs"
INFO [04-12|21:45:58.331] Commit new sealing work                  number=44 sealhash=054209..b97179 uncles=0 txs=0 gas=0 fees=0 elapsed="299.329µs"
INFO [04-12|21:45:58.944] Looking for peers                        peercount=1 tried=0 static=1
INFO [04-12|21:45:59.817] Successfully sealed new block            number=44 sealhash=054209..b97179 
...
```

```
miner.stop()
null
```

现在查看coinbase账户余额，发现有钱了

```
web3.eth.getBalance(eth.coinbase)
730000000000000000000
```

### 1.3 部署智能合约

了解钱包与挖矿的基本内容之后，我们来尝试部署一个合约到私有节点上

实际上我们已经启动了一个p2p网络，第一个节点可以看作一个服务器，而上述连接的终端可以看作是一个客户端

这部分的内容我们要配合remix编译器来操作

**1.编写合约**

我们在remix中编译一个简单的智能合约，关于solidity语言的基本语法，请参考鄙人的博客：https://blog.csdn.net/weixin_51487151/category_11681449.html?spm=1001.2014.3001.5482

https://remix.ethereum.org

```
pragma solidity ^0.4.18;

contract test{
    function multiply(uint a)public view returns(uint d){
        return a*7;
    }
}
```

然后我们在remix中编译上述代码，在编译详情里可以看到字节码和ABI信息

```
{
	--
	"object": "608060405234801561001057600080fd5b5060bb8061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063c6888fa1146044575b600080fd5b348015604f57600080fd5b50606c600480360381019080803590602001909291905050506082565b6040518082815260200191505060405180910390f35b60006007820290509190505600a165627a7a723058209135a65fdddd7be677810243db99bc4cbe46fcf74ee4ce1a3a8cc7fdbab004ef0029",
	--
}
```

ABI文件（它实际上是一份json格式的文件）

```
[
	{
		"constant": true,
		"inputs": [
			{
				"name": "a",
				"type": "uint256"
			}
		],
		"name": "multiply",
		"outputs": [
			{
				"name": "d",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]
```

**2.通过abi创建合约对象**

现将json字符串转义成字符串，并在geth控制台中将其赋值给新变量

http://www.bejson.com

```
abi = JSON.parse('[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"multiply\",\"outputs\":[{\"name\":\"d\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]')
```

```
myContract = web3.eth.contract(abi)
```

**3.预估手续费**

注意字节码前面要加上0x

```
bytecode = "0x608060405234801561001057600080fd5b5060bb8061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063c6888fa1146044575b600080fd5b348015604f57600080fd5b50606c600480360381019080803590602001909291905050506082565b6040518082815260200191505060405180910390f35b60006007820290509190505600a165627a7a723058209135a65fdddd7be677810243db99bc4cbe46fcf74ee4ce1a3a8cc7fdbab004ef0029"
```

```
> web3.eth.estimateGas({data: bytecode})
93495
```

**4.解锁账户**

```
personal.unlockAccount(eth.coinbase)
Unlock account 0x5540f9b4470d6832284774493602d1e6482cc995
Passphrase: 
true
```

**5.部署合约**

```
contractInstance = myContract.new({data: bytecode,gas: 1000000, from: eth.coinbase}, function(e, contract){
  if(!e){
    if(!contract.address){
      console.log("Contract transaction send: Transaction Hash: "+contract.transactionHash+" waiting to be mined...");
    }else{
      console.log("Contract mined! Address: "+contract.address);
      console.log(contract);
    }
  }else{
    console.log(e)
  }
})
```

如下则说明合约创建成功

```
Contract transaction send: Transaction Hash: 0x0039afc9833eb58bdc97c430b2e0b8dfc5c397a2c302cba6ef35394c456be740 waiting to be mined...
{
  abi: [{
      constant: true,
      inputs: [{...}],
      name: "multiply",
      outputs: [{...}],
      payable: false,
      stateMutability: "view",
      type: "function"
  }],
  address: undefined,
  transactionHash: "0x0039afc9833eb58bdc97c430b2e0b8dfc5c397a2c302cba6ef35394c456be740"
}
> Contract mined! Address: 0x63f91c32529cdf94446a5ac02ea36e2bff929204
[object Object]
```

账户余额不足需要先挖矿

**6.检查一下合约**

```
> contractInstance
{
  abi: [{
      constant: true,
      inputs: [{...}],
      name: "multiply",
      outputs: [{...}],
      payable: false,
      stateMutability: "view",
      type: "function"
  }],
  address: "0x63f91c32529cdf94446a5ac02ea36e2bff929204",
  transactionHash: "0x0039afc9833eb58bdc97c430b2e0b8dfc5c397a2c302cba6ef35394c456be740",
  allEvents: function bound(),
  multiply: function bound()
}
```

**7.调用合约方法**

每调用一次只能合约，程序就会执行一次

```
> contractInstance.multiply("6",{from:eth.coinbase})
42
> contractInstance.multiply("8",{from:eth.coinbase})
56
> contractInstance.multiply("11",{from:eth.coinbase})
77
```

**8.查询交易信息（通过交易哈希）**

```
eth.getTransactionReceipt("0xf94f0f21cec364ef5bb865f513169f912ba4f2775aa7afe065361679893dacad")
```

```
{
  blockHash: "0xeec2189b8b724a6da5d848ba4df776de2abb6d131c215cc4e56b017ed5d4f1e9",
  blockNumber: 695,
  contractAddress: "0x1141f70f401aad06ecb98c174c60c0a3d54aa7c9",
  cumulativeGasUsed: 186990,
  effectiveGasPrice: 1000000000,
  from: "0x5540f9b4470d6832284774493602d1e6482cc995",
  gasUsed: 93495,
  logs: [],
  logsBloom: "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  status: "0x1",
  to: null,
  transactionHash: "0xf94f0f21cec364ef5bb865f513169f912ba4f2775aa7afe065361679893dacad",
  transactionIndex: 1,
  type: "0x0"
}
```

**转账和查看转账信息**

```
eth.sendTransaction({from:eth.coinbase,to:eth.accounts[0],value:web3.toWei(1,"ether")})
"0x84657f918f0c316501a8e95219b571f48ab4b65687dc53c2ae7aab6c36c4925b"
//要通过挖矿打包
> web3.eth.getBalance(web3.eth.accounts[0])
1000000000000000000
```

```
eth.getTransactionReceipt("0x84657f918f0c316501a8e95219b571f48ab4b65687dc53c2ae7aab6c36c4925b")
```

```
{
  blockHash: "0xc9268af2c2d62cc226aa2a8dc71d1cd7d8645298aa990176d5460dbd0c4d4836",
  blockNumber: 902,
  contractAddress: null,
  cumulativeGasUsed: 21000,
  effectiveGasPrice: 1000000000,
  from: "0x5540f9b4470d6832284774493602d1e6482cc995",
  gasUsed: 21000,
  logs: [],
  logsBloom: "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  status: "0x1",
  to: "0xf980455ca20d7b9a05d72e5439cc52f5d25f0f12",
  transactionHash: "0x84657f918f0c316501a8e95219b571f48ab4b65687dc53c2ae7aab6c36c4925b",
  transactionIndex: 0,
  type: "0x0"
}
```

转账也没问题，非常棒，以上就是在私链上部署和测试合约的所有内容

## 2 以太坊多节点连接

以上就是合约部署的全部内容，现在我们将孤立的节点，连成网络

通过如下代码获取本机IP地址

```
ifconfig | grep netmask | awk '{print $2}'
```

```
127.0.0.1
192.168.31.144
```

**1.创建并启用三个节点**

初始化

```
geth --datadir data1 init genesis.json
```

```
geth --datadir data2 init genesis.json
```

```
geth --datadir data3 init genesis.json
```

**2.启用节点**

此时请使用不同的服务监听端口和网络监听端口

```
geth --datadir "/Users/qinjianquan/privatechain/data1" --http --http.api "eth,web3,miner,admin,personal,net" --http.corsdomain "*" --nodiscover --networkid 111111 --http.addr 192.168.31.144  --http.port 8989 --port 3000 --allow-insecure-unlock
```

```
geth --datadir "/Users/qinjianquan/privatechain/data2" --http --http.api "eth,web3,miner,admin,personal,net" --http.corsdomain "*" --nodiscover --networkid 111111 --http.addr 192.168.31.144  --http.port 9049 --port 3333 --allow-insecure-unlock
```

```
geth --datadir "/Users/qinjianquan/privatechain/data3" --http --http.api "eth,web3,miner,admin,personal,net" --http.corsdomain "*" --nodiscover --networkid 111111 --http.addr 192.168.31.144  --http.port 9999 --port 3366 --allow-insecure-unlock
```

**3.连接节点**

请分别进入到data1,data2,data3文件夹中执行如下命令

```
geth attach ipc:./geth.ipc
```

**4.查看节点信息**

```
admin.nodeInfo
{
  enode: "enode://efee81ef4b7730991d6599635f00afc66b7114e3174f7b9e1384ae1f2e06a624523c02ab9706116fb791d3fad735f48fb6c228f732fdbf481b493d2ca32c24f5@127.0.0.1:3000?discport=0",
  enr: "enr:-Jy4QMZ32-F4Th9NqufPZiAMYTepbFWbnrc_r-M4Fnhhz2dYMudesczRKS6X8CUbhgBKK4ezAsHTyuPN0dCdYLvG9IuGAYAeQRf-g2V0aMfGhJPR68WAgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQPv7oHvS3cwmR1lmWNfAK_Ga3EU4xdPe54ThK4fLgamJIRzbmFwwIN0Y3CCC7g",
  id: "d0d497d222b515c81cf34845f207bc47454446c44e971ffcd39dfbe4a0e1986d",
  ip: "127.0.0.1",
  listenAddr: "[::]:3000",
  name: "Geth/v1.10.16-stable/darwin-amd64/go1.17.6",
  ports: {
    discovery: 0,
    listener: 3000
  },
  protocols: {
    eth: {
      config: {
        byzantiumBlock: 0,
        chainId: 666,
        constantinopleBlock: 0,
        eip150Block: 0,
        eip150Hash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        eip155Block: 0,
        eip158Block: 0,
        ethash: {},
        homesteadBlock: 0,
        istanbulBlock: 0,
        petersburgBlock: 0
      },
      difficulty: 2,
      genesis: "0xd3d6bb893a6e274cab241245d5df1274c58d664fbb1bfd6e59141c2e0bc5304a",
      head: "0xd3d6bb893a6e274cab241245d5df1274c58d664fbb1bfd6e59141c2e0bc5304a",
      network: 111111
    },
    snap: {}
  }
}
```

**5.查看节点连接情况**

经过查看，确实三个节点都是独立的，未相互连接

```
> web3.net.peerCount
0
```

**6.连接节点**

将每一个节点的encode代码添加至其他节点

```
admin.addPeer("enode://efee81ef4b7730991d6599635f00afc66b7114e3174f7b9e1384ae1f2e06a624523c02ab9706116fb791d3fad735f48fb6c228f732fdbf481b493d2ca32c24f5@127.0.0.1:3000?discport=0"
```

```
web3.net.peerCount
2
```

现在，每一个节点都与其他节点相连接了