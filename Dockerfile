FROM golang:1.14-alpine
MAINTAINER Yan-Bin-Lin

# work dir
WORKDIR /creater


# add file to image
ADD app .
# run image
CMD ["./app"]

