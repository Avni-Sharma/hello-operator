FROM registry.access.redhat.com/ubi7/ubi-minimal:7.6

ENV OPERATOR=/usr/local/bin/hello-operator \
    USER_UID=1001 \
    USER_NAME=hello-operator

# install operator binary
COPY build/_output/bin/hello-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
