include github.com/natural-affinity/makefiles/golang.bin.mk

PACKR := $(value GOPATH)\bin\go-bindata.exe
DBDIR := data
SRC := $(SRC) $(wildcard $(DBDIR)/*.*)
DB := bindata.go

# build when changed including embedding static content
$(BIN): $(SRC) $(PACKR)
	go-bindata -ignore=.go -o $(DB) $(DBDIR)
	go build -o $(BIN)

# fetch static embed tool
$(PACKR):
	go get -u github.com/natural-affinity/go-bindata/...

# remove application and intermediary files
clean:
	@rm -f $(DB)
	@go clean -i
