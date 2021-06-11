FROM golang
WORKDIR /go/src/app
COPY . .
ENTRYPOINT ["./skillQuiz"]
CMD go run mainn.go