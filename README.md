# Trade Engine
> This just for interview homework.
## 目錄
* [Engine](#Engine)
* [Client](#Client)
* [Tool](#Tool)
* [Example](#Example)

## Engine
> 使用 priority queue 機制儲存交易者的訂單，基底結構採用紅黑樹；
> 依 FIFO 規則進行交易匹配。

## Client
> 測試用 Agent，啟用後，可透過 Command Line 進行；
> 1. Buy；
> 2. Sell;
> 3. Cancel;
> 4. QueryOrder;
     > 等操作。

## Tool
> ./proto/generate.bat 執行此工具可以產生所需 proto 檔。

## Example
### Engine
> 執行 engine.go 即可啟動 TradeMatcher
> ![](https://i.imgur.com/5SbVirM.png)


### Client
> 執行 ./client/command_client.go 成功連線後；即可在 Termial 視窗輸入以下指令進行測試。
#### 指令參照
* Buy - **[Cmd] [Stock ID] [Quantity] [Price]**
    * e.g. b 1000 2 500
    * ![](https://i.imgur.com/PIHUplL.png)

* Sell - **[Cmd] [Stock ID] [Quantity] [Price]**
    * e.g. s 1000 10 500
    * ![](https://i.imgur.com/ipFqYi4.png)

* Cancel - **[Cmd] [Stock ID] [Trade ID]**
    * e.g. c 1001
    * ![](https://i.imgur.com/cIMjBx1.png)

* QueryOrder - **[Cmd]**
    * e.g. l
    * ![](https://i.imgur.com/kzLOSwx.png)

