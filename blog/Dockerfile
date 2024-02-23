  FROM alpine
  WORKDIR /libra/blog
  COPY ./target/blog .
  COPY ./config/config-docker.yaml .
  RUN  mkdir config && mv config-docker.yaml config/config.yaml
  EXPOSE 8082 8883
  ENTRYPOINT ["./blog"]