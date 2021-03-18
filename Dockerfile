FROM golang:1.14.15

# Add Maintainer Info
LABEL maintainer="Vitaly Volozhinov <vitaly@simply-vc.com.mt>"

WORKDIR /app

COPY /src/go.mod /src/go.sum ./

RUN go mod download

COPY . .
WORKDIR ./src

RUN go build -o main .

EXPOSE 8686

CMD ["./main"]