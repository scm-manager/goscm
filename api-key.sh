#!/usr/bin/env bash

set -o errexit
set -o pipefail

username=${1}
keyid=${2}
host="${3:-"ecosystem.cloudogu.com"}"

if [[ -z ${username} ]] || [[ -z ${keyid} ]]; then
    echo "Usage: $(basename "${BASH_SOURCE[0]}") <username> <keyid> [host]"
    exit 1
fi

curl \
    --silent \
    --fail \
    --request POST \
    --user "${username}" \
    --header "Content-Type: application/json" \
    --data "{\"apiKey\":\"${keyid}\"}" \
    "https://${host}/scm/api/v2/cli/login"
