FROM golang:1.15

# Meta data:
LABEL maintainer="matthewgleich@gmail.com"
LABEL description="📣 A user-friendly status outputting library for go"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

CMD ["make", "local-test"]
