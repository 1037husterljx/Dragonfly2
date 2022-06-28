#!/bin/bash

set -o nounset
set -o xtrace
set -o errexit
set -o pipefail

D7Y_VERSION=${D7Y_VERSION:-"latest"}
D7Y_REGISTRY=${D7Y_REGISTRY:-reg.docker.alibaba-inc.com/dragonfly2}
curDir=$(cd "$(dirname "$0")" && pwd)
cd "${curDir}/../" || return

docker-push() {
    docker push "${D7Y_REGISTRY}"/"${1}":"${D7Y_VERSION}"
}

main() {
    case "${1-}" in
    cdn)
        docker-push cdn
        ;;
    dfdaemon)
        docker-push dfdaemon
        ;;
    scheduler)
        docker-push scheduler
        ;;
    manager)
        docker-push manager
    esac
}

main "$@"