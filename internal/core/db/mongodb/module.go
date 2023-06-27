package mongodb

import (
	"go.uber.org/fx"
)

var MongoDBModule = fx.Module("MongoDBModule",
	fx.Provide(func() (ClientWrapper, CollectionWrapper) {
		client := NewMongoClient("mongodb+srv://connextor:ConnextorPassword1@connextor.lfncyii.mongodb.net/?retryWrites=true&w=majority")
		clientWrap := NewClientWrapper(client)
		db := client.Database("saver-service-dev")
		coll := db.Collection("pockets")
		collWrap := Collection{Collection: coll}
		return clientWrap, collWrap
	}),
)
