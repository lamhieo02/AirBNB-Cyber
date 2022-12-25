# Docker file quy định nội dung của 1 docker image - dựa theo Dockerfile, Docker biết cần làm những gì để create 1 docker image
#Build image base on base image
FROM golang:1.19-alpine as builder
#Run: dùng để thực thi một câu lệnh trong images
RUN mkdir /app
#Copy all files from project to images
COPY ./ /app/
#Set working directory
WORKDIR /app/
#Cài đặt các dependencies cho project
RUN go mod download
#Buid project
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-airbnb .

FROM alpine
WORKDIR /app/

COPY --from=builder /app/go-airbnb /app/config ./

CMD ["/app/go-airbnb"]