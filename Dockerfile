FROM golang:1.8

RUN mkdir -p $GOPATH/src/my-first-beego-project
COPY . $GOPATH/src/my-first-beego-project
WORKDIR $GOPATH/src/my-first-beego-project
ENTRYPOINT sh $GOPATH/src/my-first-beego-project/command/init.sh $GOPATH && $GOPATH/src/my-first-beego-project/my-first-beego-project
#CMD ['./my-first-beego-project']
