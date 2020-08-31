FROM golang:1.15.0-alpine3.12
COPY . /app/service
WORKDIR /app/service
EXPOSE 6007
CMD [ "go", "run", "main.go"]