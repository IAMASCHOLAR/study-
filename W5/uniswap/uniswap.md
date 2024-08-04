Uniswap
消息源：网页
目标：
1，uniswap的发展历程
2，uniswap为什么会被提出
3，uniswap有哪些应用
4，uniswap的源码

之前不太好用的方法是EtherDelta

uniswap是最大的去中心化（dex）交易协议，

uniswap V0概念阶段，仅适用于单一 ETH/ERC 20 交易对。仅适用于单一流动性提供者（LP）。

Uniswap V1 仅仅ETH 和Token直接交换----2018年1月

Uniswap V2 允许任何 ERC20 之间直接交换以及链式交换----2020 年 3 月

uniswapV3设计了限定区间的Concentrated Liquidity（集中式流动性），为了解决上述V2的资金效率问题


V1的问题在于所有代币需要先和ETH进行交换，对于频繁DAI/USDT货币效率低下

V2使得ERC-20代币直接可以交换

V2的问题，效率低下，1，在处理低波动性代币时，仅在较小的价格范围内才需要流动性。然而，当前的设计在所有价格范围内均匀分配流动性。

2，大额订单对价格造成冲击，

3，代币对的局限性，每个流动型池至少有2种代币，

4，无常损失，当流动性提供者加入流动性池时，代币的价格比率是固定的，但如果市场条件变化导致价格比率变动，他们可能会在撤出流动性时遭受损失。

V3：


QA：
工厂合约
AMM机制
价格预言机
TWAP
滑点
滑点问题
永续合约
期权


uniswap  白皮书
https://binschool.app/uniswap-whitepaper/uniswap-whitepaper-v1.html

