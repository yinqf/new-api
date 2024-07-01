docker buildx create --use

docker buildx build --platform linux/arm64,linux/amd64 -t yinqf7437/new-api:2.5.0 . --push

docker buildx build --platform linux/amd64 -t yinqf7437/new-api:2.5.0 . --push
