clean:
	rm -rf *.pb.go

FILE := date.proto
REL_PATH := google/type
VENDOR_PATH := ../vendor/google.golang.org/genproto/googleapis

$(VENDOR_PATH)/$(REL_PATH)/$(FILE):
	mkdir -p $(VENDOR_PATH)/$(REL_PATH)
	curl \
		"https://raw.githubusercontent.com/googleapis/api-common-protos/master/$(REL_PATH)/$(FILE)" \
		-o $(VENDOR_PATH)/$(REL_PATH)/$(FILE)

# Downloads the dependency to the repo.
test4_vendored.pb.go: test4_vendored.proto $(VENDOR_PATH)/$(REL_PATH)/$(FILE)
	protoc \
		-I=. \
		-I=$(VENDOR_PATH) \
		--go_out=. \
		$<
	go build
