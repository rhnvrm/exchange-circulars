---
category: trading
circular_id: 8fe15274ac3b1695
date: '2026-11-25'
description: MCX introduces separate Product IDs and Product Names for Call and Put
  option contracts (such as OFGOLDMC and OFGOLDMP for GOLDM) effective November 26,
  2026.
draft: false
feed_type: ''
guid: https://www.mcxindia.com/docs/default-source/circulars/english/2026/november/circular-664_2026.pdf
impact: medium
impact_ranking: medium
importance_ranking: medium
justification: Requires trading members to update front-end systems, order management,
  and cancellation logic for split Call and Put product IDs.
pdf_url: https://www.mcxindia.com/docs/default-source/circulars/english/2026/november/circular-664_2026.pdf
processing:
  attempts: 1
  content_hash: 609f0ee76fc7d53f
  processed_at: '2026-11-26T12:08:07.452394'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-11-24T00:00:00+05:30'
rss_url: https://www.mcxindia.com/docs/default-source/circulars/english/2026/november/circular-664_2026.pdf
severity: medium
source: mcx
stocks: []
tags:
- mcx
- options
- modification-in-the-contract
- commodity-derivatives
- trading
title: 'MCX Exchange Circular (TECH): Introduction of Separate Product IDs and Product
  Name for Call and Put Options'
---

## Summary

Multi Commodity Exchange of India Limited (MCX) announced the segregation of composite option product IDs into separate Product IDs and Product Names for Call and Put options, effective November 26, 2026, in continuation of Circular No: MCXTECH/261/2026.

## Key Points

- **Product ID Segregation**: Option contracts with common Product IDs are separated into distinct Call (C) and Put (P) Product IDs and names.
- **GOLDM Options Bifurcation**: The existing product `OFGOLDM` is split into `OFGOLDMC` (Call Options) and `OFGOLDMP` (Put Options).
- **Mass Order Cancellation**: Mass order cancellation requests must be submitted using the respective split Product ID.
- **GTC/GTD Orders**: All active GTC/GTD orders associated with affected contracts will be cancelled upon the introduction of split Product IDs.
- **Unaffected Parameters**: No changes to Symbol names, Bhav Copy, Trade Files, RT feeds, or API structures.

## Regulatory Changes

- Implementation of distinct product identifiers for Call and Put options to enhance operational efficiency.

## Compliance Requirements

- Members must ensure readiness of front-end systems and make necessary configurations to support split Product IDs for order management and cancellation processes.

## Important Dates

- **Circular Date**: November 25, 2026
- **Effective Date**: November 26, 2026

## Impact Assessment

Trading members must update order routing, cancellation workflows, and system configurations to seamlessly handle separate Call and Put product identifiers.