FROM alpine
RUN apk -U add ca-certificates && \
    rm -rf /var/cache/apk/
COPY release/vault-envexport /sbin/entrypoint
ENTRYPOINT /sbin/entrypoint
