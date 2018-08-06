FROM iron/go
EXPOSE 8080

#RUN apk add --update iptables
#RUN iptables -L

WORKDIR /app
ADD routes-service /app/
CMD ["./routes-service"]