FROM golang:1.10.0




# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/minhajuddinkhan/dockergo/
# Add the source code:
WORKDIR $SRC_DIR
ADD . $SRC_DIR

CMD ["go", "run", "cmd/dockergo/main.go"]