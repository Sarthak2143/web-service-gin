# web-service-gin

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

RESTful API written in Go with Gin.

API returns data about vintage jazz records.

## Endpoints

- ``/albums``

    - ``GET`` – Get a list of all albums, returned as JSON.
    - ``POST`` – Add a new album from request data sent as JSON.

- ``/albums/:id``

    - ``GET`` – Get an album by its ID, returning the album data as JSON.


## Prerequisites

- **An installation of Go 1.16 or later**. For installation instructions, see [Installing Go.](https://go.dev/doc/install)

- **A tool to edit your code**. Any text editor you have will work fine.

- **A command terminal**. Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows.

- **The curl tool**. On Linux and Mac, this should already be installed. On Windows, it’s included on Windows 10 Insider build 17063 and later. For earlier Windows versions, you might need to install it. For more, see [Tar and Curl Come to Windows.](https://docs.microsoft.com/en-us/virtualization/community/team-blog/2017/20171219-tar-and-curl-come-to-windows)

## Setup

```bash
git clone https://github.com/Sarthak2143/web-service-gin
cd web-service-gin/
go get .
go run .
```
It started the web server on `localhost:8080`.

## Examples

- Getting all albums

```bash
curl http://localhost:8080
```

```json
[
    {
        "id": "1",
        "title": "After Dark",
        "artist": "Mr. Kitty",
        "Price": 56.99
    },
    {
        "id": "2",
        "title": "Where is my mind?",
        "artist": "Pixies",
        "Price": 59.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "Price": 48.99
    }
]
```

- Adding an album

```bash
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Wed, 02 Jun 2021 00:34:12 GMT
Content-Length: 116
```
```json

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}
```

- Getting an album by `ID`

```bash
curl http://localhost:8080/albums/2
```

```json
{
    "id": "2",
    "title": "Where is my mind?",
    "artist": "Pixies",
    "Price": 59.99
}
```

---
