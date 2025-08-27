module github.com/ONSdigital/dp-zebedee-sdk-go

go 1.24

require (
	github.com/ONSdigital/dp-net/v2 v2.22.0
	github.com/smartystreets/goconvey v1.8.1
)

require (
	github.com/ONSdigital/dp-api-clients-go/v2 v2.266.0 // indirect
	github.com/ONSdigital/dp-net/v3 v3.2.0 // indirect
	github.com/ONSdigital/log.go v1.1.0 // indirect
	github.com/ONSdigital/log.go/v2 v2.4.6 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/hokaccha/go-prettyjson v0.0.0-20211117102719-0474bc63780f // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/justinas/alice v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/smarty/assertions v1.16.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
)

// Remove vulnarability for [CVE-2024-45337] CWE-863: Incorrect Authorization
tool golang.org/x/crypto
