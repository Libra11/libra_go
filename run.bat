  chcp 65001
  cd user
  docker build -t user:latest .
  cd ..
  docker-compose up -d