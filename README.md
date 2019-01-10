<p align="center"><img src="https://blog.golang.org/gopher/gopher.png"></p>

# booksapi
Golang microservice sample

This is a sample of microservice with basic HTTP methods built with amazing Golang and awesome PostgreSQL database

To run it on your machine, follow those steps:
* Make sure you have PostgreSQL 9.4 or greater and run the `migrations/201901060312_migration.sql` to create the `library` database and the schema;
* On terminal, go to `src/config` dir and execute: `cp config/config.json.example config/config.json` to create a `config.json` file with all necessary settings;
* To run this example, also make sure you have downloaded all Golang libs:

```
go get github.com/gorilla/mux
go get github.com/lib/pq
```

* Open your `config.json` file and make sure your database settings are placed to access your server and database that you should already be created;

This example implements the basics HTTP verbs to provide the CRUD operations.

* To run it, on terminal, go to `src` directory and execute:

`go run *.go`

If everything went go, you should see the following message:

`waiting routes on port :8081...`

If you wish to change the port or you already have this one in use, you can change it on your `config.json` file on `service.port` setting item;

To see it in action, you can make HTTP requests using curl or any other tool you like. The following example uses curl on terminal;
* To insert a book, you should make an HTTP POST request like:

```
curl --request POST \
  --url http://localhost:8081 \
  --data '{"title": "book1", "ibsn": "0123456789123", "about": "about book1" "edition": 1,	"page_number": 177", "release_date": "2019-01-11",	"author_id": 1}'
```

  The result will be the id crated on boook table and HTTP status code 201 (created);

  * Let's assume the returned id was 1, to update this book, you should make an HTTP PUT request like:

```
curl --request PUT \
  --url http://localhost:8081/1 \
  --data '{"title": "new name book1", "ibsn": "0123456789123", "about": "about book1" "edition": 1,	"page_number": 177", "release_date": "2019-01-11",	"author_id": 1}'
```

  You can update any attribute you want, except the `id`, that one is our primary key. The result will be HTTP status code 200;

  * To remove and book, you should make an HTTP DELETE request like:

```
curl --request DELETE \
  --url http://localhost:8081/1
```

  No data param is required, only que id on URL schema. The result should be a 200 HTTP status code;

  * If you wish the get an specific book, you shoul make an HTTP GET request passing the `id` on URL, like:

```
curl --request GET \
  --url http://localhost:8081/1
  ```

  The result will be an book object and a 200 HTTP status code;

  * Let's say you want to navigate through the books, like paging data. To do that, you can make a GET HTTP request passing `X-PageSize` and `X-CurrentPage` headers. Pay attention on `X-CurrentPage` header, it starts on 0 index:

  ```
  curl --request GET \
  --url http://localhost:8081 \
  --header 'X-CurrentPage: 0' \
  --header 'X-PageSize: 2'
  ```

  Wrapping all up, the HTTP verbs that this example implements have the following schema:

  ```
  POST http://localhost
  PUT http://localhost/{id}
  DELETE http://localhost/{id}
  GET http://localhost/{id}
  GET http://localhost
  ```

  I hope you enjoy all of this. If you have any questions and any suggestions, please send me a feed back on: washington.moises@gmail.com, I'll be glad to hear from you.

  Thanks!