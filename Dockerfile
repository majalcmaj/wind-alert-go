FROM golang:1.25 AS build
WORKDIR /build

RUN --mount=type=bind,source=.,target=src,rw \
  cd src && \
  make build && \
  ls && \
  mv bin/lambda ..
  # go build -tags lambda.norpc -o ../main main.go

# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /build/lambda ./lambda
ENTRYPOINT [ "./lambda" ]
