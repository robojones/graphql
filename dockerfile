FROM golang as builder

WORKDIR /go/src/github.com/steebchen/graphql

RUN curl -sL https://deb.nodesource.com/setup_10.x | bash - && apt-get install -y nodejs
RUN npm i -g prisma

# install godep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only

COPY prisma/ ./prisma/

RUN prisma generate

COPY api/*.graphqls api/**/*.graphqls ./api/
COPY gqlgen/ ./gqlgen/

RUN ls

RUN go run gqlgen/cmd.go -c gqlgen/gqlgen.yml

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /main .

FROM scratch
COPY --from=builder /main /main
CMD ["/main"]
