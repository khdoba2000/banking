FROM golang:1.18 as builder

#
RUN mkdir -p $GOPATH/src/github.com/khdoba2000/banking
WORKDIR $GOPATH/src/github.com/khdoba2000/banking

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/banking /

FROM alpine
COPY --from=builder banking .
COPY .env ./.env
COPY ./configs/rbac_model.conf ./configs/rbac_model.conf
COPY ./configs/models.csv ./configs/models.csv
COPY ./db/migrations/ ./db/migrations/

ENTRYPOINT ["./banking"]
