module github.com/camaoag/cert-manager-webhook-project-pinto

go 1.16

require (
	github.com/jetstack/cert-manager v1.3.1
	github.com/jinzhu/copier v0.3.2
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/oauth2 v0.0.0-20210218202405-ba52d332ba99
	google.golang.org/grpc v1.39.1 // indirect
	k8s.io/api v0.19.0
	k8s.io/apiextensions-apiserver v0.19.0
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
)

// v1.30.0 removes an API which is needed for go.etcd.io/etcd/proxy/grpcproxy, but etcd does not pin its dependency
replace google.golang.org/grpc v1.39.1 => google.golang.org/grpc v1.29.1
