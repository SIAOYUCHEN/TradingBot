### 步驟 1: 安裝 Docker 和 Docker Compose

如果您尚未安裝 Docker 和 Docker Compose，請根據您的操作系統安裝它們。您可以參考 Docker 官方網站上的文檔以進行安裝：[Docker 官方網站](https://docs.docker.com/get-docker/)

### 步驟 2: 啟動服務

在終端機中，切換到包含 Docker Compose 配置文件（通常是 `docker-compose.yml`）的目錄，然後執行以下命令：

```
docker-compose up -d
```

### 步驟 3: 匯入資料表和資料

在你的Sql 編譯器執行或匯入tableData.sql，就會新增相對應資料表和資料


### 步驟 4: 訪問服務
```
go run .
```
一旦服務容器正確啟動，您可以使用您的瀏覽器或 API 工具訪問服務。

1. 網址：http://localhost:8080/swagger/index.html

2. 帳號：test

3. 密碼：123456