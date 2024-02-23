  chcp 65001
  cd user
  docker build -t user:latest .
  cd ..
  cd blog
  docker build -t blog:latest .
  cd ..
  cd api
  docker build -t api:latest .
  cd ..
  docker-compose up -d