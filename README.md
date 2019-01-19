# VideoManager（名称未定）
動画自動保存LINE BOT & 閲覧アプリ

## Todo
- [ ] CSSフレームワークの選定

### ざっくりな動き（ミニマム）
0. LINE BOTをグループに追加
1. 動画をグループに投稿
2. 投稿をトリガーにLINE BOTに設定してあるwebhook URLが叩かれる
3. heroku上で稼働しているGoサーバーにリクエストが届く
4. リクエストに含まれる、`type`が`video` かつ `contentProvider.type`が`line` の場合、動画データを`content`エンドポイントから取得する
    - リクエストの `id (messageId)` をキーにGET
    - `GET https://api.line.me/v2/bot/message/{messageId}/content`
    - API doc: https://developers.line.biz/ja/reference/messaging-api/#wh-video
5. 取得した動画データをgoogle photo apisにRESTで投げることで、google photoに動画を保存
6. Vue.jsで作ったWebページをデプロイ -> 動画閲覧
    - indexページには各ビデオページへのリンクを全件表示する
    - 各ビデオページにはvideoタグで表示

## メモ
- 動画サムネイル生成

- google photo apis -> 10000 req/day
  - https://developers.google.com/photos/library/guides/api-limits-quotas
  
- CSSフレームワーク
  - とりあえずレスポンシブ対応しているやつなら何でも良い（はず）

- 個人開発でデザインに苦戦しないためのデザインテクニック
  - https://iritec.jp/design/19604/
