FROM golang as build

COPY . /udpecho/
WORKDIR /udpecho
ENV CGO_ENABLED=0
RUN make

FROM alpine
COPY --from=build /udpecho/udpecho /bin/
CMD /bin/udpecho
