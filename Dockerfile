FROM jbonachera/vault
COPY release/vault-envexport /sbin/entrypoint
ENTRYPOINT /sbin/entrypoint
