package domain


type SubmitScore struct {
	Id       int    `json:"id"`
	GameId   string `json:"game_id"`
	Score    int64  `json:"score"`
}

type UserRanking struct {
	ID        int    `json:"id"`
    UserName  string `json:"user_name"`
	Score     int64  `json:"score"`
	Rank      int64  `json:"rank"`
}