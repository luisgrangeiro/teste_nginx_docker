FROM golang:latest

COPY . .

RUN go build main.go

EXPOSE 9000

ENTRYPOINT [ "./main" ]