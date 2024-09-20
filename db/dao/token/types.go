package token

type Token struct {
	AccessToken string
	AccessUUid  string
	AtExp       int64
}

type RefreshToken struct {
	RefreshToken string
	RefreshUUid  string
	ReExp        int64
}
