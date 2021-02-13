FROM scratch
COPY echo-server /echo-server
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/echo-server"]
