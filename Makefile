VERSION=3.1.0
BIN=kustomize

.PHONY: setup clean version $(EXAMPLES)

setup: ${BIN}

${BIN}:
	curl -sSL -o ${BIN} https://github.com/kubernetes-sigs/kustomize/releases/download/v${VERSION}/kustomize_${VERSION}_linux_amd64
	chmod +x ${BIN}

clean:
	rm -f ${BIN}

version: ${BIN}
	./${BIN} version
