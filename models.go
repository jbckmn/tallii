package main

import (
  "time"
  "labix.org/v2/mgo/bson"
)

// Models
type Account struct {
  Id bson.ObjectId `bson:"_id" json:"_id"`
  Location struct {
    Latitude float32 `json:"latitude"`
    Longitude float32 `json:"longitude"`
  } `json:"location"`
  Username string `json:"username"`
  Gender string `json:"gender"`
  Birthdate time.Time `json:"birthdate"`
  Admin bool `json:"admin"`
  FullAccess bool `bson:"full_access" json:"full_access"`
  Key string `json:"key"`
  Created time.Time `json:"created"`
  Updated time.Time `json:"updated"`
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