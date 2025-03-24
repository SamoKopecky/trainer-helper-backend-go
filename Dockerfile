FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN go build 

FROM golang:1.24 AS runner

WORKDIR /app

RUN apt update -y && apt install -y netcat-traditional

COPY --from=builder /app/trainer-helper trainer-helper
COPY migrations/ migrations/
COPY wait-for.sh wait-for.sh

EXPOSE 2001
CMD ["/app/trainer-helper"]

