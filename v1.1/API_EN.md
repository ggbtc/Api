# **API  Document**

## **2019-4-12**

## 1.1 /  Private interface signature process 

### Encryption Methon
The token got from server when user logined, token={key:apikey, value:secretkey},example:

 - apikey=K28f43a180373d7e
 - secretkey=Ve70aee66b8a3e2bf

1.URL Request String: concatenating the request method URL including any query-string parameters, your public API Key, and a user-generated nonce (we recommend a UNIX-style timestamp in epoch millisecond format).

2.Add the secretkey to the end of URL Request String.

3.Secret as well as the Sha512 yielded use the string, get the sign value.

4.Add the API Server address before the URL Request String, Add the sign parameter to the end.

Online encryption tool: http://tool.oschina.net/encrypt?type=2

#### example: Get user recharge address 

url: https://private.ggbtc.com:55558/user/logined/getNewAddress

- #### Parameters  list:

| FieldName | Description      |
| --------- | ---------------- |
| key       | apikey           |
| nonce     | 12343546         |
| currency  | eth              |
| sign      | encrypted string |

- #### Signature process 

URL of formatting string's request:

```
/user/logined/getNewAddress?key=K28f43a180373d7e&nonce=12343536&currency=eth
```

The string above is followed by 'secretkey' 

```
/user/logined/getNewAddress?key=K28f43a180373d7e&nonce=12343536&currency=ethVe70aee66b8a3e2bf
```

Sha512 encryption of the above string, the result is :

```
45d05a916083d4949ac090308be193d79c135090de44f6e950601e6658f09e74a624d0044e4628127eaaced232b0840315b248a729846803055d6122b7e00d16
```

Final add the API Server Address, the full request's URL: 

```
http://apiserver/user/logined/getNewAddress?key=K28f43a180373d7e&nonce=12343536&currency=eth&sign=45d05a916083d4949ac090308be193d79c135090de44f6e950601e6658f09e74a624d0044e4628127eaaced232b0840315b248a729846803055d6122b7e00d16
```

-   All interfaces marked for signature ,must follow the above procedure to sign and add the required parameters ,the following interface details will not repeated .

1.2 / Standard return value 
----------------

```
    {
     	  "message":	"success", 		//'success' represets success, if there are other	
     	  "result":	{} 					//Data is needed only for success
    }
```



-   All interfaces have return values that represent data in'result'. If  no , comment the standard return value. 

2.0 /Get user token（long-term）
---------------------------

- #### URL: /user/logined/getLongToken		Need to sign 

- #### Parameters  list: 

  | FieldName | Description      |
  | --------- | ---------------- |
  | key       | apikey           |
  | nonce     | timestamp        |
  | sign      | encrypted string |

- #### Return value:

  | FieldName   | Description   | DataType |
  | ----------- | ------------- | -------- |
  | User        | UserName      | string   |
  | Key         | apikey        | string   |
  | Value       | secretkey     | string   |
  | PayPwdCheck | pay pwd check | int      |
  | LongToken   | long token    | bool     |

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

2.1 / Get user information
------------------

- #### URL: /user/logined/getInfo			 Need to sign 

- #### Parameters  list: 

  | FieldName | Description      |
  | --------- | ---------------- |
  | key       | apikey           |
  | nonce     | timestamp        |
  | sign      | encrypted string |

- #### Return value:

  | FieldName          | Description                     | DataType    |
  | ------------------ | ------------------------------- | ----------- |
  | Id                 | userId1324                      |             |
  | User               | userName                        | string      |
  | Referrer           | referrer                        | string      |
  | Time               | registration time               | string      |
  | Email              | e-mail                          | string      |
  | Address            | recharge address                | HashMap     |
  | ~~Balance~~        | ~~balance (Prepare to remove)~~ | ~~HashMap~~ |
  | Group              | user group                      | string      |
  | IsLocked           | If the user is locked out       | bool        |
  | IsBalanceLocked    | If the balance is locked out    | bool        |
  | WithdrawLockBefore | User withdrawal lock time       | unixTime    |

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

2.2 / Request user recharge address
----------------------

- #### URL: /user/logined/getNewAddress 			 Need to sign 

- #### Parameters list	

| FieldName | Description      | Remark             |
| --------- | ---------------- | ------------------ |
| currency  | currency         | "eth"              |
| key       | apikey           | "K28f43a180373d7e" |
| nonce     | timestamp        | "1542020339856"    |
| sign      | encrypted string | encrypted string   |

- #### Return value

  | FieldName | Description | DataType |
  | --------- | ----------- | -------- |
  | Address   | Address     | string   |
  | Currency  | Currency    | string   |
  | User      | User        | string   |

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

2.3 / Get user balance
------------------

- #### 	URL: /user/logined/getBalance 			Need to sign 

- #### Parameters list 

| FieldName | Description      | Remark             |
| --------- | ---------------- | ------------------ |
| key       | apikey           | "K28f43a180373d7e" |
| nonce     | timestamp        | "1542020339856"    |
| sign      | encrypted string | encrypted string   |

- #### Return value

​                        ~~Balance           user balance              HashMap~~
```
    {
     	  "message":	"success", 			
     	  "result":		{
                "eth":				1000000,				//eth available balance
                "eth-trade":		10111, 					//eth locked balance
                "eth-locked":		100, 					//eth locked balance
                "btc":				11123456, 				//btc available balance
                "btc-trade":		11123456, 				//btc locked balance
                "hkdt":				11123456,			    //hkdt available balance
                "hkdt-trade":		11123456 			    //hkdt locked balance(locked balance in trade)
            }
    }
```


//~~record:Each of the above fields may not exist, the nonexistent fields are treated as 0, and the above data contains 6 decimal places~~ 

//record1:'-trade' and '-locked'  both represent locked balances 

//record2: When the balance is empty, the field may not exist 



3.1 / Trade / get trade history
-----------------------------

- #### URL: /trade/public/getTradeHistory

- #### Parameters list	

| FieldName | Description                    | Remark     |
| --------- | ------------------------------ | ---------- |
| pair      | a pair of trading              | "eth-usdt" |
| limit     | number of branches (optional ) | 100        |

- #### Return value

  | FieldName | Description   | DataType |
  | --------- | ------------- | -------- |
  | P         | price         | float64  |
  | A         | amount        | float64  |
  | T         | time          | string   |
  | Tp        | trade type    | string   |
  | B         | buy order ID  | string   |
  | S         | sell order ID | string   |

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

3.2 / trade / get 24 hours of data 
---------------------------

- #### URL: /trade/public/getTicker

- #### Parameters list	

| FieldName | Description       | Remark   |
| --------- | ----------------- | -------- |
| pair      | a pair of trading | eth-usdt |

- #### Return value
  | FieldName | Description                                                              | DataType |
  | --------- | ------------------------------------------------------------------------ | -------- |
  | chg       | increase or decrease ratio                                               | string   |
  | high      | 24 hour maximum price                                                    | string   |
  | low       | 24 hour minimum price                                                    | string   |
  | name      | a pair of trading                                                        | string   |
  | id        | sort                                                                     | string   |
  | last      | last transaction price                                                   | string   |
  | turn      | 24 hours turnover (the number of transactions in the following currency) | string   |
  | vol       | 24 hours volume (the number of transactions in the previous currency)    | string   |

  Returns  jsonStr when the pair  for a single currency 
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

  Returns	map\[string\]jsonStr when the pair is all

```
    {
        "message":  "success",
        "result":   {
            "ankr-btc":
                {
                    "chg":"0",
                    "high":"0.00000309",
                    "last":"0.00000203",
                    "low":"0.00000189",
                    "name":"ankr-btc",
                    "turn":"1.735185602580714",
                    "vol":"712810.62287742"
                },
            "apk-usdt":
                {
                    "chg":"0.00003413910824501265",
                    "closed":"true",
                    "high":"0.055",
                    "id":"131",
                    "last":"0.02899999",
                    "low":"0.0273",
                    "name":"apk-usdt",
                    "turn":"19590.240049685468",
                    "vol":"667533.286858"
                },
            }
    }
```

3.3 / trade / getOrderBook
-----------------------

- #### URL: /trade/public/getOrderbook

- #### Parameters list

  | FieldName | Description                    | Remark   |
  | --------- | ------------------------------ | -------- |
  | pair      | a pair of trading              | eth-usdt |
  | limit     | number of branches (optional ) | 100      |

- #### Return value

  | FieldName | Description              | DateType |
  | --------- | ------------------------ | -------- |
  | Time      | timestamp                | int64    |
  | Last      | latest transaction price | float64  |
  | Sell      | sell orders              | array    |
  | Buy       | buy orders               | array    |

```

    {
        "message":  "success",
        "result":   {
            "Sell":     //sell orders
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
            "Buy":      //buy orders
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

3.4 / trade /getLine
---------------------

- #### URL: /trade/public/getKline

- #### Parameters list	

| FieldName | Description                    | Remark               |
| --------- | ------------------------------ | -------------------- |
| pair      | a pair of trading(required)    | eth-usdt             |
| basis     | unit(required)                 | hour,minute,day,week |
| start     | StartTime(optional)            | unixTime             |
| end       | EndTime(optional)              | unixTime             |
| limit     | number of branches (optional ) | 100                  |

- #### Return value
  
  | FieldName | Description      | DataType |
  | --------- | ---------------- | -------- |
  | T         | timestamp        | int64    |
  | S         | start price      | float64  |
  | E         | end price        | float64  |
  | H         | highest price    | float64  |
  | L         | the lowest price | float64  |
  | A         | amount           | float64  |
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


4.1 /trade / orderCreate
---------------------

- #### URL: /trade/logined/orderCreate                             Need to sign

- #### Parameters list	

| FieldName | Description       | Remark             |
| --------- | ----------------- | ------------------ |
| pair      | a pair of trading | "eth-usdt"         |
| price     | order price       | 222.3              |
| amount    | order quantity    | 10                 |
| tradeType | trade type        | "sell"             |
| key       | apikey            | "K28f43a180373d7e" |
| nonce     | timestamp         | 1542020339856      |
| sign      | encrypted string  | encrypted string   |

- #### Return value

  tradeID(success)

```
    {
        "message":      "success",
        "result":       6027829     //trade ID
    }
```

4.2 / trade / userCancelOrder
-------------------------

- #### URL: /trade/logined/orderCancel 		               Need to sign

- #### Parameters list

  | FieldName | Description       | Remark             |
  | --------- | ----------------- | ------------------ |
  | pair      | a pair of trading | "eth-usdt"         |
  | id        | orderID           | 612342             |
  | key       | apikey            | "K28f43a180373d7e" |
  | nonce     | timestamp         | 1542020339856      |
  | sign      | encrypted string  | encrypted string   |

- #### Return value

Standard return value 

asynchronous  processing  ,so be sure to return  success
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
4.3 / trade / user query order
-------------------------

- #### URL: /trade/logined/orderGet 		          Need to sign

- #### Parameters list	

| FieldName | Description       | Remark             |
| --------- | ----------------- | ------------------ |
| pair      | a pair of trading | "eth-usdt"         |
| id        | orderID           | 612342             |
| key       | apikey            | "K28f43a180373d7e" |
| nonce     | timestamp         | 1542020339856      |
| sign      | encrypted string  | encrypted string   |

- #### Return value
  | FieldName                                                                                                                                                               | Description                                                              | DataType |
  | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------ | -------- |
  | user                                                                                                                                                                    | user                                                                     | string   |
  | id                                                                                                                                                                      | order id                                                                 | int64    |
  | pair                                                                                                                                                                    | a pair of trading                                                        | string   |
  | tradetype                                                                                                                                                               | trade type                                                               | string   |
  | orderprice                                                                                                                                                              | order price                                                              | float64  |
  | orderamount                                                                                                                                                             | order quantity-int-The last six digits are the decimal points            | float64  |
  | time                                                                                                                                                                    | order time                                                               | int64    |
  | ismarket                                                                                                                                                                | is market                                                                | bool     |
  | feescale                                                                                                                                                                | fee scale                                                                | float64  |
  | remainamount                                                                                                                                                            | remaining unsold quantity-int-The last six digits are the decimal points | float64  |
  | feeamount                             | this is the user selling the eth and getting the amount of the usdt      | float64  |
  | fillamount                            | When it is a sell order: Here is the sale of eth on behalf of the user, the number of usdt is obtained; when it is paid: here represents the number of unused eth users | float64                                                                  |
  | avgprice                                                                                                                                                                | average transaction price                                                | float64  |
  | state                                                                                                                                                                   | order state(open\closed\canceled)                                        | string   |
  | lastupdate                                                                                                                                                              | last update time                                                         | string   |

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

4.4 / trade / user query the current pending order 
-----------------------------

- #### URL: /trade/logined/orderGetUser		 Need to sign

- #### Parameters list	

| FieldName | Description                                  | Remark                |
| --------- | -------------------------------------------- | --------------------- |
| pair      | a pair of trading                            | "eth-usdt"            |
| limit     | query quantity(optional)                     | 100                   |
| time      | query deadline(optioanal)                    | unixTime  millisecond |
| tradeType | Check the buy order or sell order (optional) | sell,buy              |
| state     | query order state(optional)                  | open,closed,canceled  |
| key       | apikey                                       | "K28f43a180373d7e"    |
| nonce     | timestamp                                    | 1542020339856         |
| sign      | encrypted string                             | encrypted string      |

- #### Return value
  | FieldName                                                                                                                                                               | Description                                                              | DataType |
  | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------ | -------- |
  | user                                                                                                                                                                    | user                                                                     | string   |
  | id                                                                                                                                                                      | order id                                                                 | int64    |
  | pair                                                                                                                                                                    | a pair of trading                                                        | string   |
  | tradetype                                                                                                                                                               | trade type                                                               | string   |
  | orderprice                                                                                                                                                              | order price                                                              | float64  |
  | orderamount                                                                                                                                                             | order quantity-int-The last six digits are the decimal points            | float64  |
  | time                                                                                                                                                                    | order time                                                               | int64    |
  | ismarket                                                                                                                                                                | is market                                                                | bool     |
  | feescale                                                                                                                                                                | fee scale                                                                | float64  |
  | remainamount                                                                                                                                                            | remaining unsold quantity-int-The last six digits are the decimal points | float64  |
  | feeamount                                                                                                                                                               | this is the user selling the eth and getting the amount of the usdt      | float64  |
  | fillamount                            | When it is a sell order: Here is the sale of eth on behalf of the user, the number of usdt is obtained; when it is paid: here represents the number of unused eth users | float64                                                                  |
  | avgprice                                                                                                                                                                | average transaction price                                                | float64  |
  | state                                                                                                                                                                   | order state(open\closed\canceled)                                        | string   |
  | lastupdate                                                                                                                                                              | last update time                                                         | string   |
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

4.5 / trade / user pending transaction details 
-----------------------------

- #### URL: /trade/logined/transGetByID                           Need to sign

- #### Parameters  list

  | FieldName | Description       | Remark             |
  | --------- | ----------------- | ------------------ |
  | pair      | a pair of trading | "eth-usdt"         |
  | id        | user pending id   | 60006              |
  | key       | apikey            | "K28f43a180373d7e" |
  | nonce     | timestamp         | 1542020339856      |
  | sign      | encrypted string  | encrypted string   |

- #### Return value
 | FieldName | Description | DataType |
 | --------- | ----------- | -------- |
 | P         | price       | float64  |
 | A         | amount      | float64  |
 | T         | time        | string   |

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

4.6 / trade / user transaction flow 
-------------------------

- #### URL: /trade/logined/getTradeHistory                     Need to sign

- #### Parameters list

  | FieldName | Description                                          | Remark             |
  | --------- | ---------------------------------------------------- | ------------------ |
  | pair      | a pair of trading                                    | "eth-usdt"         |
  | limit     | return the number of transactions                    | 0-1000 999         |
  | time      | returns the flow before this time (unix microsecond) | 1544061142176525   |
  | tradeType | user trade type                                      | sell / buy         |
  | key       | apikey                                               | "K28f43a180373d7e" |
  | nonce     | timestamp                                            | 1542020339856      |
  | sign      | encrypted string                                     | encrypted string   |

- #### Return value

  | FieldName | Description                                                                                                                | DataType |
  | --------- | -------------------------------------------------------------------------------------------------------------------------- | -------- |
  | P         | price                                                                                                                      | float64  |
  | A         | amount                                                                                                                     | float64  |
  | T         | time                                                                                                                       | string   |
  | B         | buyer ID                                                                                                                   | int64    |
  | S         | seller ID                                                                                                                  | int64    |
  | Tp        | the direction of the transaction,the direction of the actual transaction,not the direction of the user's pending order | string   |
  | ID        | transaction flow ID, this ID is unique                                                                                     | string   |

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


-   Record:

1.The Tp in the field has nothing to do with the tradeType submitted. It only represents the trade direction of the actual transaction success, and does not represent the direction of the user to hang the order.

 Ex. : He hung a bill, and then he hung a single sell single deal their own bill, then the direction of the deal to sell 

2\. When the tradeType parameter is submitted, only the user's account as (buy/sell order) is returned.

Ex. :

The user hangs a single price is 200 yuan, the quantity is 10 sell sheet, after the user hangs a single again same price and quantity buy a bill, after clinching a deal The user queries the transaction record :

1\) The tradeType field is not submitted, or the field is empty and the return value of the   flow  is :

```
T=15xxxxxxxxxx P=200 A=10 S=123 B=124 Tp=buy ID=873213214cfds3
```

2\) The submitted tradeType field is sell, and the return value is:

```
T=15xxxxxxxxxx P=200 A=10 S=123 B=124 Tp=buy ID=873213214cfds3
```

At this time the return of the turnover, are the user for "sell"  flow 

3\) The submitted tradeType field is buy,, and the return value is:

```
T=15xxxxxxxxxx P=200 A=10 S=123 B=124 Tp=buy ID=873213214cfds3
```

At this time the return of the turnover , are the user "buy"  flow 

As above: the return value is exactly the same on self-dealing. 
Repeated  flow can be obtained by submitting different tradetypes.
