FROM golang:1.16-alpine

#COPY go.mod .
COPY go.sum .
#RUN go mod download

COPY *.go .

RUN go build main.go

EXPOSE 8080

CMD [ "./main" ]