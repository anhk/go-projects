FROM golang as build

COPY . /httpecho/
WORKDIR /httpecho
ENV CGO_ENABLED=0
RUN make

FROM alpine
COPY --from=build /httpecho/httpecho /bin/
CMD /bin/httpecho
