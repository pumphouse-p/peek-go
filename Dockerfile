FROM registry.access.redhat.com/ubi8-minimal:8.4-208

EXPOSE 8080

COPY cmd/peek-go /opt/app-root/src

ENTRYPOINT ["/opt/app-root/src/peek-go"]