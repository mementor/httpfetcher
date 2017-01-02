
# -----------------------------------------------------------------------------
# docker-httpfetcher
#
# Builds a basic docker image that can run httpfetcher
#
# Authors: Roman Kubar
# Updated: 02.01.2017
# Require: Docker (http://www.docker.io/)
# -----------------------------------------------------------------------------

# Base system is Ubuntu 16.04
FROM   easypi/alpine-arm

RUN apk add ca-certificates

ADD httpfetcher /bin/httpfetcher

VOLUME ["/data"]
ENTRYPOINT    ["/bin/httpfetcher"]
