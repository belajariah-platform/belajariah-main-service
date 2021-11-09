FROM golang:1.15.0-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go mod download
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main
RUN go build -o main
RUN go get -d -v

# FROM alpine:latest
# COPY --from=builder /app/main .
# COPY . .


#For small binnary

# Run the binary.
# EXPOSE 3000
CMD ["/app/main"]