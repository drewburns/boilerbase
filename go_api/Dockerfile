FROM golang:alpine AS go-build-env
ADD . /src
RUN cd /src && go build ./cmd/my_api
RUN cd /src && go build ./cmd/migrations

FROM alpine
COPY --from=go-build-env /src/my_api /app/
COPY --from=go-build-env /src/migrations /app/
ENTRYPOINT /app/my_api
