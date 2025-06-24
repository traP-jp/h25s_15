package db

import (
	"context"
	"fmt"

	"github.com/traP-jp/h25s_15/internal/games/internal/domain"
)

type Ranking struct {
	UserName   string `db:"user_name"`
	Wins       int    `db:"wins"`
	Losses     int    `db:"losses"`
	Draws      int    `db:"draws"`
	TotalScore int    `db:"total_score"`
}

func (r *Repo) GetRanking(ctx context.Context, limit int) ([]domain.RankingItem, error) {
	var ranking []Ranking
	err := r.db.DB(ctx).Select(&ranking, `
		WITH player_stats AS (
			SELECT 
				gp.user_name,
				-- 勝利数: ゲームが終了していて、そのゲームで最高得点を獲得したプレイヤー
				COUNT(CASE 
					WHEN g.status = 'finished' 
						 AND gp.score > (
							 SELECT gp2.score 
							 FROM game_players gp2 
							 WHERE gp2.game_id = gp.game_id AND gp2.player_id != gp.player_id
						 )
					THEN 1 
				END) AS wins,
				-- 負けた数: ゲームが終了していて、そのゲームで相手より得点が低かったプレイヤー
				COUNT(CASE 
					WHEN g.status = 'finished' 
						 AND gp.score < (
							 SELECT gp2.score 
							 FROM game_players gp2 
							 WHERE gp2.game_id = gp.game_id AND gp2.player_id != gp.player_id
						 )
					THEN 1 
				END) AS losses,
				-- 引き分けの数: ゲームが終了していて、そのゲームで相手と同じ得点だったプレイヤー
				COUNT(CASE 
					WHEN g.status = 'finished' 
						 AND gp.score = (
							 SELECT gp2.score 
							 FROM game_players gp2 
							 WHERE gp2.game_id = gp.game_id AND gp2.player_id != gp.player_id
						 )
					THEN 1 
				END) AS draws,
				-- 合計得点
				COALESCE(SUM(gp.score), 0) AS total_score
			FROM game_players gp
			LEFT JOIN games g ON gp.game_id = g.id WHERE g.created_at > "2025-06-22 11:00:00" -- 試遊の部分を除く
			GROUP BY gp.user_name
		)
		SELECT 
			user_name,
			wins,
			losses,
			draws,
			total_score
		FROM player_stats
		WHERE wins > 0 OR losses > 0 OR draws > 0  -- 1回以上プレイしたプレイヤーのみ
		ORDER BY 
			wins DESC,           -- 勝利数が多い順
			total_score DESC     -- 同じ勝利数の場合は合計得点が多い順
		LIMIT ?
	`, limit)
	if err != nil {
		return nil, fmt.Errorf("get ranking: %w", err)
	}

	rankingRes := make([]domain.RankingItem, 0, len(ranking))
	for _, item := range ranking {
		rankingRes = append(rankingRes, domain.RankingItem{
			UserName:   item.UserName,
			Wins:       item.Wins,
			Losses:     item.Losses,
			Draws:      item.Draws,
			TotalScore: item.TotalScore,
		})
	}

	return rankingRes, nil
}
