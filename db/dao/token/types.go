package token

type Token struct {
	AccessToken  string
	RefreshToken string
	AccessUUid   string
	RefreshUUid  string
	AtExp        int64
	ReExp        int64
}
type AccessDetails struct {
	AccessUid string
	Userid    int64
}
