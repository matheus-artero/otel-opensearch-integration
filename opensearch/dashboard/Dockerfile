FROM opensearchproject/opensearch-dashboards:2.9.0
RUN /usr/share/opensearch-dashboards/bin/opensearch-dashboards-plugin remove securityDashboards
COPY --chown=opensearch-dashboards:opensearch-dashboards opensearch_dashboards.yaml /usr/share/opensearch-dashboards/config/opensearch_dashboards.yml