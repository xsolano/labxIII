##################################
# Build small image without code #
##################################

FROM alpine:3.10
WORKDIR /webhook

# copy local "service" binary to docker.
COPY service .

EXPOSE 8080
CMD ["sh", "/webhook/service"]