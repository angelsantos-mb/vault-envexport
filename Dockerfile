FROM jbonachera/alpine
COPY release/vault-envexport /sbin/entrypoint
ENTRYPOINT /sbin/entrypoint
