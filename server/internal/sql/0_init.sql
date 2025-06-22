-- ゲームテーブル
CREATE TABLE games (
    id VARCHAR(36) PRIMARY KEY,
    status ENUM('waiting', 'running', 'finished', 'canceled') NOT NULL,
    started_at DATETIME NOT NULL,
    ended_at DATETIME DEFAULT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_status (status)
);

-- 待機中プレイヤーテーブル
CREATE TABLE waiting_players (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_name VARCHAR(32) UNIQUE NOT NULL,
    waiting BOOLEAN NOT NULL DEFAULT TRUE,
    created_at DATETIME NOT NULL
);

-- ゲーム参加プレイヤーテーブル
CREATE TABLE game_players (
    game_id VARCHAR(36) NOT NULL,
    player_id TINYINT NOT NULL,
    user_name VARCHAR(32) NOT NULL,
    score INT DEFAULT 0,
    PRIMARY KEY (game_id, player_id),
    FOREIGN KEY (game_id) REFERENCES games(id),
    UNIQUE KEY unique_game_user (game_id, user_name)
);

-- 式テーブル
CREATE TABLE expressions (
    id VARCHAR(36) PRIMARY KEY,
    game_id VARCHAR(36) NOT NULL,
    player_id TINYINT NOT NULL,
    expression TEXT NOT NULL,
    value TEXT NOT NULL,
    points INT NOT NULL,
    success BOOLEAN NOT NULL,
    submitted_at DATETIME NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games(id),
    INDEX idx_game_id (game_id)
);

-- カードテーブル
CREATE TABLE cards (
    id VARCHAR(36) PRIMARY KEY,
    game_id VARCHAR(36) NOT NULL,
    type ENUM('operand', 'operator', 'item') NOT NULL,
    value VARCHAR(32) NOT NULL,
    owner_player_id TINYINT DEFAULT NULL,
    location ENUM('field', 'hand', 'used') NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games(id),
    INDEX idx_game_id (game_id)
);

-- 手札カード上限テーブル
CREATE TABLE hand_cards_limits (
    game_id VARCHAR(36) NOT NULL,
    player_id TINYINT NOT NULL,
    hand_cards TINYINT NOT NULL,
    PRIMARY KEY (game_id, player_id),
    FOREIGN KEY (game_id) REFERENCES games(id)
);

-- フィールドカード上限テーブル
CREATE TABLE field_cards_limits (
    game_id VARCHAR(36) PRIMARY KEY,
    field_cards TINYINT NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games(id)
);

-- ターンテーブル
CREATE TABLE turns (
    game_id VARCHAR(36) NOT NULL,
    player_id TINYINT NOT NULL,
    turn_number INT NOT NULL,
    start_at DATETIME NOT NULL,
    end_at DATETIME NOT NULL,
    PRIMARY KEY (game_id, turn_number),
    FOREIGN KEY (game_id) REFERENCES games(id)
);