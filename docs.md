# TrueReview backend Firestore Wrapper

This application provides an API wrapper around firestore to manage authentication and the database itself. 

## API Calls

There are a number of supported opperations all connecting directly to firestore.

### GET Call

http://truereview.network/api/movies/title_here

Response body if found 

```json
{
    "message": "Movie has been found",
    "movies": {
        "date": "02/04/1999",
        "score": "88",
        "title": "AuthMovie"
    },
    "status": true
}
```

If movie has not been found

```json
{
    "message": "Movie not found",
    "status": false
}
```

#### To get all movies

http://truereview.network/api/movies/all

```json
{
    "message": "Movies have been found",
    "movies": [
        {
            "date": "02/04/1999",
            "score": "88",
            "title": "AuthMovie"
        },
        {
            "date": "02/04/1999",
            "score": "88",
            "title": "MovieToDelete"
        },
        {
            "date": "02/04/1999",
            "score": "88",
            "title": "MovieToDeleteTwo"
        },
        {
            "date": "02/04/1999",
            "score": "88",
            "title": "PostMovie"
        }
    ],
    "status": true
}
```

### POST Call

http://truereview.network/api/movies/p

Unfortunately, there has to be some kind of string after '/movies/'

Posting must have a body of type 'application/json' and a body that matches the move:

```json
{
	"title" : "MovieToDeleteTwo",
	"score" : "88",
	"date" : "02/04/1999"
}
```

A successful post will be responded with

```json
{
    "message": "Movie has been created and stored",
    "movie": {
        "title": "MovieToDeleteTwo",
        "score": "88",
        "date": "02/04/1999"
    },
    "status": true
}
```

### DELETE Call

http://truereview.network/api/movies/title_here

A DELETE will always respond with a successful message

```json
{
    "message": "Movie has been deleted",
    "status": true
}
```

## Deployment

Deployed on a Kubernetes cluster managed by GKE

### Building

```
make build
```

### Running 

```
make run
```

### Deploying

```
make upload
```

### Calling

http://truereview.network/api/movies/