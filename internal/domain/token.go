package domain

type Session struct {
	Token string `bson:"token"`
	Exp   int64  `bson:"exp"`
}
