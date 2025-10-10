package domain


type SubmitScore struct {
	UserId   int    `db:"user_id"   json:"user_id"`
	UserName string `db:"user_name" json:"user_name"`
	GameId   string `db:"game_id"   json:"game_id"`
	Score    int64  `db:"score"     json:"score"`
}

type UserRanking struct {
    UserName  string   `json:"user_name"`
	Score     float64  `json:"score"`
	Rank      int64    `json:"rank"`
}

