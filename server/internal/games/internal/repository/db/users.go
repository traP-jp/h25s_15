package db

import "context"

func (r *Repo) GetUsersCount(ctx context.Context) (int, error) {
	var count int
	err := r.db.DB(ctx).Get(&count, `
		WITH users as  (
			SELECT user_name FROM game_players AS gp
			LEFT JOIN games AS g ON gp.game_id = g.id
			WHERE g.created_at > "2025-06-22 11:00:00"
			GROUP BY user_name 
		) 
		SELECT COUNT(user_name) FROM users
	`)
	if err != nil {
		return 0, err
	}
	return count, nil
}
