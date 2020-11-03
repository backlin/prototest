clean:
	rm -rf *.pb.go

DEPENDENCY_PATH := ../vendor/google.golang.org/genproto/googleapis/type/date
DEPENDENCY_FILE := $(DEPENDENCY_PATH)/date.proto

setup: $(DEPENDENCY_FILE)

# go modules will not download proto files, so we need to do that ourselves.
$(DEPENDENCY_FILE):
	mkdir -p $(DEPENDENCY_PATH)
	curl \
		"https://raw.githubusercontent.com/googleapis/api-common-protos/master/google/type/date.proto" \
		-o $(DEPENDENCY_FILE)

# -I=../vendor
# -I=.
#	See previous example ../3_import_internal_package/Makefile.
#
# --go_out=.
#	Generate the output using the standard go bindings, not gogofaster used previously.
#	This is necessary because the vendored package is not compatible with gogofaster.
test4.pb.go: test4.proto $(DEPENDENCY_FILE)
	protoc \
		-I=. \
		-I=../vendor \
		--go_out=. \
		$<
	go build
