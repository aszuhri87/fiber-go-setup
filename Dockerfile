FROM golang:1.22.2-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

EXPOSE 8070

RUN go mod tidy

RUN go build 

# CMD [ "go", "run", "main.go", "--host=0.0.0.0"] 
