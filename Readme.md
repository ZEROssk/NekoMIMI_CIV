#### 鯖にぶん投げたTwitterの画像をwebで見れるものを作る

#### 適当にメモするところ
- ~ユーザー認証の実装~
- webフロントの実装
  - UIなんもわからん
- サーバーサイドの実装
  - APIサーバーとDBサーバーの連携
    - DB
      - DBサーバーのCREATE USERやCREATE DATABASEは.sqlに書かない(本番実行時)
	  - TABLE userID・fileNAME・ID・CreateDATE(2019/07/09時点)  
	- API
	  - APIはページ番号を整数値でのみ受け取る。(1.1などはエラーを返す)
	  - 検索用APIの作成
---
#### 要件メモとか
##### API
- rootページ(変更の可能性あり)
  - `URL: https://host-name/image_viewer`
- 画像一覧表示のページ(変更の可能性あり)
  - `URL: https://host-name/image_viewer/images/thumbnail?p={PageNum}`
  - `API: https://host-name:port/v1/image_viewer/imgs/twimg/data/{PageNum}`  
```json
{
    headers: {
        'Content-Type': 'application/json; charset=utf-8'
        'Date': 'Mon, 1 Jan 2019 00:00:00 GMT'
        'Version': 'v1'
        'Server': 'host-name'
        'Status': 'up'
    }
    PageNumber: PageNum
    Thumbnail: [
        {
            fileName: "NAME",
            userID: "ID"
        },
        {
            fileName: "NAME",
            userID: "ID"
        }
    ]
}
```

- userIDで検索した場合の一覧表示ページ(変更の可能性あり)
  - `URL: https://host-name/image_viewer/images/search?id={UserID}`
  - `API: https://host-name:port/v1/image_viewer/imgs/twimg/data/search/{UserID}/{PageNum}`  
```json
{
    headers: {
        'Content-Type': 'application/json; charset=utf-8'
        'Date': 'Mon, 1 Jan 2019 00:00:00 GMT'
        'Version': 'v1'
        'Server': 'host-name'
        'Status': 'up'
    }
    UserID: "userID"
    PageNumber: PageNum
    Thumbnail: [
        {
            fileName: "NAME",
            userID: "ID"
        },
        {
            fileName: "NAME",
            userID: "ID"
        }
    ]
}
```

- 一覧表示から画像の個別表示ページ(変更の可能性あり)
  - `URL: https://host-name/image_viewer/image/original?id={UserID}&img={ImageID}`
  - `API: https://host-name:port/v1/image_viewer/imgs/twimg/data/original?id={UserID}&img={ImageID}`  
```json
{
    headers: {
        'Content-Type': 'application/json; charset=utf-8'
        'Date': 'Mon, 1 Jan 2019 00:00:00 GMT'
        'Version': 'v1'
        'Server': 'host-name'
        'Status': 'up'
    }
    Image: [
        {
            fileName: "NAME",
            userID: "ID"
        }
    ]
}
```

##### DB
  - Example

|:-------:|:-------:|:--------:|:----------------------:|
| ID      | userID  | fileNAME | CreateDATE             |
| 1       | a       | test-a   | YYYY/MM/DD Date-Time   |
| 2       | b       | test-b   | YYYY/MM/DD Date-Time   |

