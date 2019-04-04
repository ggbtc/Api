# **API文档**

## **2018-10-11**

## 1.1 / 私有接口签名流程

#### **实例 : 获取用户充值地址** 

地址: https://private.ggbtc.com:55558/user/logined/getNewAddress

- #### 参数列表:

已经从服务端获取的token={key:keykey,value:valuevalue}

| 字段名称 | 描述           |
| -------- | -------------- |
| key      | keykey         |
| value    | valuevalue     |
| nonce    | 12343546       |
| currency | eth            |
| sign     | 字符加密后的串 |

- #### 签名过程

格式化字符串请求URL为:

```
/user/logined/getNewAddress?key=keykey&nonce=12343536&currency=eth
```

上面的字符串后面加上value

```
/user/logined/getNewAddress?key=keykey&nonce=12343536&currency=ethvaluevalue
```

对上面的字符串进行sha512加密,得到的结果是

```
dea67c903dd1af28270f68f9b35c52ef6597d2082694189adbdf9ad805b10e0c5ac73461cfd08e292036ac1a8fb626fe7a5bbd76d5fda33a112a707a96404c87
```

最后完整的请求的URL为:

```
http://apiserver/user/logined/getNewAddress?key=keykey&nonce=12343536&currency=eth&sign=dea67c903dd1af28270f68f9b35c52ef6597d2082694189adbdf9ad805b10e0c5ac73461cfd08e292036ac1a8fb626fe7a5bbd76d5fda33a112a707a96404c87
```

-   所有注明需要签名接口,按照以上流程签名并增加签名所需参数,以下接口详情不在重复提示.

1.2 / 标准返回值
----------------

- ```
    {
     	  "message":	"success", 			//success为成功,若有其他结果,会体现
     	  "result":		{} 				    //只有在成功时才有需要的数据
    }
    
    ```

    


-   所有接口有返回值均代表result内数据,若无数据,则注释标准返回值.

2.0 / 获取用户token（长期）
---------------------------

- #### URL: /user/logined/getLongToken		 **需签名**

- #### 参数列表 

  ​			none

- #### 返回值

  | 字段名称 | 数据类型 |
  | -------- | -------- |
  | Key      | string   |
  | Value    | string   |

2.1 / 获取用户信息
------------------

- #### URL: /user/logined/getInfo			 **需签名**

- #### 参数列表 

  ​			none

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

2.2 / 请求用户充值地址
----------------------

- #### URL: /user/logined/getNewAddress **需签名**

- #### 参数列表	

  | 字段名称 | 描述 | 备注  |
  | -------- | ---- | ----- |
  | currency | 币种 | "eth" |

- #### 返回值

  | 字段名称 | 描述 | 数据类型 |
  | -------- | ---- | -------- |
  | Address  | 地址 | string   |

2.3 / 获取用户余额
------------------

- #### URL: /user/logined/getBalance **需签名**

- #### 参数列表 

  ​						none

- #### 返回值

​                        ~~Balance            用户余额              HashMap~~

```
{

	"eth":				1000000,				 //eth可用余额

	"eth-trade":		10111, 					//eth锁定余额

	"eth-locked":		100, 					//eth锁定余额

	"btc":				11123456, 				//btc可用余额

	"btc-trade":		11123456, 				//btc锁定余额

	"hkdt":				11123456,			   //hkdt可用余额

	"hkdt-trade":		11123456 			   //hkdt锁定余额(交易中的锁定余额)

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
  | -------- | ---------- | -------- |
  | pair     | 交易对     | eth-usdt |
  | limit    | 条数(可选) | 100      |

- #### 返回值

```
[
	T 		时间			 unixTime

	Tp 		类型(买卖) 		sell,buy

	P 		价格 			 float64

	A 		数量

	B		买单ID

	S 		卖单ID
]
```

​	

3.2 / 交易 / 获取24小时数据
---------------------------

- #### URL: /trade/public/getTicker

- #### 参数列表	

  | 字段名称 | 描述   | 备注     |
  | -------- | ------ | -------- |
  | pair     | 交易对 | eth-usdt |

- #### 返回值

  当pair为all时,返回	map\[string\]jsonStr

  当pair为单一币种时,返回jsonStr

```
{

	"Chg": 		"-0.398026", 			//增加或减少的比例

	"High":		 "201.00541", 			//24小时最高价

	"Low": 		"121.00003",	 		//24小时最低价

	"Name": 	"eth-usdt", 			//币种

	"Turn": 	"6331.0673987", 		//24小时成交额(既-后面的币种的成交数量)

	"Vol": 		"31.497", 				//24小时成交量(既-前面的币种的成交数量)

	"Last": 	"121.000003"			//上一次成交价格

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

```
{

	Sell :[ 			//卖单

	{P:100.0,A:30},

		.....

	]

	Buy[], 				//买单

	Last:	123.1 		//最新成交价格

}

```



3.4 / 交易 /获取Kline
---------------------

- #### URL: /trade/public/getKline

- #### 参数列表	

  | 字段名称 | 描述           | 数据类型             |
  | -------- | -------------- | -------------------- |
  | start    | 开始时间(可选) | unixTime             |
  | end      | 结束时间(可选) | unixTime             |
  | basis    | 单位           | hour,minute,day,week |
  | limit    | 条数(可选)     | 100                  |

- #### 返回值

```
[

	T 		时间 		 unixTime

	S 		开始价格	float64

	E 		结束价格 	float64

	H 		最高价格

	L 		最低价格

	A 		数量

]

```



4.1 / 交易 / 用户下单
---------------------

- #### URL: /trade/logined/orderCreate **需签名**

- #### 参数列表	

  | 字段名称  | 描述     | 备注       |
  | --------- | -------- | ---------- |
  | pair      | 交易对   | "eth-usdt" |
  | price     | 下单价格 | 222.3      |
  | amount    | 下单数量 | 10         |
  | tradeType | 交易类型 | "sell"     |

- #### 返回值

  交易ID号(如成功)

4.2 / 交易 / 用户取消订单
-------------------------

- #### URL: /trade/logined/orderCancel 		**需签名**

- #### 参数列表

  | 字段名称 | 描述   | 备注       |
  | -------- | ------ | ---------- |
  | pair     | 交易对 | "eth-usdt" |
  | id       | 订单号 | 612342     |

- #### 返回值

标准返回值

异步处理 , 故这里必返回success

4.3 / 交易 / 用户查询订单
-------------------------

- #### URL: /trade/logined/orderGet 		**需签名**

- #### 参数列表	

  | 字段名称 | 描述   | 备注       |
  | -------- | ------ | ---------- |
  | pair     | 交易对 | "eth-usdt" |
  | id       | 订单号 | 612342     |

- #### 返回值

  ```
  {
      id 					//id号
  
  	user
  
  	pair
  
  	tradetype
  
  	orderprice 			//下单价格
  
  	orderamount 		//下单数量-int-后六位为小数点
  
  	time
  
  	feescale 			//手续费比例
  
  	remainamount		 //未成交剩余数量-int-同上
  
  	feeamount
  
  	fillamount 			//当为卖单: 这里代表用户卖掉eth,获取usdt的数 量
  
   						   //当为买单 : 这里代表用户未使用eth的数量
  	avgprice 			//成交平均价格
  
  	state 				//订单状态		 open,closed,canceled
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

- #### 返回值

```
{
    id 					//id号

	user

	pair

	tradetype

	orderprice 			//下单价格

	orderamount 		//下单数量-int-后六位为小数点

	time

	feescale 			//手续费比例

	remainamount		 //未成交剩余数量-int-同上

	feeamount

	fillamount 			//当为卖单: 这里代表用户卖掉eth,获取usdt的数 量

 						   //当为买单 : 这里代表用户未使用eth的数量
	avgprice 			//成交平均价格

	state 				//订单状态		 open,closed,canceled
}
```

4.5 / 交易 / 用户挂单交易详情
-----------------------------

- #### URL: /trade/logined/transGetByID 需签名

- #### 参数列表

  | 字段名称 | 描述           | 备注       |
  | -------- | -------------- | ---------- |
  | pair     | 交易对         | "eth-usdt" |
  | id       | 用户挂单	id | 60006      |

- #### 返回值

```
[

	{

		T 		时间 		unixTime

		P 		价格 		float64

		A 		数量

	}...

]

```



4.6 / 交易 / 用户交易流水
-------------------------

- #### URL: /trade/logined/getTradeHistory 需签名

- #### 参数列表

  | 字段名称  | 描述                           | 备注             |
  | --------- | ------------------------------ | ---------------- |
  | pair      | 交易对                         | "eth-usdt"       |
  | limit     | 返回交易条数                   | 0-1000 999       |
  | time      | 返回此时间之前的流水(unix微秒) | 1544061142176525 |
  | tradeType | 用户挂单类型                   | sell \| buy      |

- #### 返回值

```
[

	{

		T 		时间 		unixTime

		P 		价格		 float64

		A 		数量

		B 		买家ID

		S 		卖家ID

		Tp 		交易方向,交易实际成交的方向,并不是用户的挂单方向

		ID 		交易流水ID,此ID唯一

	}...

]

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

5.0 公告内容相关API
-------------------

站内公告相关内容使用wordpress开源CMS，当前使用**API地址**：https://api.100ma.com/wp-json

一般使用axios的get请求内容，未整合进中间层，发起请求时**拼接url**：【API地址】+【请求接口?参数(可选)】

发起Get请求时，有**两种传参**方式，如下带背景色文字为传入参数：

1，带参数名称，问号"?"开始,"&"拼接多个参数，/wp/v2/posts?categories=7&page=1

2，不带参数名称，一般只有一个参数，默认参数名称为上一级接口名称，/wp/v2/categories/&lt;id&gt;

举个例子：

get接口获取id为7的分类下的所有文章列表 ，

【url】=【https://api.100ma.com/wp-json】+【/wp/v2/posts?categories=7】

无参数时，默认获取所有，如下例子获取所有文章列表，

【url】=【https://api.100ma.com/wp-json】+【/wp/v2/posts】

5.1 获取分类列表
----------------

- #### 获取全部分类列表

接口：/wp/v2/categories

返回：默认10个长度的分类数组

![](C:\Users\liuqi\Desktop\markdown\img1.png)

- #### 获取特定id号的分类

接口：/wp/v2/categories/&lt;id&gt;

返回：对应id的分类对象

![](C:\Users\liuqi\Desktop\markdown\img2.png)

5.2 获取文章列表
----------------

- #### 获取全部文章列表

接口：/wp/v2/posts

返回：默认10个长度的文章数组

![](C:\Users\liuqi\Desktop\markdown\img3.png)

- #### 获取特定分类下的文章列表

接口：/wp/v2/posts?categories=&lt;id&gt;

返回：默认10个长度的文章数组

![](C:\Users\liuqi\Desktop\markdown\img4.png)



5.3 获取文章内容
----------------

接口：/wp/v2/posts/<id>

返回：对应id的文章

![](C:\Users\liuqi\Desktop\markdown\img5.png)
