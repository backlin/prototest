# --gogofaster_out=.
#	Tells protoc to generate go files (since it can be used to generate code in many
#	other languages too). gogofaster is a popular alternative to the default go
#	bindings, called using --go_out.
test1.pb.go: test1.proto
	protoc --gogofaster_out=. $<
	go build
