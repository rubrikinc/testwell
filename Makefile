.PHONY: all
all: generate test testbuild

.PHONY: generate
generate: ./internal/tests/tests.go ./assert/assert.go ./expect/expect.go

.PHONY: test
test:
	go test -vet off ./internal/tests/

./assert/assert.go: ./internal/assert.tmpl
	rm -f $@
	go generate ./assert/

./expect/expect.go: ./internal/assert.tmpl
	rm -f $@
	go generate ./expect/

./internal/tests/tests.go: ./internal/assert.tmpl
	rm -f $@
	go generate ./internal/tests/

.PHONY: testbuild
testbuild:
	go build ./assert/
	go build ./expect/

.PHONY: clean
clean:
	rm -f \
		./internal/codegen/codegen \
		./internal/tests/tests.go \
		./assert/assert.go \
		./expect/expect.go
