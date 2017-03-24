FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/manlycode/get-to-work

ENV USER manlycode
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET c0Z00Q67pqPqssXM

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://manlycode@localhost:5432/get-to-work?sslmode=disable

WORKDIR /go/src/github.com/manlycode/get-to-work

RUN godep go build

EXPOSE 8888
CMD ./get-to-work