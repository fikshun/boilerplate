FROM golang:1.22.0-alpine
WORKDIR /boilerplate
COPY . /boilerplate
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /boilerplate
EXPOSE 8080
CMD ["/boilerplate"]
