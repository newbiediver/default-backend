FROM alpine
ENV GOTRACEBACK=crash
RUN apk add --no-cache tzdata
WORKDIR /app
RUN mkdir -p /app/web
COPY "./web/*" "/app/web/"
ADD [ "default-backend", "/app/" ]
ENTRYPOINT [ "/app/default-backend" ]