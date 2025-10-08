package domain


type SubmitScore struct {
	GameId   string `json:"game_id"`
	Score    int64  `json:"score"`
}

type UserRanking struct {
    UserName  string `json:"user_name"`
	Score     int64  `json:"score"`
	Rank      int64  `json:"rank"`
}