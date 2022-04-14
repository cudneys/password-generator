FROM scratch
COPY dist/linux/amd64/server /
COPY docker/etc/ /etc

USER appuser
ENTRYPOINT ["/server"]