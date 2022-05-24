FROM golang:1.18.2

ENV REPO_URL=github.com/rajesh4b8/users-api-batch-2
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL

ENV WORKPATH=$APP_PATH

WORKDIR $WORKPATH

COPY ./src ./src
COPY ./go.mod ./
COPY ./go.sum ./

RUN go build -o users-api ./src

EXPOSE 8080

CMD [ "./users-api" ]