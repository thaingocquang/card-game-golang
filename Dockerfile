# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /app

COPY . .

RUN ls -a

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -buildvcs=false -o /buildfile

EXPOSE 8080

EXPOSE $PORT

CMD [ "/buildfile" ]