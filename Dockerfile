FROM golang:1.25
WORKDIR /usr/src/app
# copy all the files to the container
COPY . .
RUN apt update \
    && apt install -y imagemagick 
CMD ["go","test"]