package structs

type Pokemon struct {
	No              int    `json:"no"`
	Name            string `json:"Name"`
	Form            string `json:"form"`
	IsMegaEvolution bool   `json:"isMegaEvolution"`
	Evolutions      []int    `json:"evolutions"`
	Types           []string `json:"types"`
	Abilities       []string `json:"abilities"`
	HiddenAbilities []string `json:"hiddenAbilities"`
	Stats           struct {
		H int `json:"hp"`
		A int `json:"attack"`
		B int `json:"defence"`
		C int `json:"spAttack"`
		D int `json:"spDefence"`
		S int `json:"speed"`
	} `json:"stats"`
}

// POSTのBody変換用
type PostReq struct {
	ID int `json:"id"`
}

// 環境変数取得
type Env struct {
	Port string `envconfig:"PORT" default:"80"` 
}