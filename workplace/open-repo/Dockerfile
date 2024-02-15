ARG ADDITION_ENV=standard

#### Builder
FROM golang:1.20-alpine as builder

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . /app/

ENV CGO_ENABLED=0

### Test
FROM builder as builder-test
RUN mkdir -p /reports
RUN go test -tags=test -v -covermode=count -coverprofile="/reports/cover.out" ./... 2>&1 |
	tee /reports/report.txt |
	go-junit-report >/reports/report.xml &&
	go tool cover -func="/reports/cover.out"

### Artifacts
FROM scratch as artifacts
COPY --from=builder-test /reports /reports

### Check if there was failures
FROM alpine as builder-test-report
COPY --from=builder-test /reports/report.xml /reports/report.xml
RUN [ $(grep -e "<error" -e "<failure" /reports/report.xml | wc -l) -eq 0 ] || (
	echo "/!\ FAILURES IN UNIT TEST"
	exit 1
)

### Unit Test
FROM builder as unit-test
RUN go test -short ./...

### Race Detector
FROM builder as race-detector
ENV CGO_ENABLED=1
RUN go test -race -short ./...

### Build
FROM builder as builder-build
ARG GIT_COMMIT=dev
ARG VERSION=dev
ARG APP
ARG ADDITION_ENV=standard
RUN echo "✅ Build ${APP} for Linux | with Addition=${ADDITION_ENV}"
RUN GOOS=linux go build \
	-ldflags="-w -s \
	-X https://github.com/open-thienhang-com/Open-Repo.git/libs/utils.BuildDate=$(date +'%Y-%m-%dT%H:%m:%S%:z') \
	-X https://github.com/open-thienhang-com/Open-Repo.git/libs/utils.GitCommit=${GIT_COMMIT} \
	-X https://github.com/open-thienhang-com/Open-Repo.git/libs/utils.Version=${VERSION}" \
	-a -installsuffix cgo \
	-o app cmd/${APP}/main.go

#### Runtime
FROM gcr.io/distroless/base:nonroot-amd64 AS distroless_standard

FROM gcr.io/distroless/base:nonroot-amd64 AS distroless_uploader
COPY --from=mwader/static-ffmpeg:5.1 /ffprobe /usr/local/bin/
COPY --from=mwader/static-ffmpeg:5.1 /ffmpeg /usr/local/bin/

FROM distroless_${ADDITION_ENV}
ARG APP
ENV APP_HTTP_DEBUG=false
EXPOSE 8080

WORKDIR /app
COPY --from=builder-build /app/app /app/cmd/${APP}/*.yml /app/cmd/${APP}/*.env /app/
# COPY --from=builder-build /app/bin/${APP}/migrations /app/migrations

ENTRYPOINT ["/app/app"]

CMD ["server"]
