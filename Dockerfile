FROM golang:1.18
WORKDIR /src
COPY . .
RUN go build -o action .