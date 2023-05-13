package models

/*Tweet estructura del tweet*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
