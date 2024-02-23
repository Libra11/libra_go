  chcp 65001
  cd user
  docker build -t user:latest .
  docker tag user:latest libra001/user:latest
  cd ..
  cd blog
  docker build -t blog:latest .
  docker tag blog:latest libra001/blog:latest
  cd ..
  cd api
  docker build -t api:latest .
  docker tag api:latest libra001/api:latest
  cd ..
  docker push libra001/user:latest
  docker push libra001/blog:latest
  docker push libra001/api:latest