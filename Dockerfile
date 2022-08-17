FROM golang:1.18.4

RUN apt-get update && apt-get upgrade -y \
  && apt-get install -y ca-certificates \
  && apt-get clean

# Add build-level env vars (these typically come from CICD pipeline)
ARG COMMIT_HASH
ENV COMMIT_HASH=${COMMIT_HASH}

ARG COMMIT
ENV COMMIT=${COMMIT}

# $GOPATH
ARG DIR=$GOPATH/src/demo/
WORKDIR $DIR

# Get dependancies - will also be cached if we won't change mod/sum
COPY go.mod $DIR
COPY go.sum $DIR
RUN go mod download

COPY . $DIR

# install requirements
RUN go get -d -v

# compile to executable
RUN go build -o demo

EXPOSE $PORT

CMD ["./demo"]
