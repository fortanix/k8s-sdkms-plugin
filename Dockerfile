FROM alpine:latest
RUN apk --no-cache add ca-certificates && update-ca-certificates

COPY build/k8s-sdkms-plugin /bin/k8s-sdkms-plugin
ENTRYPOINT ["/bin/k8s-sdkms-plugin"]
