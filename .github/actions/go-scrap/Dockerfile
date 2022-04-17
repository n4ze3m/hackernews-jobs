FROM golang:1.18

WORKDIR /go/src/go-scrap

COPY . .

RUN go mod download

RUN go build -o go-scrap

CMD [ "go-scrap" ]