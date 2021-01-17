FROM golang:1.15

WORKDIR /go/src/github.com/Omelman/task-management/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/task-management github.com/Omelman/task-management/cmd/task-management

###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/task-management /usr/local/bin/task-management
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["task-management"]