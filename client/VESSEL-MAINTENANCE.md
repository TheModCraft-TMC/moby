# Vessel split-SDK compatibility

Based on upstream `client/v0.6.3`, commit
`2ca8f919108134b1d6e79ed95b1c55ffc32f9c2c`.

`ServerVersionResult` retains the four top-level engine diagnostics already decoded
by the SDK: GitCommit, GoVersion, KernelVersion and BuildTime. They remain separate
from Components, preserving older daemon responses and modern responses containing
both forms. Existing result fields, response hooks, decoding/error behavior and
transport/negotiation are unchanged. No daemon packages are added to this module.

The complete client module race suite passes on Go 1.26.11. A real HTTP regression
covers componentless API 1.24 and coexistence with modern component metadata.
Vessel can now remove its secondary body-observation adapter rather than modifying
or bypassing the newer SDK's intentional hook-body restrictions.
