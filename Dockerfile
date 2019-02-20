FROM alpine:3.9

ENV PORT 80

EXPOSE 80

WORKDIR /

ADD web /

ADD ui/dist /public

CMD ["/web"]
