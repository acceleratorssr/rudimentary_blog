# 镜像基于:
FROM ubuntu:20.04

COPY blog /service/blog
COPY conf.yaml /service/conf.yaml
COPY log /service/log

WORKDIR /service

#CMD ["/service/blog"]
ENTRYPOINT ["/service/blog"]