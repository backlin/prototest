clean:
	rm -rf *.pb.go

# -I=$$GOPATH/src/googleapis
#	Imports the dependncy from outside the repo. You must clone it manually for this to
#	work.
#
# -I=.
#	See previous example ../3_import_internal_package/Makefile.
#
# --go_out=.
#	Generate the output using the standard go bindings, not gogofaster used previously.
#	This is necessary because the vendored package is not compatible with gogofaster.
test4.pb.go: test4.proto
	protoc \
		-I=. \
		-I=$$GOPATH/src/googleapis \
		--go_out=. \
		$<
	go build
