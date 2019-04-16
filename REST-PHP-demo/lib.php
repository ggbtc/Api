<?php

class ggbtc {
    private $api = 'http://apitest.ggbtc.cc:3030'; //todo : replace the real api server address
    private $apiKey = 'K5e3c3825ede07bb';   //todo : replace your apikey
    private $secretKey = 'V9923e2213085aea2';   //todo : replace your secretkey
    public $limit = 30;
    public $apiMethod = '';

    function __construct() {
    }

    //获取订单薄
    function getOrderBook($pair) 
    {
        $params = [];
        $params['pair'] = $pair;
        $params['limit'] = $this->limit;

        $this->apiMethod = '/trade/public/getOrderbook';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, '');
    
        $res = $this->curl($this->api . $strReqUrl);
        return $res;
    }

    //获取用户充值地址
    function getNewAddress($currency)
    {
        $params = [];
        $params['currency'] = $currency;   
        $params['key'] = $this->apiKey;

        $this->apiMethod = '/user/logined/getNewAddress';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, $this->secretKey);

        $res = $this->curl($this->api . $strReqUrl);
        return $res;
    }

    //获取用户信息
    function getUserInfo()
    {
        $params = [];
        $params['key'] = $this->apiKey;

        $this->apiMethod = '/user/logined/getInfo';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, $this->secretKey);

        $res = $this->curl($this->api . $strReqUrl);
        return $res;   
    }

    //获取用户余额
    function getUserBalance()
    {
        $params = [];
        $params['key'] = $this->apiKey;

        $this->apiMethod = '/user/logined/getBalance';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, $this->secretKey);

        $res = $this->curl($this->api . $strReqUrl);
        return $res;    
    }

    //获取交易历史记录
    function getTradeHistory($pair)
    {
        $params = [];
        $params['pair'] = $pair;
        $params['limit'] = $this->limit;

        $this->apiMethod = '/trade/public/getTradeHistory';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, '');

        $res = $this->curl($this->api . $strReqUrl);
        return $res;    
    }

    //获取24小时交易数据
    function getTicker($pair)
    {
        $params = [];
        $params['pair'] = $pair;

        $this->apiMethod = '/trade/public/getTicker';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, '');

        $res = $this->curl($this->api . $strReqUrl);
        return $res;    
    }

    //获取K线数据
    function getKline($pair, $basis)
    {
        $params = [];
        $params['pair'] = $pair;
        $params['basis'] = $basis;
        $params['limit'] = $this->limit;
        // $params['start'] = 1555391850775;
        // $params['end'] = 1555437883220;

        $this->apiMethod = '/trade/public/getKline';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, '');

        $res = $this->curl($this->api . $strReqUrl);
        return $res;    
    }

    //用户下单
    function createOrder($price, $amount, $pair, $tradeType)
    {
        $params = [];
        $params['pair'] = $pair;
        $params['price'] = $price;
        $params['amount'] = $amount;
        $params['tradeType'] = $tradeType;
        $params['key'] = $this->apiKey;

        $this->apiMethod = '/trade/logined/orderCreate';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, $this->secretKey);

        $res = $this->curl($this->api . $strReqUrl);
        return $res;  
    }

    //取消下单
    function cancelOrder($pair, $id)
    {
        $params = [];
        $params['pair'] = $pair;
        $params['id'] = $id;
        $params['key'] = $this->apiKey;

        $this->apiMethod = '/trade/logined/orderCancel';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, $this->secretKey);

        $res = $this->curl($this->api . $strReqUrl);
        return $res;   
    }

    //用户查询订单
    function GetOrder($pair, $id)
    {
        $params = [];
        $params['pair'] = $pair;
        $params['id'] = $id;
        $params['key'] = $this->apiKey;

        $this->apiMethod = '/trade/logined/orderGet';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, $this->secretKey);

        $res = $this->curl($this->api . $strReqUrl);
        return $res;  
    }

    //用户条件查询多条订单详情
    function getOrders($pair)
    {
        $params = [];
        $params['pair'] = $pair;
        $params['limit'] = $this->limit;
        $params['key'] = $this->apiKey;

        $this->apiMethod = '/trade/logined/orderGetUser';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, $this->secretKey);

        $res = $this->curl($this->api . $strReqUrl);
        return $res; 
    }

    //用户挂单交易详情
    function getTransByID($pair, $id)
    {
        $params = [];
        $params['pair'] = $pair;
        $params['id'] = $id;
        $params['key'] = $this->apiKey;

        $this->apiMethod = '/trade/logined/transGetByID';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, $this->secretKey);

        $res = $this->curl($this->api . $strReqUrl);
        return $res; 
    }

    //获取用户历史交易
    function getUserTradeHistory($pair)
    {
        $params = [];
        $params['pair'] = $pair;
        $params['limit'] = $this->limit;
        $params['key'] = $this->apiKey;

        $this->apiMethod = '/trade/logined/getTradeHistory';
        $strReqUrl = $this->getReqUrl($this->apiMethod, $params, $this->secretKey);

        $res = $this->curl($this->api . $strReqUrl);
        return $res; 
    }

    //拼接请求参数
    function bindParam($params)
    {
        $s = '';
        foreach($params as $key => $v) {
            $s .= $key . '=' . $v . '&';
        }

        $s = rtrim($s, '&');
        return $s;
    }

    //拼接请求字符串
    function getReqUrl($apiMethod, $params, $secretKey) 
    {
        $isSign = false;
        if (array_key_exists('key', $params)) {
            $isSign = true;
        }

        $stringParams = $this->bindParam($params);
        //拼接请求方法和参数
        $strReqUrl = $apiMethod . '?' . $stringParams;
        //加密获取加密字符串
        if ($isSign) {
            $sign = hash('sha512', $strReqUrl . $secretKey);
            $strReqUrl .= '&sign=' . $sign;
        }

        return $strReqUrl;
    }

    //get请求
    function curl ($url)
    {
		$ch = curl_init();
		curl_setopt($ch,CURLOPT_URL, $url);
		curl_setopt($ch,CURLOPT_RETURNTRANSFER,1);
		curl_setopt($ch,CURLOPT_HEADER,0);
		curl_setopt($ch, CURLOPT_TIMEOUT,60);
		$output = curl_exec($ch);
        $info = curl_getinfo($ch);
		curl_close($ch);
		return $output;
    }
}