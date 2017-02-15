VERSION="X.Y.Z"

all:
	go build -ldflags "-X main.version=$(VERSION)"