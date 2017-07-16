# Simple REST API

A proof of concept of a simple REST API in Golang.

All data is held in memory, all transfer is via JSON.

All testing can be with __curl__ - although Postman should work too.

## Features

- uses Gorilla MUX (github.com/gorilla/mux)
- returns appropriate HTTP status codes
- modify Person method implemented
- uses __JSON__
 
In __Firefox__ at least, specifying "application/json" allows interpretation of the JSON:

![JSON in Firefox](./json_in_firefox.png)

## Installation

- __Go__ is required (version 1.7 or later)

Fetch this project as follows:

	$ go get -u github.com/mramshaw/Simple-REST-API

This could also be done with __git__ of course:

	$ git clone https://github.com/mramshaw/Simple-REST-API.git

Any dependencies will be fetched in the __build__ process.

## Building

- __make__ is required

This will fetch all dependencies and build the executable:

	$ make

All dependencies will be stored in a local __go__ directory.

## Usage

Simply run the go code (Ctrl-C to terminate):

	$ go run RestfulGorillaMux.go

Or run the executable (Ctrl-C to terminate):

	$ ./RestfulGorillaMux

The API will then be accessible at:

	http://localhost:8100/people

## Testing

Use the following __curl__ commands to test (Postman should work too).

GET (All):

	$ curl -v localhost:8100/people

GET (Individual):

	$ curl -v localhost:8100/people/5

	$ curl -v localhost:8100/people/1

POST (Create):

	$ curl -v -X POST -H "Content-Type: application/json"   \
	       -d '{"firstname":"Tommy","lastname":"Smothers"}' \
	       localhost:8100/people/5

PUT (Update):

	$ curl -v -X PUT -H "Content-Type: application/json" \
	       -d '{"firstname":"Tom","lastname":"Smothers","address":{"city":"Hollywood","state":"CA"}}' \
	       localhost:8100/people/5

DELETE (Delete):

	$ curl -v -X DELETE -H "Content-Type: application/json" \
	       localhost:8100/people/5

[Specifying __-v__ shows the HTTP status codes; this can be omitted if the status codes are not of interest.]

## Credits

Largely based (with some changes) upon this great tutorial by Nic Raboy:

	https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/

There is also a YouTube video (which is linked to from the article).
