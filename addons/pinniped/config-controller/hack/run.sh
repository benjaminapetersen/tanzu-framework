#!/bin/bash

# Copyright 2022 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

RED='\033[0;31m'
RESET='\033[0m'

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Always run from config-controller directory for reproducibility
cd "${MY_DIR}/.."

tag="dev"
# tag="$(uuidgen)" # Uncomment to create random image every time
image="harbor-repo.vmware.com/tkgiam/$(whoami)/pinniped-config-controller-manager:$tag"

# be run to run ./generate.sh first to get current RBAC
RBAC_FILE=./config/rbac/role.yaml
if [ ! -f "${RBAC_FILE}" ]
then
    echo -e "${RED}without RBAC the controller will not run. please run ./hack/generate.sh${RESET}"
    exit 1
else
    kubectl apply -f "${RBAC_FILE}"
fi

docker build -t "$image" .
docker push "$image"
ytt --data-value "image=$image" -f deployment.yaml | kbld -f - | kapp deploy -a pinniped-config-controller-manager -f - -y
kubectl logs -n pinniped deploy/pinniped-config-controller-manager -f
