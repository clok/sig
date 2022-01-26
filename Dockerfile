FROM alpine:3.15.0

COPY sig /usr/local/bin/sig
RUN chmod +x /usr/local/bin/sig

RUN mkdir /workdir
WORKDIR /workdir

ENTRYPOINT [ "/usr/local/bin/sig" ]