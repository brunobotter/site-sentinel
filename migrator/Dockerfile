FROM postgres:16-alpine
RUN apk add --no-cache bash
COPY run-migrations.sh /run-migrations.sh
RUN chmod +x /run-migrations.sh
ENTRYPOINT ["bash","/run-migrations.sh"]
