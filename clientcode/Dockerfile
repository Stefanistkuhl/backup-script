FROM alpine:latest
RUN apk add --no-cache rsync --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community
COPY entrypoint.sh /bin/entrypoint.sh
RUN chmod +x /bin/entrypoint.sh
EXPOSE 27017
ENTRYPOINT ["/bin/sh", "/bin/entrypoint.sh"]
