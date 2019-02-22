FROM golang as builder

WORKDIR /go/src/github.com/steebchen/graphql

# install godep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /main .

FROM scratch
COPY --from=builder /main /main
CMD ["/main"]
