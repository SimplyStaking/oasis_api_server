FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Vitaly Volozhinov <vitaly@simply-vc.com.mt>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]