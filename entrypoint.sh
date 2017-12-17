#!/bin/bash

VAULT_ADDR=${VAULT_ADDR:-"https://identity.cloud.vx-labs.net"}
SECRETS_PATH=${SECRETS_PATH:-"secrets/ci"}
APPROLE_ID=${APPROLE_ID}
APPROLE_SECRET=${APPROLE_SECRET}

echo '{"role_id":"'$APPROLE_ID'", "secret_id":"'$APPROLE_SECRET'"}'  | vault write -format=json auth/approle/login @/dev/stdin
