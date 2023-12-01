FROM golang:1.21

# Add Maintainer Info
LABEL maintainer="Vitaly Volozhinov <vitaly@simply-vc.com.mt>"

WORKDIR /app

COPY /src/go.mod /src/go.sum ./

RUN go mod tidy

COPY . .
WORKDIR ./src

RUN go build -o main .

EXPOSE 8686

CMD ["./main"]