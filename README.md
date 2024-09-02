# Opentelemetry-Openseacrh integration example

This is a "simple" example of an integration between opentelemetry and opensearch. The ideia behind this example is to:

- use Otel SDK and API to generate trace data
- export trace data to Otel Collector
- export data from Otel Collector to Opensearch Dataprepper
- export data from Dataprepper to Opensearch
- Vizualize Trace Data with Opensearch Dashboards