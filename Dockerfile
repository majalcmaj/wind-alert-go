FROM golang:1.25 AS build
WORKDIR /build

RUN --mount=type=bind,source=.,target=src,rw \
  cd src && \
  make build && \
  mv bin/wind-alert-go ..

# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /build/wind-alert-go ./main
ENTRYPOINT [ "./main" ]
