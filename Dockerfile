# Docker file quy định nội dung của 1 docker image - dựa theo Dockerfile, Docker biết cần làm những gì để create 1 docker image
#stage build
#Build image base on base image
FROM golang:1.19-alpine as builder

#Copy all files from project to images
COPY ./ /app/

#Set working directory is folder app
WORKDIR /app/

#Cài đặt các dependencies cho project
RUN go mod download
#Buid project
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-airbnb .

#stage runner
FROM alpine
WORKDIR /app/
# copy file thực thi go từ stage trước đó
COPY --from=builder /app/go-airbnb .
#copy config
COPY config/config.yaml ./config/
#copy migrations
COPY migrations migrations
#CMD ["make migrate_up"]
CMD ["/app/go-airbnb"]