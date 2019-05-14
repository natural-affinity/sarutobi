include github.com/natural-affinity/makefiles/golang.bin.mk
PACKR := $(value GOPATH)\bin\go-bindata.exe
DBDIR := wisdom
DB := $(DBDIR)\$(DBDIR)_generated.go

# build when changed including embedding static content
$(BIN): $(SRC) $(PACKR)
	go-bindata -o $(DB) -pkg $(DBDIR) $(DBDIR)
	go build -o $(BIN)

# fetch static embed tool
$(PACKR):
	go get -u github.com/tj/go-bindata/...

# remove application and intermediary files
clean:
	@rm -f $(DB)
	@go clean -i
