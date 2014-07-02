package main

import (
  "time"
  "labix.org/v2/mgo/bson"
)

// Models
type Account struct {
  Id bson.ObjectId `bson:"_id"`
  Location struct {
    Latitude float32
    Longitude float32
  }
  Username string
  Gender string
  Birthdate time.Time
  Admin bool
  FullAccess bool `bson:"full_access"`
  Key string
  Created time.Time
  Updated time.Time
}

type Datum struct {
  Id bson.ObjectId `bson:"_id"`
  Quantity float32 `param:"quantity"`
  Category bson.ObjectId
  CategoryName string `bson:"category_name"`
  Account bson.ObjectId
  Created time.Time
  Updated time.Time
}

type Category struct {
  Id bson.ObjectId `bson:"_id"`
  Name string
  data []bson.ObjectId
}