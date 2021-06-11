FROM golang:1.6.1-alpine
WORKDIR /go/src/app
COPY . .
EXPOSE 8080
ENTRYPOINT ["./skillQuiz"]
CMD go run mainn.go


FROM keinos/sqlite3:latest
COPY init.sql /docker-entrypoint-initdb.d/