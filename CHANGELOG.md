CHANGELOG
=========

## HEAD (Unreleased)

- Update upstream ClickHouse Terraform provider from v3.11.1 to v3.18.0
- Add new resources: `Clickpipe`, `ClickpipeCdcInfrastructure` (ClickPipes reverse private
  endpoints), `OrganizationSettings`, `PostgresService`, `Role`, `RoleAssignment`,
  `ServiceScheduledScaling`, `ServiceUpgradeWindow`
- Add `ComputeID` for `clickhouse_clickpipe_cdc_infrastructure` (no `id` attribute upstream)
- Require Go >= 1.26 (required by upstream)
- Regenerate all SDKs (Go, Node.js, Python, .NET)
- Pin pulumi CLI to 3.237.0 in CI workflows (SDK codegen runs through
  `pulumi package gen-sdk`, so the CLI version determines generated output)

---

## v2.0.4 (2026-05-22)

Release-only tag to publish v2 language SDKs and provider artifacts under a fresh version
(avoid re-upload conflicts from partial v2.0.2/v2.0.3 release runs).

- No provider or schema changes since v2.0.2

---

## v2.0.2 (2026-05-21)

- Update upstream ClickHouse Terraform provider from v1.0.0 to v3.11.1
- Update pulumi-terraform-bridge from v3.89.0 to v3.125.0
- Migrate pf bridge imports to unified v3 module path
- Update golangci-lint configuration to v2 format
- Add new resources: `ServiceTransparentDataEncryptionKeyAssociation`, `ApiKey.GetId`
- Add `ComputeID` for resources without an `id` attribute
- Require Go >= 1.25
- Move Go SDK module to `sdk/v2/go/clickhouse`
- Regenerate all SDKs (Go, Node.js, Python, .NET)

---
