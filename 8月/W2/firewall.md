链接列表

[查看external关键字](external.md)
[查看memory关键字](memory.md)




QA：
1，项目是什么，准备实现什么样的功能
2，功能有哪些模块，已经实现了什么模块
3，这些模块之间的关系是怎样的
4，项目的实现路径是怎样的
5，

智能合约防火墙
目的：为了有效拦截攻击交易
（攻击交易是可能对合约造成损害的交易）
总体上有三个模块：1，Router，2，Register，3，Protection
Router模块：最开始有一个交易项目，防火墙想要对这个交易项目进行检测，需要将这个交易信息路由到防火墙内，是防火墙在被保护的交易项目中的路口（可能在路由前后需要保持原子化操作，不可被跳过）


Router模块实现的事情：
///@notice ProtectInfosize is the function getProtectInfo's returndata size.
    ///@notice info is the function getProtectInfo's returndata info.
    ///@notice is_ProjectPaused 表示从registry中查找项目是否暂停.
    ///@param data 用作接受当前交易调用受保护项目的信息

    function executeWithDetect(bytes memory data) external returns (bool) {
        // 通过代理从registry中获取项目受保护信息
[查看external关键字](external.md)
[查看memory关键字](memory.md)
        bytes memory Proxy_data = abi.encodeWithSignature(
            "getProtectInfo(address,bytes4)",
            msg.sender,
            bytes4(data)
        );
        registry_Proxy.CallOn(Proxy_data);
        // 获取项目受保护信息
        bytes memory ProtectInfo;
        assembly {
            let ProtectInfosize := returndatasize()
            ProtectInfo := mload(0x40)
            mstore(ProtectInfo, ProtectInfosize)
            returndatacopy(add(ProtectInfo, 0x20), 0, ProtectInfosize)
            mstore(0x40, add(add(ProtectInfo, 0x20), ProtectInfosize))
        }
        console.logBytes(ProtectInfo);
        FireWallRegistry.ProtectInfo memory info = abi.decode(
            ProtectInfo,
            (FireWallRegistry.ProtectInfo)
        );
        // 判断是否暂停(项目暂停，函数暂停)
        bytes memory puaseData = abi.encodeWithSignature(
            "pauseMap(address)",
            msg.sender
        );
        registry_Proxy.CallOn(puaseData);
        bytes memory pauseMapInfo;
        // 解析returndata
        assembly {
            let size := returndatasize()
            pauseMapInfo := mload(0x40)
            mstore(pauseMapInfo, size)
            returndatacopy(add(pauseMapInfo, 0x20), 0, size)
            mstore(0x40, add(add(pauseMapInfo, 0x20), size))
        }
        bool is_ProjectPaused = abi.decode(pauseMapInfo, (bool));
        require(!is_ProjectPaused, "project is pause interaction");
        require(!info.is_pause, "project function is pause interaction");
        
        // 遍历
        // 利用接受到的当前交易信息以及获取到的受保护信息对当前交易是否合法进行判断
        for (uint256 index = 0; index < info.enableModules.length; index++) {
            address detectMod = info.enableModules[index];
            // 拆开参数
            string[] memory args = info.params;
            IModule(detectMod).detect(msg.sender, args, data);
        }
        return true;
    }










代码实现：

Register模块：相当于数据库，存储的是项目信息，这个也应该在防火墙内，不能让非权限用户修改，

代码实现：

Protection模块：包含实现了不同防火逻辑的防护模块
实现功能：1，黑名单防护，2，可疑参数防护
待增加：1，价格操纵防护，2，重入防护

代码实现：

工作流程：发起交易，路由模块将交易信息路由到防火墙内，注册模块根据项目信息（具体是？）获取注册模块中的受保护信息，在防护模块检测项目是否应该拦截（是串行检测每一个防护模块中的内容嘛，速度会很慢嘛），若未拦截，检测模块给路由模块传递信息，通过路由模块回到原项目合约继续进行。若拦截，则：：（报出拦截信息）

模块间交互的代码实现细节：


![[Pasted image 20240811095830.png]]
无代理框架
为了实现数据和业务逻辑分开存储所做的改变在哪：


proxyRegister 和 proxyRouter是什么

有代理框架和无代理框架有什么区别

proxy是代理
![[Pasted image 20240811100908.png]]

- **  
    src**
    - **Implemention：核心模块**
        - **Interface**
            - **IModule.sol：防护模块通用接口**
            - **IAuthenicationModule.sol：黑名单防护模块专用接口**
            - **IParamCheckModule.sol：参数防护模块专用接口**
            - **interface.sol：test测试合约使用接口**
        - **AuthenticationModule.sol：黑名单防护模块合约**
        - **AuthenticationModuleV2.sol：黑名单防护模块V2**
        - **ParamCheckModule.sol：参数防护模块合约**
        - **Registry.sol：注册表合约**
        - **RegistryV2.sol：注册表合约V2**
        - **Router.sol：路由合约**
        - **RouterV2.sol：路由合约V2**
    - **proxy：与代理相关的代码实现**
        - **access：实现访问控制代码**
        - **interface：代理部分使用接口**
        - **utils：代理部分使用辅助功能**
        - **ERC1967Proxy.sol：1967代理基础合约**
        - **ERC1967Upgrade.sol：1967合约可升级基础合约**
        - **ProxyAdmin.sol：升级权限管理合约**
        - **proxyForRegistry.sol：注册表合约的代理**
        - **proxyForRouter.sol：路由合约的实现**
    - **example：一些测试合约**
        - **test_contract.sol**
        - **testCoinToken.sol**
        - **testFireWallexp.sol**
- **test：智能合约防火墙链上代码测试**
    - **InteractWithOnChainData.t.sol：使用链上真实攻击案例对防火墙测试代码**
    - **proxy_test.t.sol：对智能合约防火墙代理实现测试代码**
    - **uintTest.t.sol：对智能合约核心模块测试代码**




