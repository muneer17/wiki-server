# Wiki

Build a web server implementing the API described below to run a wiki. In this case, a wiki is basically a key-value storage. You can store any content under a certain name. Later you can retrieve the latest stored content for a particular name.

## API

The server should implement the following REST

### Operation: List stored articles

HTTP Request: GET /articles/

HTTP Response: HTTP/1.1 200 OK Content-Type: application/json (The payload is a JSON array listing the names of all stored articles)

### Operation: Store an article

HTTP Request: PUT /articles/:name

HTTP Response: If a new article was created: HTTP/1.1 201 Created No payload If an existing article was updated: HTTP/1.1 200 OK No payload

### Operation: Read an article

HTTP Request: GET /articles/:name

HTTP Response:If the article is not found: HTTP/1.1 404 Not Found No payload If the article is found: HTTP/1.1 200 OK Content-Type: text/html The payload is the latest content stored under this name

Notes:

- Article name can be any non-empty valid Unicode string
- Article content can be any valid Unicode string
- No articles are persisted between server restarts

## Allowed Technologies

Programming Language: GO or NodeJS

## Additional Acceptance Criteria

- The server can be built using `go build` (or `go run`)
- The solution contains an appropriate level of unit testing that is run using `go test`
- All steps to build and run the server are documented in a README file

When the server is running locally it should be accessible using:

- curl http://localhost:9090/articles/
- curl ‘http://localhost:9090/articles/rest api’
- curl -X PUT http://localhost:9090/articles/wiki -d ‘A wiki is a knowledge base website'

// End of file