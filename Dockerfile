FROM golang:1.25 as build-dist

RUN git config --global credential.helper store
COPY  ./.git-credentials /root/.git-credentials 
RUN mkdir -p /xm-workspace/NEW-SHADOW-SERVER/ 
WORKDIR  /xm-workspace/NEW-SHADOW-SERVER
COPY . .
RUN rm -f gin-init  &&  mkdir  gin-init  && mkdir -p gin-init/conf && cp -r conf/* ./gin-init/conf 
RUN go build -o ./gin-init/gin-init ./
ENTRYPOINT ["./gin-init"]

FROM debian:11
RUN mkdir -p /xm-workspace/NEW-SHADOW-SERVER/ 
WORKDIR  /xm-workspace/NEW-SHADOW-SERVER
COPY --from=build-dist --chown=root:root /xm-workspace/NEW-SHADOW-SERVER/gin-init/  /xm-workspace/NEW-SHADOW-SERVER/ 

ENTRYPOINT ["./gin-init"]