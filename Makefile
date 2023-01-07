
# add some tests
# add default shell
# add build and install scripts
# add docker build and install (run in docker always)

test:
	@go run cmd/eddie/main.go

testctl:
	@go run cmd/eddiectl/main.go
