FROM gocv/opencv:4.5.3

RUN go get -u -d gocv.io/x/gocv

COPY main.go go.* /src/

WORKDIR /src/

CMD go run .

