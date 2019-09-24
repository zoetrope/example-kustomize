#!/bin/bash

TARGET=$1
CMD=${2:-v3}

SCRIPT_DIR=$(cd $(dirname $0); pwd)
KUSTOMIZE_V2=${SCRIPT_DIR}/kustomize2
KUSTOMIZE_V3=${SCRIPT_DIR}/kustomize3
KUSTOMIZE=${KUSTOMIZE_V3}

TARGET_BASE=${SCRIPT_DIR}/example/base


if [ "${CMD}" = "v2" ]; then
    KUSTOMIZE=${KUSTOMIZE_V2}
elif [ "${CMD}" = "diff" ]; then
    echo "=== diff kustomize v2 -> kustomize v3==="
    diff -u <(${KUSTOMIZE_V2} build ${TARGET}) <(${KUSTOMIZE_V3} build ${TARGET})
    exit 0
fi

echo "=== kustomize version ==="
${KUSTOMIZE} version

echo "=== kustomize build ==="
${KUSTOMIZE} build ${TARGET}

echo "=== diff base -> target ==="
diff -u <(${KUSTOMIZE} build ${TARGET_BASE}) <(${KUSTOMIZE} build ${TARGET})
