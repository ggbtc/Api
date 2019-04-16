<?php

// 定义参数
// define('ACCESS_KEY','K5e3c3825ede07bb'); // your ACCESS_KEY , had set at lib.php
// define('SECRET_KEY', 'V9923e2213085aea2'); // your SECRET_KEY , had set at lib.php

//加载类库
require_once("lib.php");

//实例化类库
$ggbtc = new ggbtc();

echo $ggbtc->getOrderBook('eth-usdt');
// echo $ggbtc->getNewAddress('eth');
// echo $ggbtc->getUserInfo();
// echo $ggbtc->getUserBalance();
// echo $ggbtc->getTradeHistory('eth-usdt');
// echo $ggbtc->getTicker('eth-usdt');
// echo $ggbtc->getKline('eth-usdt','m');
// echo $ggbtc->createOrder(210, 5.9, 'eth-usdt', 'buy');
// echo $ggbtc->cancelOrder('eth-usdt', 6045001);
// echo $ggbtc->GetOrder('eth-usdt', 6045001);
// echo $ggbtc->getOrders('eth-usdt');
// echo $ggbtc->getTransByID('eth-usdt', 6027830);
// echo $ggbtc->getUserTradeHistory('eth-usdt');