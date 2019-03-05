FROM golang as builder

WORKDIR /app

RUN curl -sL https://deb.nodesource.com/setup_10.x | bash - && apt-get install -y nodejs
RUN npm i -g prisma

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY prisma/ ./prisma/

RUN prisma generate

COPY api/*.graphqls api/**/*.graphqls ./api/
COPY gqlgen/ ./gqlgen/

RUN go run gqlgen/cmd.go -c gqlgen/gqlgen.yml

COPY . ./

RUN go run github.com/google/wire/cmd/wire

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /main .

FROM scratch
COPY --from=builder /main /main
CMD ["/main"]
