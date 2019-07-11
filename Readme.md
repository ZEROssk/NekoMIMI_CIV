#### 鯖にぶん投げたTwitterの画像をwebで見れるものを作る

#### 実装
- ~ユーザー認証の実装~
- webフロントの実装
  - UIなんもわからん
- サーバーサイドの実装
  - APIサーバーとDBサーバーの連携
    - DB
      - DBサーバーのCREATE USERやCREATE DATABASEは.sqlに書かない(本番実行時)
	  - TABLE userID・fileNAME・ID・CreateDATE(2019/07/09時点)
	  - | Example |
		| Center align | Center align | Center align | Center align |
		| ID | userID | fileNAME | CreateDATE |
		|:--:|:--:|:--:|:--:|
		| 1 | a | test-a | YYYY/MM/DD Date-Time |
		| 1 | a | test-a | YYYY/MM/DD Date-Time |
		| ... |
	- API
	  - APIはページ番号を整数値でのみ受け取る。(1.1などはエラーを返す)
	  - 検索用APIの作成
---
#### 要件メモとか
- rootページ(変更の可能性あり)
  - `URL:https://host-name/image_viewer`
- 画像一覧表示のページ(変更の可能性あり)
  - `URL:https://host-name/image_viewer/image?p={PageNum}`
  - `API:https://host-name:port/image_viewer/imgs/twimg/data/{PageNum}`
- userIDで検索した場合の一覧表示ページ(変更の可能性あり)
  - `URL:https://host-name/image_viewer/image/search?id={UserID}`
  - `API:https://host-name:port/image_viewer/imgs/twimg/data/search/{UserID}/{PageNum}`
- 一覧表示から画像の個別表示ページ(変更の可能性あり)
  - `URL:https://host-name/image_viewer/image/original?id={UserID}&img={ImageID}`

