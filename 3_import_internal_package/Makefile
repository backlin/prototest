clean:
	rm -rf *.pb.go

# -I=..
#	Enables protoc to find the import (2_multiple_files/other.proto).
#
# -I=.
#	Enables protoc to find the input file (test3.proto). -I=. is default if no -I is
#   given, but since since we added -I=.. we have to set it explicitly.
#
# --gogofaster_out=M2_multiple_files/other.proto=github/...
#	Overrides the import path in the output file (test3.pb.go) from "2_multiple_files"
#	to "github.com/.../2_multiple_files".
test3.pb.go: test3.proto
	protoc \
		-I=. \
		-I=.. \
		--gogofaster_out=M2_multiple_files/other.proto=github.com/backlin/prototest/2_multiple_files:. \
		$<
	go build
