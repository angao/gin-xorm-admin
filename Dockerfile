# Build the manager binary
FROM golang:1.10.3 as builder

# Copy in the go src
WORKDIR /go/src/github.com/angao/gin-xorm-admin
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o gin-xorm-admin github.com/angao/gin-xorm-admin/

# Copy the gin-xorm-admin into a empty image
FROM scratch
WORKDIR /
COPY --from=builder /go/src/github.com/angao/gin-xorm-admin/ .
ENTRYPOINT ["./gin-xorm-admin"]
