# DB スキーマ

## `games`

| 列名       | 型                                     | 制約            | 既定値              | 説明        |
| ---------- | -------------------------------------- | --------------- | ------------------- | ----------- |
| id         | `VARCHAR(36)`                          | **PK**          | ―                   | ゲーム UUID |
| status     | `ENUM('waiting','running','finished')` | **NOT NULL**    | ―                   | ゲーム状態  |
| started_at | `DATETIME`                             |                 | `NULL`              | 開始時刻    |
| ended_at   | `DATETIME`                             |                 | `NULL`              | 終了時刻    |
| created_at | `DATETIME`                             |                 | `CURRENT_TIMESTAMP` | 作成時刻    |
| **索引**   | ―                                      | `INDEX(status)` |                     | 状態検索用  |

## `waiting_players`

| 列名       | 型            | 制約                             | 既定値 | 説明         |
| ---------- | ------------- | -------------------------------- | ------ | ------------ |
| id         | `INT`         | **PK**, **AUTO_INCREMENT**       | ―      | 行連番       |
| user_name  | `VARCHAR(32)` | **PK**, **UNIQUE**, **NOT NULL** | ―      | 参加ユーザー |
| created_at | `DATETIME`    | **NOT NULL**                     | ―      | 作成時刻     |

## `game_players`

| 列名      | 型            | 制約                                          | 既定値 | 説明               |
| --------- | ------------- | --------------------------------------------- | ------ | ------------------ |
| game_id   | `VARCHAR(36)` | **PK(1)**, **FK → games.id**, **NOT NULL**    | ―      | ゲーム UUID        |
| player_id | `TINYINT`     | **PK(2)**, **NOT NULL**                       | ―      | スロット番号 (0/1) |
| user_name | `VARCHAR(32)` | **UNIQUE (game_id, user_name)**, **NOT NULL** | ―      | 参加ユーザー       |
| score     | `INT`         |                                               | `0`    | 現在のスコア         |

PK(1,2) = 複合主キー (game_id, player_index)

## `expressions`

| 列名         | 型            | 制約                                       | 既定値 | 説明             |
| ------------ | ------------- | ------------------------------------------ | ------ | ---------------- |
| id           | `VARCHAR(36)` | **PK**                                     | ―      | 行連番           |
| game_id      | `VARCHAR(36)` | **FK → games.id**, **INDEX**, **NOT NULL** | ―      | ゲーム           |
| player_id    | `TINYINT`     | **NOT NULL**                               | ―      | プレイヤーの番号 |
| expression   | `TEXT`        | **NOT NULL**                               | ―      | 入力式           |
| value        | `TEXT`         | **NOT NULL**                               | ―      | 計算結果         |
| points       | `INT`         | **NOT NULL**                               | ―      | 加算点           |
| success      | `BOOLEAN`     | **NOT NULL**                               | ―      | 目標達成か       |
| submitted_at | `DATETIME`    | **NOT NULL**                               | ―      | 送信時刻         |

## `cards`

| 列名            | 型                                  | 制約                                       | 既定値 | 説明                         |
| --------------- | ----------------------------------- | ------------------------------------------ | ------ | ---------------------------- |
| id              | `VARCHAR(36)`                       | **PK**                                     | ―      | カード UUID                  |
| game_id         | `VARCHAR(36)`                       | **FK → games.id**, **INDEX**, **NOT NULL** | ―      | ゲーム                       |
| type            | `ENUM('operand','operator','item')` | **NOT NULL**                               | ―      | 種別                         |
| value           | `VARCHAR(32)`                       | **NOT NULL**                               | ―      | 数字・演算子、アイテム名など |
| owner_player_id | `TINYINT`                           |                                            | `NULL` | 0/1/NULL                     |
| location        | `ENUM('field','hand','used')`       | **NOT NULL**                               | ―      | 現在位置                     |

## `hand_cards_limits`

| 列名       | 型            | 制約                                    | 既定値 | 説明               |
| ---------- | ------------- | --------------------------------------- | ------ | ------------------ |
| game_id    | `VARCHAR(36)` | **PK**, **FK → games.id**, **NOT NULL** | ―      | ゲーム UUID        |
| player_id  | `TINYINT`     | **PK**, **NOT NULL**                    | ―      | スロット番号 (0/1) |
| hand_cards | `TINYINT`     | **NOT NULL**                            | ―      | hand cardsの上限   |

## `field_cards_limits`

| 列名        | 型            | 制約                                    | 既定値 | 説明              |
| ----------- | ------------- | --------------------------------------- | ------ | ----------------- |
| game_id     | `VARCHAR(36)` | **PK**, **FK → games.id**, **NOT NULL** | ―      | ゲーム UUID       |
| field_cards | `TINYINT`     | **NOT NULL**                            | ―      | field cardsの上限 |

## `turns`

| 列名        | 型            | 制約                                    | 既定値 | 説明               |
| ----------- | ------------- | --------------------------------------- | ------ | ------------------ |
| game_id     | `VARCHAR(36)` | **PK**, **FK → games.id**, **NOT NULL** | ―      | ゲーム UUID        |
| player_id   | `TINYINT`     | **NOT NULL**                    | ―      | スロット番号 (0/1) |
| turn_number | `INT`         | **PK**, **NOT NULL**                            | ―      | ターン番号         |
| start_at    | `DATETIME`    | **NOT NULL**                            | ―      | 開始時刻           |
| end_at      | `DATETIME`    | **NOT NULL**                            | ―      | 終了時刻           |
