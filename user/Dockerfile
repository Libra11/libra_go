  FROM alpine
  WORKDIR /Initial
  COPY ./target/user .
  COPY ../config/user/config-docker.yaml .
  RUN  mkdir config && mv config-docker.yaml config/config.yaml
  EXPOSE 8080 8881
  ENTRYPOINT ["./user"]
