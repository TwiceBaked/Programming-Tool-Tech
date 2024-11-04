# Use the official Golang image as the base image
FROM golang:1.23

WORKDIR /app

COPY go.mod .
COPY gws.go .
COPY hello-j.json .
COPY help.html .
COPY syllabus.json .

RUN go build -o bin .
EXPOSE 9000

ENTRYPOINT [ "/app/bin" ]