module github.com/speak2jc/k-op

require (
	github.com/NYTimes/gziphandler v1.0.1 // indirect
	github.com/operator-framework/operator-sdk v0.9.1-0.20190802152409-7104d8d7d0e8
	github.com/spf13/pflag v1.0.3
	k8s.io/api v0.0.0-20190808180749-077ce48e77da
	k8s.io/apimachinery v0.0.0-20190808180622-ac5d3b819fc6
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.0.0-20190808180452-d0071a119380
	k8s.io/kubernetes v1.11.8-beta.0.0.20190124204751-3a10094374f2
	k8s.io/sample-controller v0.0.0-20190808182125-a52d0d8c67c5 // indirect
	sigs.k8s.io/controller-runtime v0.1.12
	sigs.k8s.io/controller-tools v0.1.10
)

// Pinned to kubernetes-1.13.4
replace (
	k8s.io/api => k8s.io/api v0.0.0-20190222213804-5cb15d344471
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190228180357-d002e88f6236
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190221213512-86fb29eff628
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190228174230-b40b2a5939e4
)

replace (
	github.com/coreos/prometheus-operator => github.com/coreos/prometheus-operator v0.29.0
	k8s.io/kube-state-metrics => k8s.io/kube-state-metrics v1.6.0
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.1.12
	sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.1.11-0.20190411181648-9d55346c2bde
)

replace github.com/operator-framework/operator-sdk => github.com/operator-framework/operator-sdk v0.9.0
