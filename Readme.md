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
	- API
	  - APIはページ番号を整数値でのみ受け取る。(1.1などはエラーを返す)
	  - 検索用APIの作成(未完了)
