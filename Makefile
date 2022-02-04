test:
	go test -count=1 -race -cover ./...
.PHONY: test

audit:
	go list -json -m all | nancy sleuth
.PHONY: audit

build:
	go build ./...
.PHONY: build

lint:
	#./ci/scripts/doGoLint.sh
	./ci/scripts/doGoFmt.sh
	#exit
.PHONY: lint