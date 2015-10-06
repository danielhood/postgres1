# postgres1

CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -installsuffix netgo -ldflags '-w' .
