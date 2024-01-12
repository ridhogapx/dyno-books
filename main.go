package main

import "dyno-books/helper"

var user1  = helper.Data{
  UserEmail: "john@mail.com",
  Book: []helper.Book{
    {
      Title: "Attack Of Titan",
      Author: "Hajime",
    },
    {
      Title: "Evangelion",
      Author: "Anon",
    },
  },
}

var user2  = helper.Data{
  UserEmail: "neko@mail.com",
  Book: []helper.Book{
    {
      Title: "Sword Art Online",
      Author: "Budi",
    },
    {
      Title: "Blend W",
      Author: "Supri",
    },
    {
      Title: "Dokter Stone",
      Author: "Agus",
    },
  },
}

func main() {

  // Add some data -> helper.AddNote(&user1)
  // Add some data -> helper.AddNote(&user2)

  // Find the data by "user_email"
  helper.FindBook("john@mail.com")
}
