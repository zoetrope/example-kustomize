BIN_DIR = bin
PLUGIN_DIR = plugins
PLUGIN_SRC = $(shell find $(PLUGIN_DIR) -type f -name "*.go" -not -name "*_test.go")

V2_VERSION=2.0.3
V3_VERSION=3.1.0

.PHONY: setup build clean distclean version

setup:
	# GOPATH=${PWD} go get -u sigs.k8s.io/kustomize/v3/cmd/kustomize
	# curl -sSL -o ${BIN_DIR}/kustomize2 https://github.com/kubernetes-sigs/kustomize/releases/download/v${V2_VERSION}/kustomize_${V2_VERSION}_linux_amd64
	curl -sSL -o ${BIN_DIR}/kustomize3 https://github.com/kubernetes-sigs/kustomize/releases/download/v${V3_VERSION}/kustomize_${V3_VERSION}_linux_amd64
	# chmod +x ${BIN_DIR}/kustomize2
	chmod +x ${BIN_DIR}/kustomize3

build: $(PLUGIN_SRC)
	cd $(PLUGIN_DIR)/kustomize/plugin/myapiversion/mykind && go build -buildmode plugin -o mykind.so mygenerator.go

clean:
	:

distclean:
	# -rm -f ${BIN_DIR}/kustomize2
	-rm -f ${BIN_DIR}/kustomize3

version:
	# ${BIN_DIR}/kustomize2 version
	${BIN_DIR}/kustomize3 version
