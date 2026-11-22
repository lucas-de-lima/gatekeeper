param(
  [string]$RunId = "gatekeeper-onboarding-$([DateTime]::UtcNow.ToString('yyyyMMddTHHmmssZ'))"
)

$ErrorActionPreference = "Stop"
$root = (Resolve-Path (Join-Path $PSScriptRoot "..\..")).Path
$runDir = Join-Path $root ".harness\runs\$RunId"
New-Item -ItemType Directory -Force -Path $runDir | Out-Null

Set-Location $root
$commit = (& git rev-parse HEAD 2>$null).Trim()
$status = (& git --no-pager status --porcelain=v1 2>$null)
$remote = (& git remote get-url origin 2>$null).Trim()

$docs = Get-ChildItem -LiteralPath (Join-Path $root "docs") -File -ErrorAction SilentlyContinue
$docNames = @($docs | ForEach-Object { $_.Name })
$blueprintFiles = @(Get-ChildItem -Recurse -LiteralPath (Join-Path $root ".harness\blueprint") -File)

$checks = [ordered]@{
  git_root = $root
  commit = $commit
  remote = $remote
  untracked_or_modified_count = @($status).Count
  target_project_manifest = Test-Path (Join-Path $root ".harness\project-manifest.yaml")
  blueprint_copy_present = $blueprintFiles.Count -gt 0
  blueprint_file_count = $blueprintFiles.Count
  brainstorming_doc_count = $docNames.Count
  brainstorming_documents = $docNames
  executable_source_present = ((Test-Path (Join-Path $root "go.mod")) -or
                              (Test-Path (Join-Path $root "package.json")) -or
                              (Test-Path (Join-Path $root "build.gradle.kts")) -or
                              (Test-Path (Join-Path $root "Cargo.toml")) -or
                              (Test-Path (Join-Path $root "pom.xml")))
  status = "waiting_for_request"
  result = "completed"
  implementation_started = $false
}

$run = [ordered]@{
  schema_version = 1
  run_id = $RunId
  workspace_kind = "managed_target"
  project_id = "gatekeeper"
  work_mode = "target_discovery"
  status = "waiting_for_request"
  started_at = [DateTime]::UtcNow.ToString("yyyy-MM-ddTHH:mm:ssZ")
  objective = "Validate target-local bootstrap and preserve brainstorming-only status."
  checks = $checks
}
$runJson = $run | ConvertTo-Json -Depth 8
[IO.File]::WriteAllText((Join-Path $runDir "run.json"), $runJson)

$event = [ordered]@{
  occurred_at = [DateTime]::UtcNow.ToString("yyyy-MM-ddTHH:mm:ssZ")
  event_type = "target_discovery_completed"
  run_id = $RunId
  result = "completed"
  payload = $checks
}
[IO.File]::WriteAllText((Join-Path $runDir "events.jsonl"), ($event | ConvertTo-Json -Compress) + [Environment]::NewLine)

$evidenceFiles = @("run.json","events.jsonl")
$evidenceIndex = [ordered]@{ run_id=$RunId; evidence=@() }
foreach($name in $evidenceFiles){
  $path=Join-Path $runDir $name
  $hash=(Get-FileHash -Algorithm SHA256 -LiteralPath $path).Hash.ToLower()
  $evidenceIndex.evidence += [pscustomobject]@{ location=$name; sha256=$hash; result="observed" }
}
[IO.File]::WriteAllText((Join-Path $runDir "evidence-index.json"), ($evidenceIndex | ConvertTo-Json -Depth 8))

$handoff = @"
# Gatekeeper Target Discovery Handoff

Status: waiting_for_request

## Observed
- Blueprint copy present: $($checks.blueprint_copy_present)
- Blueprint files: $($checks.blueprint_file_count)
- Brainstorming documents: $($docNames.Count)
- Executable source present: $($checks.executable_source_present)
- Implementation started: false

## Evidence
- run.json
- events.jsonl
- evidence-index.json

## Next
- Confirm MVP scope and architecture decisions.
- Receive an explicit implementation request.
"@
[IO.File]::WriteAllText((Join-Path $runDir "handoff.md"), $handoff)

Write-Output "Gatekeeper target discovery completed: $runDir"
