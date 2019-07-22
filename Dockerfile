FROM node:10.16-alpine as build-node

WORKDIR $GOPATH/src/md-wiki

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM golang:1.12-alpine3.10 as build-go

WORKDIR $GOPATH/src/md-wiki

COPY . .

RUN go get -d -v ./...

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM nginx:1.16.0-alpine

USER root

COPY --from=build-node $GOPATH/src/md-wiki/dist /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf

COPY nginx/nginx.conf /etc/nginx/conf.d

RUN echo "daemon off;" >> /etc/nginx/nginx.conf

WORKDIR /go/src/md-wiki

COPY --from=build-go /go/src/md-wiki .

EXPOSE 8080

CMD ./main & nginx