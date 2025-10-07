FROM golang:1.25 AS build
WORKDIR /build

RUN --mount=type=bind,source=.,target=src,rw \
  cd src && \
  make build && \
  mv bin/lambda ..

# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /build/lambda ./main
ENTRYPOINT [ "./main" ]
