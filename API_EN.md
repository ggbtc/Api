# **API  Document**

## **2018-10-11**

## 1.1 /  Private interface signature process 

#### Interface: Get user recharge address 

url: https://private.ggbtc.com:55558/user/logined/getNewAddress

- #### Parameters  list:

Token got from server  : token={key:keykey,value:valuevalue} 

| FieldName | Description        |
| --------- | ------------------ |
| key       | keykey             |
| value     | valuevalue         |
| nonce     | 12343546           |
| currency  | eth                |
| sign      | encrypted   string |

- #### Signature process 

URL of formatting string's request:

```
/user/logined/getNewAddress?key=keykey&nonce=12343536&currency=eth
```

The string above is followed by 'value' 

```
/user/logined/getNewAddress?key=keykey&nonce=12343536&currency=ethvaluevalue
```

Sha512 encryption of the above string, the result is :

```
dea67c903dd1af28270f68f9b35c52ef6597d2082694189adbdf9ad805b10e0c5ac73461cfd08e292036ac1a8fb626fe7a5bbd76d5fda33a112a707a96404c87
```

Final full request's URL: 

```
http://apiserver/user/logined/getNewAddress?key=keykey&nonce=12343536&currency=eth&sign=dea67c903dd1af28270f68f9b35c52ef6597d2082694189adbdf9ad805b10e0c5ac73461cfd08e292036ac1a8fb626fe7a5bbd76d5fda33a112a707a96404c87
```

-   All interfaces marked for signature ,must follow the above procedure to sign and add the required parameters ,the following interface details will not repeated .

1.2 / Standard return value 
----------------

- ```
    {
     	  "message":	"success", 			//'success' represets success, if there are other 										//results, it will be reflected
     	  "result":		{} 				    //Data is needed only for success
    }
    
    ```



-   All interfaces have return values that represent data in'result'. If  no , comment the standard return value. 

2.0 /Get user token（long-term）
---------------------------

- #### URL: /user/logined/getLongToken		Need to sign 

- #### Parameters  list: 

  ​			none

- #### Return value:

  | FieldName | DataType |
  | --------- | -------- |
  | Key       | string   |
  | Value     | string   |

2.1 / Get user information
------------------

- #### URL: /user/logined/getInfo			 Need to sign 

- #### Parameters  list: 

  ​			none

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

2.2 / Request user recharge address
----------------------

- #### URL: /user/logined/getNewAddress 			 Need to sign 

   #### Parameters list	

| FieldName | Description | Remark |
| --------- | ----------- | ------ |
| currency  | currency    | "eth"  |

- #### Return value

  | FieldName | Description | Remark |
  | --------- | ----------- | ------ |
  | Address   | Address     | string |

2.3 / Get user balance
------------------

- #### 	URL: /user/logined/getBalance 			Need to sign 

- #### Parameters list 

  ​						none

- #### Return value

​                        ~~Balance           user balance              HashMap~~

```
{

	"eth":				1000000,				 //eth available balance

	"eth-trade":		10111, 					//eth locked balance

	"eth-locked":		100, 					//eth locked balance

	"btc":				11123456, 				//btc available balance

	"btc-trade":		11123456, 				//btc locked balance

	"hkdt":				11123456,			   //hkdt available balance

	"hkdt-trade":		11123456 			   //hkdt locked balance(locked balance in trade)

}

```



//~~record:Each of the above fields may not exist, the nonexistent fields are treated as 0, and the above data contains 6 decimal places~~ 

//record1:'-trade' and '-locked'  both represent locked balances 

//record2: When the balance is empty, the field may not exist 



3.1 / Trade / get trade history
-----------------------------

- #### URL: /trade/public/getTradeHistory

	 #### Parameters list	

| FieldName | Description                    | Remark   |
| --------- | ------------------------------ | -------- |
| pair      | a pair of trading              | eth-usdt |
| limit     | number of branches (optional ) | 100      |

- #### Return value

```
[
	T 		time			 unixTime

	Tp 		type(trade) 	 sell,buy

	P 		price 			 float64

	A 		quantity

	B		buyID

	S 		sellID
]
```

​	

3.2 / trade / get 24 hours of data 
---------------------------

- #### URL: /trade/public/getTicker

	 #### Parameters list	

| FieldName | Description       | Remark   |
| --------- | ----------------- | -------- |
| pair      | a pair of trading | eth-usdt |

- #### Return value

  Returns	map\[string\]jsonStr when the pair is all

  Returns  jsonStr when the pair  for a single currency 

```
{

	"Chg": 		"-0.398026", 			//The proportion of increase or decrease

	"High":		 "201.00541", 			//24 hour's high

	"Low": 		"121.00003",	 		//24 hour's low

	"Name": 	"eth-usdt", 			//currency

	"Turn": 	"6331.0673987", 		//transaction amount in one hour(as-The number of 										  //transactions in subsequent currencies)

	"Vol": 		"31.497", 			   //24 hours volume(as-The number of transactions in the 									    //previous currency)

	"Last": 	"121.000003"			//last transaction price

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

```
{

	Sell :[ 			//sell

	{P:100.0,A:30},

		.....

	]

	Buy[], 				//buy

	Last:	123.1 		//latest transaction price

}

```



3.4 / trade /getLine
---------------------

- #### URL: /trade/public/getKline

	 #### Parameters list	

| FieldName | Description                    | DataType             |
| --------- | ------------------------------ | -------------------- |
| start     | StartTime(optional)            | unixTime             |
| end       | EndTime(optional)              | unixTime             |
| basis     | unit                           | hour,minute,day,week |
| limit     | number of branches (optional ) | 100                  |

- #### Return value

```
[

	T 		time 		 unixTime

	S 		startPrice	float64

	E 		endPrice 	float64

	H 		peak price

	L 		bottom price

	A 		quantity

]

```



4.1 /trade / orderCreate
---------------------

- #### URL: /trade/logined/orderCreate                             Need to sign

   #### Parameters list	

| FieldName | Description       | Remark     |
| --------- | ----------------- | ---------- |
| pair      | a pair of trading | "eth-usdt" |
| price     | order price       | 222.3      |
| amount    | order quantity    | 10         |
| tradeType | trade type        | "sell"     |

- #### Return value

  tradeID(success)

4.2 / trade / userCancelOrder
-------------------------

- #### URL: /trade/logined/orderCancel 		               Need to sign

- #### Parameters list

  | FieldName | Description       | Remark     |
  | --------- | ----------------- | ---------- |
  | pair      | a pair of trading | "eth-usdt" |
  | id        | orderID           | 612342     |

- #### Return value

Standard return value 

asynchronous  processing  ,so be sure to return  success

4.3 / trade / user query order
-------------------------

- #### URL: /trade/logined/orderGet 		          Need to sign

	 #### Parameters list	

| FieldName | Description       | Remark     |
| --------- | ----------------- | ---------- |
| pair      | a pair of trading | "eth-usdt" |
| id        | orderID           | 612342     |

- #### Return value

  ```
  {
      id 					//id
  
  	user
  
  	pair
  
  	tradetype
  
  	orderprice 			//order price
  
  	orderamount 		//order quantity-int-The last six digits are the decimal points
  
  	time
  
  	feescale 			//commission ratio
  
  	remainamount		//remaining unsold quantity-int-The last six digits are the 						   //decimal points
  
  	feeamount		    // This is the user selling the eth and getting the amount  						   //of the usdt
  
  	fillamount 			// This is the number of eth's that the user is not using
  	avgprice 			//average transaction price
  
  	state 				//orderState		 open,closed,canceled
  }
  ```

4.4 / trade / user query the current pending order 
-----------------------------

- #### URL: /trade/logined/orderGetUser		 Need to sign

	 #### Parameters list	

| FieldName | Description                                  | Remark                |
| --------- | -------------------------------------------- | --------------------- |
| pair      | a pair of trading                            | "eth-usdt"            |
| limit     | query quantity(optional)                     | 100                   |
| time      | query deadline(optioanal)                    | unixTime  millisecond |
| tradeType | Check the buy order or sell order (optional) | sell,buy              |
| state     | query order state(optional)                  | open,closed,canceled  |

- #### Return value

```
{
    id 					//id

	user

	pair

	tradetype

	orderprice 			//order price

	orderamount 		//order quantity   -int-  The last six digits are the decimal points

	time

	feescale 			//commission ratio

	remainamount		//remaining unsold quantity  -int-  The last six digits are the 						   //decimal points

	feeamount			//This is the user selling the eth and getting the amount of 						 	    //the usdt			

	fillamount 			//This is the number of eth's that the user is not using
	avgprice 		    //average transaction price
	state 				//order state		 open,closed,canceled
}
```

4.5 / trade / user pending transaction details 
-----------------------------

- #### URL: /trade/logined/transGetByID                           Need to sign

- #### Parameters  list

  | FieldName | Description       | Remark     |
  | --------- | ----------------- | ---------- |
  | pair      | a pair of trading | "eth-usdt" |
  | id        | user pending id   | 60006      |

- #### Return value

```
[

	{

		T 		time 		unixTime

		P 		price 		float64

		A 		quantity

	}...

]

```



4.6 / trade / user transaction flow 
-------------------------

- #### URL: /trade/logined/getTradeHistory                     Need to sign

- #### Parameters list

  | FieldName | Description                                          | Remark           |
  | --------- | ---------------------------------------------------- | ---------------- |
  | pair      | a pair of trading                                    | "eth-usdt"       |
  | limit     | return the number of transactions                    | 0-1000 999       |
  | time      | returns the flow before this time (unix microsecond) | 1544061142176525 |
  | tradeType | user trade type                                      | sell \| buy      |

- #### Return value

```
[

	{

		T 		time 		unixTime

		P 		price		 float64

		A 		quantity

		B 		buyerID

		S 		sellerID

		Tp 		trade direction,The direction that trades clinch a deal actually, not be the 				 hang sheet direction of the user

		ID 		Transaction  flow ID and this ID is unique

	}...

]

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
