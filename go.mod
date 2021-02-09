module github.com/coredns/coredns

go 1.13

require (
	github.com/Azure/azure-sdk-for-go v40.6.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.13
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.5
	github.com/Infoblox-CTO/atlas.onprem.authn v1.0.0 // indirect
	github.com/aws/aws-sdk-go v1.36.2
	github.com/caddyserver/caddy v1.0.5
	github.com/cenkalti/backoff/v4 v4.1.0
	github.com/coredns/caddy v1.1.0
	github.com/dnstap/golang-dnstap v0.3.0
	github.com/farsightsec/golang-framestream v0.3.0
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/infobloxopen/go-trees v0.0.0-20190313150506-2af4e13f9062
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/miekg/dns v1.1.35
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/prometheus/client_golang v1.8.0
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.14.0
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200306183522-221f0cc107cb
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	golang.org/x/sys v0.0.0-20201015000850-e3ed0017c211
	google.golang.org/api v0.29.0
	google.golang.org/grpc v1.29.1
	gopkg.in/DataDog/dd-trace-go.v1 v1.27.1
	k8s.io/api v0.19.4
	k8s.io/apimachinery v0.19.4
	k8s.io/client-go v11.0.1-0.20191029005444-8e4128053008+incompatible
	k8s.io/klog v1.0.0
)
