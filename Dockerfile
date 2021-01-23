FROM docker:stable

COPY start-clickhouse.sh /start-clickhouse.sh
RUN chmod +x /start-clickhouse.sh

ENTRYPOINT ["/start-clickhouse.sh"]