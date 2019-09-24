module github.com/zchee/spinctl

require (
	cloud.google.com/go v0.36.0
	contrib.go.opencensus.io/exporter/ocagent v0.4.6 // indirect
	contrib.go.opencensus.io/exporter/stackdriver v0.9.1
	github.com/antihax/optional v0.0.0-20180407024304-ca021399b1a6
	github.com/aws/aws-sdk-go v1.17.2 // indirect
	github.com/census-instrumentation/opencensus-proto v0.1.0 // indirect
	github.com/evanphx/json-patch v4.1.0+incompatible // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/google/go-cmp v0.2.1-0.20181115012043-2248b49eaa8e
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf // indirect
	github.com/google/pprof v0.0.0-20190208070709-b421f19a5c07 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/json-iterator/go v0.0.0-20181112064556-d05f387f50c0
	github.com/minio/sha256-simd v0.0.0-20190131020904-2d45a736cd16
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/sergi/go-diff v1.0.1-0.20180205163309-da645544ed44
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/zchee/go-open v1.0.0
	github.com/zchee/go-xdgbasedir v1.0.3
	go.opencensus.io v0.22.1
	go.uber.org/atomic v1.3.3-0.20181018215023-8dc6146f7569 // indirect
	go.uber.org/multierr v1.1.1-0.20180122172545-ddea229ff1df // indirect
	go.uber.org/zap v1.9.2-0.20190122184550-82e3b47c1fe4
	golang.org/x/build v0.0.0-20190221024721-9e83d8587038 // indirect
	golang.org/x/crypto v0.0.0-20190219172222-a4c6cb3142f2 // indirect
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd
	golang.org/x/oauth2 v0.0.0-20190220154721-9b3c75971fc9
	golang.org/x/sys v0.0.0-20190220154126-629670e5acc5 // indirect
	google.golang.org/api v0.1.0
	google.golang.org/genproto v0.0.0-20190219182410-082222b4a5c5 // indirect
	google.golang.org/grpc v1.18.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190202010724-74b699b93c15 // indirect
	k8s.io/apimachinery v0.0.0-20190117220443-572dfc7bdfcb // indirect
	k8s.io/cli-runtime v0.0.0-20190202014047-491c94071cfa
	k8s.io/client-go v2.0.0-alpha.0.0.20190202011228-6e4752048fde+incompatible
	k8s.io/klog v0.2.0 // indirect
	sigs.k8s.io/yaml v1.1.0
)

replace gopkg.in/DataDog/dd-trace-go.v1 v1.9.0 => github.com/DataDog/dd-trace-go v1.9.0
