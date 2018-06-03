ARG ARCHREPO
FROM ${ARCHREPO}/golang

ARG QEMU_ARCH
COPY qemu-${QEMU_ARCH}-static /usr/bin/

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]
