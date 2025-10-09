package domain


type SubmitScore struct {
	Id       int    `db:"id"      json:"id"`
	UserId   int    `db:"user_id" json:"user_id"`
	GameId   string `db:"game_id" json:"game_id"`
	Score    int64  `db:"score"   json:"score"`
}

type UserRanking struct {
	ID        int    `json:"id"`
    UserName  string `json:"user_name"`
	Score     int64  `json:"score"`
	Rank      int64  `json:"rank"`
}