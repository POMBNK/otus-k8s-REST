FROM golang:1.22-alpine as builder

WORKDIR /usr/local/src

COPY ./ ./migrations

FROM alpine AS runner
RUN apk --no-cache add bash make git go

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY --from=builder /usr/local/src/migrations /migrations
RUN cd root/go/bin/ && mv ./goose ../../../

CMD ./goose -dir migrations postgres postgres://postgres:postgres@postgres:5432/postgres up