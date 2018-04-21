FROM scratch
COPY bin/gnats-proxy /gnats-proxy
CMD [ "/gnats-proxy", "$@" ]
