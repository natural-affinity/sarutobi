include github.com/natural-affinity/makefiles/golang.bin.mk
PACKR := $(value GOPATH)\bin\go-bindata.exe

$(BIN): $(SRC) $(PACKR)
	go-bindata -o wisdom\wisdom.go -pkg wisdom wisdom
	go build -o $(BIN)

$(PACKR):
	go get -u github.com/tj/go-bindata/...
