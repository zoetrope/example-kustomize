V2_VERSION=2.0.3
V3_VERSION=3.1.0
V2_BIN=kustomize2
V3_BIN=kustomize3

.PHONY: setup clean version

setup: $(V2_BIN) $(V3_BIN)

$(V2_BIN):
	curl -sSL -o ${V2_BIN} https://github.com/kubernetes-sigs/kustomize/releases/download/v${V2_VERSION}/kustomize_${V2_VERSION}_linux_amd64
	chmod +x ${V2_BIN}

$(V3_BIN):
	curl -sSL -o ${V3_BIN} https://github.com/kubernetes-sigs/kustomize/releases/download/v${V3_VERSION}/kustomize_${V3_VERSION}_linux_amd64
	chmod +x ${V3_BIN}

clean:
	-rm -f ${V2_BIN} ${V3_BIN}

version:
	./${V2_BIN} version
	./${V3_BIN} version
