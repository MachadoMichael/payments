package schema

type Store struct {
	Id           string `json:"id"`
	SocialReason string `json:"socialReason"`
	FantasyName  string `json:"fantasyName"`
	Bank         string `json:"bank"`
}
