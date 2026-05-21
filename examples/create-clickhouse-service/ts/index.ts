import * as clickhouse from "@pulumiverse/clickhouse";

new clickhouse.Service("example", {
    region: "us-central1",
    cloudProvider: "gcp",
    password: "1234",
    ipAccesses: [{
        source: "0.0.0.0",
        description: "Test IP"
      }]
});