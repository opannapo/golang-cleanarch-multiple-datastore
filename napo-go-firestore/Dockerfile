FROM golang:1.14.7 AS build-env
ADD . /media/WinLin/GOLANG-PROJECT/napo-go-firestore
WORKDIR /media/WinLin/GOLANG-PROJECT/napo-go-firestore
COPY go.sum .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o /media/WinLin/GOLANG-PROJECT/napo-go-firestore/app-build-output

EXPOSE 8080

ENTRYPOINT ["./app-build-output"]
CMD ["-h"]