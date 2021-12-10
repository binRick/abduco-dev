clear >/dev/null 2>&1
(
  go run . | jq -Mrc
) >&1
