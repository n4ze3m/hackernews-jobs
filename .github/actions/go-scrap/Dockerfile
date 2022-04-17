FROM golang:1.18

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o /scrap


CMD [ "scrap" ]