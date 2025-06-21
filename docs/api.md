# Open API

## `GET /users/me`

operationID: `getMe`

自分の情報を取得する。NeoShowcaseの機能を使ってtraQ IDを取得する。

ローカル環境では、クエリパラメータ `debugUserName` を指定することで、そのユーザーの`name` と `iconUrl` を取得できる。指定しなかった場合は `traP` が返る。
他のエンドポイントでもこのローカル用機能は有効であり、指定されたユーザーとしてリクエストが処理される。

### Response

成功 200

```json
{
  "name": "ikura-hamu",
  "iconUrl": "https://q.trap.jp/api/v3/public/icon/ikura-hamu"
}
```

## `POST /games`

operationID: `createGame`

ゲームの開始をリクエストする。

### Request

なし

### Response

成功 204

## `WebSocket /games/ws`

operationID: `waitGameWS`

[`POST /games`](#post-games)でゲームを開始した後に接続する。

### クライアントが受け取るイベント

マッチングが成立したときに、以下のJSONが送信される。
playerIdは、ゲーム内でのidを表し、0か1である。

```json
{
  "gameId": "54f77f89-4eba-4b3d-8744-d6698c3581a9",
  "playerId": 1
}
```

### サーバーが受け取るイベント

サーバーはイベントを受け取らない

## `WebSocket /games/{gameId}/ws`

operationID: `gameWS`

[`WebSocket /games/ws`](#websocket-gamesws) でマッチングが成立した後に接続する。

### クライアントが受け取るイベント

`type` フィールドでイベントの種類を識別する。

#### `gameReady`

ゲームの準備ができたとき。
`startTime` は`gameStarted`イベントが送信されるおおよその時刻を表す。

```json
{
  "type": "gameReady",
  "fieldCards": [
    {
      "id": "36d01514-5e67-40a1-8199-e73406aec7f5",
      "type": "operand",
      "value": 5
    }
  ],
  "playerId": 1,
  "player0": [],
  "player0HandsLimit": 10,
  "player1": [],
  "player1HandsLimit": 10,
  "currentPlayerId": 0,
  "player0Score": 0,
  "player1Score": 0,
  "startTime": "2023-10-01T12:00:00Z"
}
```

#### `gameStarted`

ゲームが開始されたとき

```json
{
  "type": "gameStarted",
  "currentPlayerId": 0,
  "turn": 1,
}
```

#### `turnStarted`

新しいターンが開始されたとき

```json
{
  "type": "turnStarted",
  "currentPlayerId": 0,
  "turn": 2,
  "turnTimeRemaining": 30
}
```

#### `cardsUpdated`

field cardsもしくはhand cardsが更新されたとき

```json
{
  "type": "cardsUpdated",
  "fieldCards": [
    {
      "id": "36d01514-5e67-40a1-8199-e73406aec7f5",
      "type": "operand",
      "value": 5
    }
  ],
  "player0": [
    {
      "id": "9568cff9-f9ce-4ee3-a9b9-1751f772a640",
      "type": "operator",
      "value": "+"
    }
  ],
  "player0HandsLimit": 10,
  "player1": [
    {
      "id": "c8192d7e-4a88-4454-9894-58d5046feadc",
      "type": "item",
      "value": "clearFieldCards"
    }
  ],
  "player1HandsLimit": 10
}
```

#### `turnTimeRemainingChanged`

ターンの残り時間が変更されたとき。1秒ごとに送信される。
残り時間が 0 のときは送信されない。

```json
{
  "type": "turnTimeRemainingChanged",
  "currentPlayerId": 0,
  "remainingSeconds": 4
}
```

#### `submissionSucceeded`

式の送信が成功したとき。
ここでの`score`は、このsubmissionによって加算された得点を表し、合計点でないことを注意する必要がある。
合計点は `scoreUpdated` イベントで送信される。

```json
{
  "type": "submissionSucceeded",
  "playerId": 0,
  "expression": "5 * 2",
  "score": 1
}
```

#### `scoreUpdated`

得点が更新されたとき。

```json
{
  "type": "scoreUpdated",
  "player0": 10,
  "player1": 8
}
```

#### `turnEnded`

ターンが終了したとき。ターンの残り時間が0になったときに送信される。
最終ターンの場合は、`nextTurn` は `null` になる。

```json
{
  "type": "turnEnded",
  "nextPlayerId": 0,
  "nextTurn": 1
}
```

#### `gameEnded`

ゲームが終了したとき。最終ターンの`turnEnded`の後に送信される。

```json
{
  "type": "gameEnded",
  "player0": 10,
  "player1": 8
}
```

### サーバーが受け取るイベント

無し

## `POST /games/{gameId}/picks`

operationID: `pickCard`

field cardsからカードを1枚選んでhand cardsに加える。

### Request

```json
{
  "cardId": "36d01514-5e67-40a1-8199-e73406aec7f5"
}
```

### Response

成功 204

自分のターンではない 400
gameIdとcardIdの組み合わせが正しくない 400

## `POST /games/{gameId}/items`

operationID: `useItem`

hand cardsからアイテムカードを1枚選んで、その効果を使う。

### Request

```json
{
  "cardId": "36d01514-5e67-40a1-8199-e73406aec7f5"
}
```

### Response

成功 204

自分のターンではない 400
gameIdとcardIdの組み合わせが正しくない 400
アイテムカードではない 400

## `POST /games/{gameId}/submissions`

operationID: `submitExpression`

式を送信する
expressionは、数字と`+`, `-`, `*`, `/` ,`(`, `)`を含む文字列である必要がある。

### Request

```json
{
  "expression": "5 * 2",
  "cards": [
    "36d01514-5e67-40a1-8199-e73406aec7f5"
  ]
}
```

### Response

成功 200

値が10でない場合も200を返すことに注意する。10でなかった場合は、`success`が`false`になる。

```json
{
  "success": true,
  "value": 10
}
```

自分のターンでない 400

## `POST /games/{gameId}/clear`

operationID: `clearHandCards`

hand cardsをクリアする。

### Request

なし

### Response

成功 204

自分のターンでない 400

## `GET /games/{gameId}/results`

operationID: `getGameResult`

ゲームの結果を取得する。

### Response

成功 200

```json
{
  "gameId": "54f77f89-4eba-4b3d-8744-d6698c3581a9",
  "player0Name": "ikura-hamu",
  "player1Name": "another-player",
  "player0Score": 10,
  "player1Score": 8,
  "player0SuccessExpressions": [
    "5 * 2"
  ],
  "player1SuccessExpressions": [
    "3 + 5 + 2"
  ],
}
```

まだゲームが終了していない 404

## `GET /ranking`

operationID: `getRanking`

ランキングを取得する。

### Request

クエリパラメータ `limit` (optional, default: 20) で取得するランキングの数を指定できる。

### Response

成功 200

`count` はランキング全体の数。1回もプレイしていないユーザーは含まない。

```json
{
  "count": 2,
  "ranking": [
    {
      "name": "ikura-hamu",
      "iconUrl": "https://q.trap.jp/api/v3/public/icon/ikura-hamu",
      "wins": 5,
      "losses": 2,
      "totalScore": 100
    },
    {
      "name": "another-player",
      "iconUrl": "https://q.trap.jp/api/v3/public/icon/another-player",
      "wins": 3,
      "losses": 4,
      "totalScore": 80
    }
  ]
}
```
