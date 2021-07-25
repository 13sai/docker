FROM golang:alpine
RUN mkdir /app
WORKDIR /app
COPY ./hi /app
COPY ./config.yaml /app
RUN chmod +x hi
CMD ["./hi"]