package main

var Routes = []Route {
  Route {
    "Index",
    "GET",
    "/",
    Index,
  },
  Route {
    "TMP",
    "GET",
    "/tmp",
    Tmp,
  },
}
