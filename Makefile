include github.com/natural-affinity/makefiles/golang.bin.mk

DBDIR := data
DBBIN := bindata.go
PACKR := $(value GOPATH)\bin\go-bindata.exe
SRC := $(SRC) $(wildcard $(DBDIR)/*.*)

# build when changed including embedding static content
$(BIN): $(SRC) $(PACKR)
	go-bindata -ignore=.go -o $(DBBIN) $(DBDIR)
	go build -o $(BIN)

# fetch static embed tool
$(PACKR):
	go get -u github.com/natural-affinity/go-bindata/...

# remove application and intermediary files
clean:
	@rm -f $(DBBIN)
	@go clean -i
