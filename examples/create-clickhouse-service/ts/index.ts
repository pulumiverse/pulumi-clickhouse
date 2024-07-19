import * as pulumi from "@pulumi/pulumi";
import * as clickhouse from "@pulumiverse/clickhouse";

new clickhouse.Service("example", {
    region: "us-central1",
    cloudProvider: "gcp",
    tier: "development",
    password: "1234",
    ipAccesses: [{
        source: "0.0.0.0",
        description: "Test IP"
      }]
});