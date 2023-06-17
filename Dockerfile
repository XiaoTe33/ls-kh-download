FROM busybox

EXPOSE 3306

COPY ./ls-kh-download /
COPY ./conf.yaml /etc/

ENTRYPOINT ["/ls-kh-download"]