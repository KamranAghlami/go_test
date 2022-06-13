FROM golang

WORKDIR /usr/src/go_test

COPY . .

RUN go build main.go

CMD ["./main"]
