#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

# builder image
FROM golang:1.20.1-alpine3.17 as builder

RUN mkdir /build
COPY cmd /build/cmd
COPY vendor /build/vendor
COPY go.mod /build/go.mod
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o aegis-init-container ./cmd/main.go

# generate clean, final image for end users
# FROM alpine:3.17
FROM gcr.io/distroless/static-debian11
COPY --from=builder /build/aegis-init-container .

# executable
ENTRYPOINT [ "./aegis-init-container" ]
CMD [ "" ]