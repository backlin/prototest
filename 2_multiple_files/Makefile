clean:
	rm -rf *.pb.go

# $^
#	Lists all proto files as inputs. If only test2.proto is listed then protoc will
#	successfully create test2.pb.go but not other.pb.go, so the type in other.proto will
#	be missing in test2.pb.go, thus it won't compile.
test2.pb.go other.pb.go: test2.proto other.proto
	protoc --gogofaster_out=. $^
	go build