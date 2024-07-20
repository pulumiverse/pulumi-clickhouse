"""A Python Pulumi program"""

import pulumiverse_clickhouse as clickhouse

service_account = clickhouse.Service(
    "example",
    region = "us-central1",
    cloud_provider = "gcp",
    tier = "development",
    password =  "1234",
    ip_accesses = [{
        "source": "0.0.0.0",
        "description": "Test IP"
      }]
)
