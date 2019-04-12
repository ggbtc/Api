# **API文档**

## **2018-10-11**

## 1.1 / 私有接口签名流程

### 加密方式
用户登陆从服务端获取到的的token={key:apikey,value:secret}

 - apikey=K28f43a180373d7e
 - secret=Ve70aee66b8a3e2bf

在线加密工具: http://tool.oschina.net/encrypt?type=2

#### **实例 : 获取用户充值地址** 

地址: https://private.ggbtc.com:55558/user/logined/getNewAddress

- #### 参数列表:

| 字段名称 | 描述           |
| -------- | -------------- |
| key      | apikey         |
| nonce    | 12343546       |
| currency | eth            |
| sign     | 字符加密后的串  |

- #### 签名过程

格式化字符串请求URL为:

```
/user/logined/getNewAddress?key=K28f43a180373d7e&nonce=12343536&currency=eth
```

上面的字符串后面加上secret

```
/user/logined/getNewAddress?key=K28f43a180373d7e&nonce=12343536&currency=ethVe70aee66b8a3e2bf
```

对上面的字符串进行sha512加密,得到的sign结果是

```
45d05a916083d4949ac090308be193d79c135090de44f6e950601e6658f09e74a624d0044e4628127eaaced232b0840315b248a729846803055d6122b7e00d16
```

最后加上API服务器地址，完整的请求的URL为:

```
https://private.ggbtc.com:55558/user/logined/getNewAddress?key=K28f43a180373d7e&nonce=12343536&currency=eth&sign=45d05a916083d4949ac090308be193d79c135090de44f6e950601e6658f09e74a624d0044e4628127eaaced232b0840315b248a729846803055d6122b7e00d16
```

-   所有注明需要签名接口,按照以上流程签名并增加签名所需参数,以下接口详情不在重复提示.

1.2 / 标准返回值
----------------
```
    {
     	  "message":	"success", 			//success为成功,若有其他结果,会体现
     	  "result":		{} 				    //只有在成功时才有需要的数据
    }
```

-   所有接口有返回值均代表result内数据，若无数据，则注释标准返回值.

2.0 / 获取用户token（长期）
---------------------------

- #### URL: /user/logined/getLongToken		 **需签名**

- #### 参数列表 

| 字段名称 | 描述           |
| -------- | -------------- |
| key      | apikey         |
| nonce    | 12343546       |
| sign     | 签名           |

- #### 返回值

  | 字段名称 | 数据类型 |
  | -------- | -------- |
  | message  | string   |
  | result   | string   |

```
    {
     	  "message":	"success", 			
     	  "result":		
     	        {
     	            "User":"18812341234",
     	            "Key":"L-376e79ac9a079",
     	            "Value":"V9ac62499374c0143",
     	            "PayPwdCheck":0,
     	            "LongToken":true
     	        } 
    }
```
2.1 / 获取用户信息
------------------

- #### URL: /user/logined/getInfo			 **需签名**

- #### 参数列表 

| 字段名称 | 描述           |
| -------- | -------------- |
| key      | apikey         |
| nonce    | 12343546       |
| sign     | 签名           |

- #### 返回值

  | 字段名称           | 描述               | 数据类型    |
  | ------------------ | ------------------ | ----------- |
  | Id                 | 用户id1324         |             |
  | User               | 用户名             | string      |
  | Referrer           | 推荐人             | string      |
  | Time               | 注册时间           | string      |
  | Email              | 信箱               | string      |
  | Address            | 充值地址           | HashMap     |
  | ~~Balance~~        | ~~余额--预备移除~~ | ~~HashMap~~ |
  | Group              | 用户组             | string      |
  | IsLocked           | 用户是否锁定       | bool        |
  | IsBalanceLocked    | 用户余额是否锁定   | bool        |
  | WithdrawLockBefore | 用户提现锁定时间   | unixTime    |

```
    {
     	  "message":	"success", 			
     	  "result":		
                {
                    "User":"18812341234",
                    "NickName":"",
                    "RealName":"",
                    "ID":"06000552",
                    "UserVerify":"",
                    "Referrer":"",
                    "Pwd":"",
                    "RegIP":"54.73.187.84",
                    "Time":"2019-04-11 11:03:33.788962153 +0800 CST m=+241966.766266219",
                    "Email":"",
                    "Phone":"18812341234",
                    "IsLocked":false,
                    "IsBalanceLocked":false,
                    "IsWithdrawLocked":0,
                    "Viplv":0,
                    "Group":"",
                    "Balance":{},
                    "Address":
                        {
                            "eos":"ggbtcdeposit",
                            "eosmemo":"06000552",
                            "xrp":"rhx8L1msK3jmJWKhngVQETBuYZXMr86666",
                            "xrpmemo":"06000552"
                        },
                    "WrAddress":[],
                    "Favourite":"",
                    "GoogleKey":"",
                    "GoogleVerify":false,
                    "PayPwd":"false",
                    "FeeScale":{},
                    "Oauth":{}
                } 
    }
```

2.2 / 请求用户充值地址
----------------------

- #### URL: /user/logined/getNewAddress **需签名**

- #### 参数列表	

  | 字段名称 | 描述 | 备注  |
  | -------- | ---- | ----- |
  | currency | 币种 | "eth" |
  | key      | apikey|"K28f43a180373d7e"|
  | nonce    | 随机数 |12343546    |
  | sign     | 签名   |签名     |

- #### 返回值

  | 字段名称 | 描述 | 数据类型 |
  | -------- | ---- | -------- |
  | Address  | 地址 | string   |
  | Currency | 币种 | string   |
  | User     | 用户 | string   |

```
    {
     	  "message":	"success", 			
     	  "result":		{
                            Address: "0x245405d2e4989e530fec39063854cffc3e9a1d14",
                            Currency: "eos",
                            User: "18812341234"
                        } 
    }
```
2.3 / 获取用户余额
------------------

- #### URL: /user/logined/getBalance **需签名**

- #### 参数列表 

  | 字段名称 | 描述    |
  | -------- | ----   |
  | key      | apikey |
  | nonce    | 随机数  |
  | sign     | 签名    |

- #### 返回值

​                        ~~Balance            用户余额              HashMap~~
```
    {
     	  "message":	"success", 			
     	  "result":		{
                "eth":				1000000,				//eth可用余额
                "eth-trade":		10111, 					//eth锁定余额
                "eth-locked":		100, 					//eth锁定余额
                "btc":				11123456, 				//btc可用余额
                "btc-trade":		11123456, 				//btc锁定余额
                "hkdt":				11123456,			    //hkdt可用余额
                "hkdt-trade":		11123456 			    //hkdt锁定余额(交易中的锁定余额)
            }
    }
```

~~//注:以上每一个字段都可能不存在,不存在的字段按照0处理,以上数据包含6位小数点~~

//注1:-trade和-locked均代表锁定余额

//注2: 当余额为空时,字段可能不存在


3.1 / 交易 / 获取交易历史记录
-----------------------------

- #### URL: /trade/public/getTradeHistory

- #### 参数列表	


 | 字段名称 | 描述       | 备注     |
 | -------- | ----------| -------- |
 | pair     | 交易对     | eth-usdt |
 | limit    | 条数(可选) | 100      |

- #### 返回值

  | 字段名称 | 描述 | 数据类型  |
  | -------- | ---- | --------  |
  | P        | 价格 | float64    |
  | A        | 数量 | float64    |
  | T        | 时间 | string     |
  | Tp       | 买卖类型 | string  |
  | B        | 买单ID | string  |
  | S        | 卖单ID | string  |
  
```
    {
        "message":  "success",
        "result":   [
            {
                "P":144.86615943060704,             
                "A":36.5,                           
                "T":"1554879426166901",            
                "Tp":"buy"                          
            },
            {
                "P":144.86615943060704,
                "A":16.5,
                "T":"1554879424559248",
                "Tp":"buy"
            },
            {
                P: 144.77601684515983,
                A: 31,
                T: "1554879376508447",
                Tp: "sell"
            },            
        ]
    }
```

3.2 / 交易 / 获取24小时数据
---------------------------

- #### URL: /trade/public/getTicker

- #### 参数列表	

  | 字段名称 | 描述   | 备注     |
  | -------- | ------ | -------- |
  | pair     | 交易对 | eth-usdt |

- #### 返回值

  | 字段名称 | 描述 | 数据类型  |
  | -------- | ------------------ | --------  |
  | chg      | 增加或减少的比例 | string    |
  | high     | 24小时最高价     | string    |
  | low      | 24小时最低价     | string  |
  | name     | 币种             | string  |
  | id       | 排序             | string     |
  | last     | 上一次成交价格    | string  |
  | turn     | 24小时成交额(既-后面的币种的成交数量) | string  |
  | vol      | 24小时成交量(既-前面的币种的成交数量) | string  |

  当pair为单一币种时,返回jsonStr

```
    {
        "message":  "success",
        "result":   
            {
                "chg":"0.0027244870418008382",
                "high":"170.96678350000002",
                "id":"1",
                "last":"160.799754",
                "low":"159.507454",
                "name":"eth-usdt",
                "turn":"835575.6609929006",
                "vol":"5040.0689999999995"
            }
    }
```
  当pair为all时,返回	map\[string\]jsonStr

```
    {
        "message":  "success",
        "result":   {
            "ankr-btc":
                {
                    "chg":"0","high":"0.00000309","last":"0.00000203","low":"0.00000189","name":"ankr-btc","turn":"1.735185602580714","vol":"712810.62287742"
                },
            "apk-usdt":
                {"chg":"0.00003413910824501265","closed":"true","high":"0.055","id":"131","last":"0.02899999","low":"0.0273","name":"apk-usdt","turn":"19590.240049685468","vol":"667533.286858"
                },
            }
    }
```



3.3 / 交易 / 获取订单薄
-----------------------

- #### URL: /trade/public/getOrderbook

- #### 参数列表

  | 字段名称 | 描述       | 备注     |
  | -------- | ---------- | -------- |
  | pair     | 交易对     | eth-usdt |
  | limit    | 条数(可选) | 100      |

- #### 返回值

  | 字段名称 | 描述 | 数据类型  |
  | -------- | --------   | --------  |
  | Time     | 时间戳      | int64    |
  | Last     | 最新成交价格 | float64  |
  | Sell     | 卖单        | array    |
  | Buy      | 买单        | array    |

```

    {
        "message":  "success",
        "result":   {
            "Sell":     //卖单
                [
                    {
                        "P":163.668664,     //价格float64
                        "A":0.025           //数量float64
                    },
                    {
                        "P":163.68891,
                        "A":0.06
                        },
                    {
                        "P":163.759771,
                        "A":0.0525
                    },
                ],
            "Buy":      //买单
                [
                    {
                        "P":159.365395,
                        "A":0.0475
                    },
                    {
                        "P":159.345641,
                        "A":0.03
                    },
                    {
                        "P":159.325887,
                        "A":0.1
                    },
                ],
            "Time":1555037133,
            "Last":172.62467339844997
            }
    }
```



3.4 / 交易 /获取Kline
---------------------

- #### URL: /trade/public/getKline

- #### 参数列表	

  | 字段名称 | 描述            | 数据类型             |
  | -------- | -------------- | -------------------- |
  | pair     | 交易对（必填）  |eth-usdt            |
  | basis    | 单位（必填）    | hour,minute,day,week |
  | start    | 开始时间(可选)  | unixTime             |
  | end      | 结束时间(可选)  | unixTime             |
  | limit    | 条数(可选)      | 100                  |

- #### 返回值

  | 字段名称 | 描述       | 数据类型  |
  | ------- | --------   | --------  |
  | T       | 时间戳      | int64    |
  | S       | 开始价格    | float64  |
  | E       | 结束价格    | float64  |
  | H       | 最高价格    | float64  |
  | L       | 最低价格    | float64  |
  | A       | 数量        | float64  |

```
    {
        "message":      "success",
        "result":       
            [
                {
                    "T":1555038000,
                    "S":160.3628475,
                    "E":161.5270295,
                    "H":161.93172199999998,
                    "L":160.185695,
                    "A":221.724
                },
                {
                    "T":1555034400,
                    "S":163.6573985,
                    "E":160.3628475,
                    "H":163.6573985,
                    "L":159.507454,
                    "A":283.04400000000004
                }
            ]
    }
```



4.1 / 交易 / 用户下单
---------------------

- #### URL: /trade/logined/orderCreate **需签名**

- #### 参数列表	

  | 字段名称  | 描述      | 备注               |
  | --------- | -------- | ---------------    |
  | pair      | 交易对   | "eth-usdt"         |
  | price     | 下单价格 | 222.3              |
  | amount    | 下单数量 | 10                 |
  | tradeType | 交易类型 | "sell"             |
  | key       | apikey   | "K28f43a180373d7e" |
  | nonce     | 随机数    | 123456             | 
  | sign      | 签名     | 签名串              | 

- #### 返回值

  交易ID号(如成功)

```
    {
        "message":      "success",
        "result":       6027829     //订单ID
    }
```

4.2 / 交易 / 用户取消订单
-------------------------

- #### URL: /trade/logined/orderCancel 		**需签名**

- #### 参数列表

  | 字段名称 | 描述   | 备注       |
  | -------- | ------ | ---------- |
  | pair     | 交易对 | "eth-usdt" |
  | id       | 订单号 | 612342     |
  | key       | apikey   | "K28f43a180373d7e" |
  | nonce     | 随机数    | 123456             | 
  | sign      | 签名     | 签名串              | 

- #### 返回值

标准返回值

异步处理 , 故这里必返回success

```
    {
        "message":      "success",      
        "result":       
    }
```
```
    {
        "message":      "cancelFailed",    
        "result":       
    }
```

4.3 / 交易 / 用户查询订单
-------------------------

- #### URL: /trade/logined/orderGet 		**需签名**

- #### 参数列表	

  | 字段名称 | 描述   | 备注       |
  | -------- | ------ | ---------- |
  | pair     | 交易对 | "eth-usdt" |
  | id       | 订单号 | 6027829     |
  | key       | apikey   | "K28f43a180373d7e" |
  | nonce     | 随机数    | 123456             | 
  | sign      | 签名     | 签名串              | 

- #### 返回值

  | 字段名称     | 描述       | 数据类型  |
  | -------     | --------   | --------  |
  | user        | 用户        | string    |
  | id          | 订单号      | int64  |
  | pair        | 交易对      | string  |
  | tradetype   | 买卖类型    | string  |
  | orderprice  | 下单价格    | float64  |
  | orderamount | 下单数量-int-后六位为小数点        | int  |
  | time        | 下单时间        | int64  |
  | ismarket    | 是否市场    | bool  |
  | feescale    | 手续费比例   | float64  |
  | remainamount| 未成交剩余数量-int-同上        | int  |
  | feeamount   | 手续费数量        | int  |
  | fillamount  | 当为卖单: 这里代表用户卖掉eth,获取usdt的数量/当为买单 : 这里代表用户未使用eth的数量        | float64  |
  | avgprice    | 成交平均价格        | float64  |
  | state       | 订单状态(open\closed\canceled)        | string  |
  | lastupdate  | 最后一次更新时间        | string  |

```  
    {
        "message":      "success",
        "result":       
            {
                "user":"18812341234",
                "id":6027829,
                "pair":"eth-usdt",
                "tradetype":"sell",
                "orderprice":200,
                "orderamount":1,
                "time":1555038765693,
                "ismarket":false,
                "feescale":1,
                "remainamount":1,
                "feeamount":0,
                "fillamount":0,
                "avgprice":0,
                "state":"canceled",
                "lastupdate":"2019-04-12 11:20:42"
            }
    }
```

4.4 / 交易 / 用户查询当前挂单
-----------------------------

- #### URL: /trade/logined/orderGetUser		 **需签名**

- #### 参数列表	

  | 字段名称  | 描述                   | 备注                 |
  | --------- | ---------------------- | -------------------- |
  | pair      | 交易对                 | "eth-usdt"           |
  | limit     | 查询数量(可选)         | 100                  |
  | time      | 查询截止时间(可选)     | unixTime毫秒         |
  | tradeType | 查询买单还是卖单(可选) | sell,buy             |
  | state     | 查询订单状态(可选)     | open,closed,canceled |
  | key       | apikey                | "K28f43a180373d7e" |
  | nonce     | 随机数                | 123456             | 
  | sign      | 签名                  | 签名串              | 

- #### 返回值

  | 字段名称     | 描述       | 数据类型  |
  | -------     | --------   | --------  |
  | user        | 用户        | string    |
  | id          | 订单号      | int64  |
  | pair        | 交易对      | string  |
  | tradetype   | 买卖类型    | string  |
  | orderprice  | 下单价格    | float64  |
  | orderamount | 下单数量-int-后六位为小数点        | int  |
  | time        | 下单时间        | int64  |
  | ismarket    | 是否市场    | bool  |
  | feescale    | 手续费比例   | float64  |
  | remainamount| 未成交剩余数量-int-同上        | int  |
  | feeamount   | 手续费数量        | int  |
  | fillamount  | 当为卖单: 这里代表用户卖掉eth,获取usdt的数量/当为买单 : 这里代表用户未使用eth的数量        | float64  |
  | avgprice    | 成交平均价格        | float64  |
  | state       | 订单状态(open\closed\canceled)        | string  |
  | lastupdate  | 最后一次更新时间        | string  |

```
    {
        "message":      "success",
        "result":
            [
                {
                    "user":"18812341234",
                    "id":6027832,
                    "pair":"eth-usdt",
                    "tradetype":"buy",
                    "orderprice":185.95,
                    "orderamount":0.954,
                    "time":1555042598190,
                    "ismarket":false,
                    "feescale":1,
                    "remainamount":0.954,
                    "feeamount":0,
                    "fillamount":177.7510926,
                    "avgprice":0,
                    "state":"open",
                    "lastupdate":"2019-04-12 12:16:38"
                },
                {
                    "user":"18812341234",
                    "id":6027831,
                    "pair":"eth-usdt",
                    "tradetype":"buy",
                    "orderprice":180,
                    "orderamount":3.56,
                    "time":1555042564628,
                    "ismarket":false,
                    "feescale":1,
                    "remainamount":3.56,
                    "feeamount":0,
                    "fillamount":642.0816,
                    "avgprice":0,
                    "state":"open",
                    "lastupdate":"2019-04-12 12:16:04"
                }
            ]
        }
```

4.5 / 交易 / 用户挂单交易详情
-----------------------------

- #### URL: /trade/logined/transGetByID **需签名**

- #### 参数列表

  | 字段名称 | 描述           | 备注       |
  | -------- | -------------- | ---------- |
  | pair     | 交易对         | "eth-usdt" |
  | id       | 用户挂单	id | 60006      |
  | key       | apikey                | "K28f43a180373d7e" |
  | nonce     | 随机数                | 123456             | 
  | sign      | 签名                  | 签名串              | 

- #### 返回值

  | 字段名称  | 描述       | 数据类型  |
  | -------  | --------   | --------  |
  | P        | 价格      | float64    |
  | A        | 数量      | float64  |
  | T        | 时间      | int64  |

```
    {
        "message":  "success",
        "result":   
            [
                {
                    "P":185.95,
                    "A":0.095,
                    "T":"1555046678004735"
                },
                {
                    "P":185.95,
                    "A":0.07,
                    "T":"1555046678006127"
                },
                {
                    "P":185.95,
                    "A":0.095,
                    "T":"1555046678007048"
                }
            ]
    }
```



4.6 / 交易 / 用户交易流水
-------------------------

- #### URL: /trade/logined/getTradeHistory **需签名**

- #### 参数列表

  | 字段名称  | 描述                           | 备注             |
  | --------- | ------------------------------ | ---------------- |
  | pair      | 交易对(必填)                         | "eth-usdt"       |
  | limit     | 返回交易条数（选填）                   | 0-1000 999       |
  | time      | 返回此时间之前的流水(unix微秒)（选填） | 1544061142176525 |
  | tradeType | 用户挂单类型（选填）                   | sell \| buy      |

- #### 返回值

  | 字段名称  | 描述       | 数据类型  |
  | -------  | --------   | --------  |
  | P        | 价格      | float64    |
  | A        | 数量      | float64  |
  | T        | 时间      | string  |
  | B        | 买家ID      | int64  |
  | S        | 卖家ID      | int64  |
  | Tp        | 交易方向,交易实际成交的方向,并不是用户的挂单方向      | string  |
  | ID        | 交易流水ID,此ID唯一      | string  |

```
    {
        "message":  "success",
        "result":
            [
                {
                    "P":200,
                    "A":0.36,
                    "T":"1555047989718225",
                    "B":6027858,
                    "S":6027830,
                    "Tp":"buy",
                    "ID":"5cb0263631caa899297e85ea"
                },
                {
                    "P":200,
                    "A":1,
                    "T":"1555047989717567",
                    "B":6027858,
                    "S":6027828,
                    "Tp":"buy",
                    "ID":"5cb0263631caa899297e85e9"
                },
                {
                    "P":185.95,
                    "A":0.07,
                    "T":"1555046678510499",
                    "B":6027832,
                    "S":6027856,
                    "Tp":"sell",
                    "ID":"5cb0211731caa899297e6680"
                }
            ]
    }
```



-   注:

1.
字段里的Tp和提交的tradeType无关,仅代表实际交易成功的交易方向,并不代表是用户挂单的方向,例:
自己挂了一单买单,之后自己又挂了一单卖单成交了自己的买单,那么成交方向为sell

2\. 当提交tradeType参数后,仅返回用户做为(买/卖单)的流水, 举例 :

用户挂了一单价格为200元,数量为10的卖单,之后用户又挂了一单相同价格和数量的买单,成交后

用户查询成交记录

1\) 不提交tradeType字段,或者字段为空,流水的返回值为

```
T=15xxxxxxxxxx P=200 A=10 S=123 B=124 Tp=buy ID=873213214cfds3
```

2\) 提交tradeType字段为sell,返回值为,

```
T=15xxxxxxxxxx P=200 A=10 S=123 B=124 Tp=buy ID=873213214cfds3
```

此时返回的流水,均为用户挂单为"卖"的流水

3\) 提交tradeType字段为buy,返回值为,

```
T=15xxxxxxxxxx P=200 A=10 S=123 B=124 Tp=buy ID=873213214cfds3
```

此时返回的流水,均为用户挂单为"买"的流水

如上: 在自交易时,返回值完全相同 .
通过提交不同的tradeType可以获得重复的流水

