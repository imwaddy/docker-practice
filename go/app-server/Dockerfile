FROM golang:alpine

WORKDIR /src/go/app

COPY . .

RUN go version

RUN go mod download

RUN go build main.go

ENTRYPOINT [ "./main" ]

#EXPOSE 8080