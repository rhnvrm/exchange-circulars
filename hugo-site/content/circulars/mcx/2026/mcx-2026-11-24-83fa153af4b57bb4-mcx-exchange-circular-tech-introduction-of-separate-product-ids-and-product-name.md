---
category: trading
circular_id: 83fa153af4b57bb4
date: '2026-11-25'
description: MCX is introducing separate Product IDs and Product Names for Call and
  Put option contracts (such as GOLDM/SILVERM) effective November 26, 2026.
draft: false
feed_type: ''
guid: https://www.mcxindia.com/docs/default-source/circulars/english/2026/november/circulars-664-2026.pdf
impact: medium
impact_ranking: medium
importance_ranking: medium
justification: Segregation of option contracts into separate Call and Put Product
  IDs affects member trading systems, mass order cancellation, and GTC/GTD orders.
pdf_url: https://www.mcxindia.com/docs/default-source/circulars/english/2026/november/circulars-664-2026.pdf
processing:
  attempts: 1
  content_hash: 4cbe04a05a1114de
  processed_at: '2026-11-26T01:47:59.653245'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-11-24T00:00:00+05:30'
rss_url: https://www.mcxindia.com/docs/default-source/circulars/english/2026/november/circulars-664-2026.pdf
severity: medium
source: mcx
stocks: []
tags:
- options
- derivatives
- modification-in-the-contract
- product-launch
- mcx
title: 'MCX Exchange Circular (TECH): Introduction of Separate Product IDs and Product
  Names for Call and Put Options'
---

## Summary

Multi Commodity Exchange of India Limited (MCX) has announced the introduction of separate Product IDs and Product Names for Call (C) and Put (P) option contracts of specified commodities, effective November 26, 2026. This change aims to enhance operational efficiency by segregating symbol OFGOLDM into separate products (OFGOLDMC for Call Options and OFGOLDMP for Put Options for SILVERM/GOLDM option contracts).

## Key Points

- **Product ID Classification**: Option contracts previously assigned a common Product ID will now have separate Product IDs and Product Names for Call and Put options.
- **Bifurcation Example**: For SILVERM/GOLDM option contracts, the existing product OFGOLDM will be bifurcated into OFGOLDMC (Call Options) and OFGOLDMP (Put Options).
- **Mass Order Cancellation**: Mass order cancellation requests must be submitted using the respective Product ID separately for Call and Put option orders.
- **Order Expiry**: Upon introduction of split Product IDs, all active GTC/GTD orders associated with the affected contract will be cancelled.
- **Unaffected Parameters**: No changes are introduced to Symbol names, API structures, member-facing files (Bhav Copy, Trade Files, RT feeds), or market data feeds (MDI, EMDI, EOBI).

## Regulatory Changes

- Segregation of commodity option symbols into distinct Call and Put product identifiers (OFGOLDMC and OFGOLDMP).
- Mandatory requirement for members to submit separate cancellation requests for Call and Put option orders.

## Compliance Requirements

- Members and vendors must ensure readiness of their front-end systems and configure them to support handling split Product IDs for order management and related processes.
- Ensure appropriate Product IDs are used when submitting Mass Order Cancellation requests.

## Important Dates

- **Circular Date**: November 25, 2026 (Circular No: MCXTECH/664/2026)
- **Effective Date**: November 26, 2026

## Impact Assessment

- **Operational Impact**: Medium. Members must update their front-end trading systems to handle split Product IDs for options contracts.
- **Order Impact**: All active GTC/GTD orders on affected contracts will be cancelled automatically upon product splitting.