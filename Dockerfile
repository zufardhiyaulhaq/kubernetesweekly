#################
# Base image
#################
FROM alpine:3.12 as kubernetesweekly-base

USER root

RUN addgroup -g 10001 kubernetesweekly && \
    adduser --disabled-password --system --gecos "" --home "/home/kubernetesweekly" --shell "/sbin/nologin" --uid 10001 kubernetesweekly && \
    mkdir -p "/home/kubernetesweekly" && \
    chown kubernetesweekly:0 /home/kubernetesweekly && \
    chmod g=u /home/kubernetesweekly && \
    chmod g=u /etc/passwd

ENV USER=kubernetesweekly
USER 10001
WORKDIR /home/kubernetesweekly

#################
# Builder image
#################
FROM golang:1.15-alpine AS kubernetesweekly-builder
RUN apk add --update --no-cache alpine-sdk
WORKDIR /app
COPY . .
RUN make build

#################
# Final image
#################
FROM kubernetesweekly-base

COPY --from=kubernetesweekly-builder /app/bin/kubernetesweekly /usr/local/bin

# Command to run the executable
ENTRYPOINT ["./app/bin/kubernetesweekly"]
