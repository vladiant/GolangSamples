# Serving static files and web apps in Go

## Basic File Server

```bash
go run main.go
curl localhost:9999/
curl localhost:9999/file2.txt
```

## Serving files on a different route

```bash
curl localhost:9999/static/file2.txt
curl localhost:9999/static/subdir/
```

### Static and dynamic content

```bash
curl localhost:9999/time
curl localhost:9999/static/subdir/
```

## Serving a complete web application

* Go to <http://localhost:9999/>

## Embedding static content into the server's binary

* Go to <http://localhost:9999/>

## Streaming the client code directly from the server

* Go to <http://localhost:9999/>

## Links
* [Serving static files and web apps in Go](https://eli.thegreenplace.net/2022/serving-static-files-and-web-apps-in-go/)
* <https://github.com/eliben/code-for-blog/>
* <https://github.com/eliben/code-for-blog/tree/master/2022/go-static-http-serve-app>
* <https://github.com/eliben/code-for-blog/tree/master/2022/go-static-http-serve-app-embed>
* <https://github.com/eliben/code-for-blog/blob/master/2022/go-static-http-server/server-with-js-fetch.go>