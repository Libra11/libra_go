  FROM alpine
  WORKDIR /libra/user
  COPY ./target/user .
  COPY ./config/config-docker.yaml .
  RUN  mkdir config && mv config-docker.yaml config/config.yaml
  EXPOSE 8080 8881
  ENTRYPOINT ["./user"]
