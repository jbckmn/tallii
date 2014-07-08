package main

import (
  "fmt"
  "net/http"

  "github.com/zenazn/goji"
  "github.com/zenazn/goji/web"
  "github.com/unrolled/render"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

func main() {

  goji.Get("/hello/:name", hello)
  goji.Get("/accounts/:id", viewAccount)
  goji.Put("/accounts/:id", updateAccount)
  goji.Post("/data", makeDatum)
  goji.Get("/data/:id", viewDatum)
  goji.Put("/data/:id", upDatum)
  goji.Delete("/data/:id", deleteDatum)
  goji.Get("/categories/:id", viewCategory)
  goji.Get("/categories", allCategories)
  goji.Get("/regenerate", regenerateKey)
  goji.Get("/", apiInfo)
  goji.Serve()
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func viewAccount(c web.C, w http.ResponseWriter, r *http.Request) {
  rend := render.New(render.Options{})
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    fmt.Println("No session")
    panic(err)
  }
  defer session.Close()
  colA := session.DB("trackme").C("accounts")
  if colA != nil {
    fmt.Println("Got an account collection")
  }
  result := Account{}
  err = colA.Find(bson.M{"key": c.URLParams["id"]}).One(&result)
  if err != nil {
    fmt.Println("No account")
    panic(err)
  }
  if &result != nil {
    colD := session.DB("trackme").C("data")
    lotsData := []Datum{}
    iter := colD.Find(bson.M{"account": result.Id}).Limit(100).Iter()
    err = iter.All(&lotsData)
    if err != nil {
      panic(err)
    }
    // jsonResp := make(map[string][]Datum)
    jsonAcct := make(map[string]*Account)
    jsonAcct["account"] = &result
    // jsonResp['data'] = &lotsData
    rend.JSON(w, http.StatusOK, jsonAcct)
  }
}