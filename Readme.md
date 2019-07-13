#### 鯖にぶん投げたTwitterの画像をwebで見れるものを作る

#### 適当にメモするところ
- ~ユーザー認証の実装~
- webフロントの実装
  - UIなんもわからん
- サーバーサイドの実装
  - APIサーバーとDBサーバーの連携
    - DB
      - DBサーバーの`CREATE USER`や`CREATE DATABASE`は.sqlに書かない(本番実行時)
	  - TABLE TwitterID・FileNAME・ID・CreateDATE(2019/07/13時点)  
	- API
	  - APIはページ番号を整数値でのみ受け取る。(1.1などはエラーを返す)
	  - 検索用APIの作成
---
#### 要件メモとか
##### API
- rootページ(変更の可能性あり)
  - `URL: https://host-name/image_viewer`
- 画像一覧表示のページ(変更の可能性あり)
  - `URL: https://host-name/image_viewer/twimg/thumbnail?p={PageNum}`
  - `API: GET https://host-name:port/api/v1/twimg/data/page/{PageNum}`  
```json
{
    "Headers": {
        "Content-Type": "application/json; charset=utf-8",
        "Date": "Mon, 1 Jan 2019 00:00:00 GMT",
        "Version": "v1",
        "Server": "host-name",
        "Status": "up"
    },
    "PageNumber": "PageNum",
    "Thumbnail": [
        {
            "TwitterID": "TwiID",
            "FileName": "NAME",
            "ImageID": "ImgID"
        },
        {
            "TwitterID": "TwiID",
            "FileName": "NAME",
            "ImageID": "ImgID"
        }
    ]
}
```

- TwitterIDで検索した場合の一覧表示ページ(変更の可能性あり)
  - `URL: https://host-name/image_viewer/twimg/search?id={TwitterID}`
  - `API: GET https://host-name:port/api/v1/twimg/data/search/{TwitterID}/{PageNum}`  
```json
{
    "Headers": {
        "Content-Type": "application/json; charset=utf-8",
        "Date": "Mon, 1 Jan 2019 00:00:00 GMT",
        "Version": "v1",
        "Server": "host-name",
        "Status": "up"
    },
    "TwitterID": "TwiID",
    "PageNumber": "PageNum",
    "Thumbnail": [
        {
            "FileName": "NAME",
            "ImageID": "ImgID"
        },
        {
            "FileName": "NAME",
            "ImageID": "ImgID"
        }
    ]
}
```

- 一覧表示から画像の個別表示ページ(変更の可能性あり)
  - `URL: https://host-name/image_viewer/image/original?id={TwitterID}&img={ImageID}`
  - `API: GET https://host-name:port/api/v1/twimg/data/original/{TwitterID}/{ImageID}`  
```json
{
    "Headers": {
        "Content-Type": "application/json; charset=utf-8",
        "Date": "Mon, 1 Jan 2019 00:00:00 GMT",
        "Version": "v1",
        "Server": "host-name",
        "Status": "up"
    },
    "Image": [
        {
            "TwitterID": "TwiID",
            "FileName": "NAME",
            "ImageID": "ImgID"
        },
     ]
}
```

##### DB
  - Example

| ID  | TwitterID | FileNAME | CreateDATE               |
|:---:|:---------:|:--------:|:------------------------:|
| 1   | a         | test-a   | YYYY/MM/DD Add-Date-Time |
| 2   | b         | test-b   | YYYY/MM/DD Add-Date-Time |
| 3   | c         | test-c   | YYYY/MM/DD Add-Date-Time |
| ・  | ・        | ・       | ・                       |
| ・  | ・        | ・       | ・                       |
| ・  | ・        | ・       | ・                       |

