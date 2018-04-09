# my-kali i
FROM my-kali

WORKDIR /app

ENV GOPATH=/opt/go

ENV SRC_DIR="$GOPATH/src/github.com/emnetag/server"

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/

ENTRYPOINT ["./myapp"]