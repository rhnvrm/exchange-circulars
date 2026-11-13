---
category: market-operations
circular_id: 1567b1558063a592
date: '2026-11-13'
description: NSE announces mock trading in Currency Derivatives on November 14, 2026,
  introducing NEAT 3.5.5 and NNF encryption migration guidelines.
draft: false
feed_type: ''
guid: https://nsearchives.nseindia.com/content/circulars/CD77414.pdf
impact: medium
impact_ranking: medium
importance_ranking: medium
justification: Requires trading members to participate in mock trading, adopt NEAT
  3.5.5, and migrate encryption mechanisms for NNF connectivity.
pdf_url: https://nsearchives.nseindia.com/content/circulars/CD77414.pdf
processing:
  attempts: 1
  content_hash: 9d2180adf57f1a17
  processed_at: '2026-11-13T14:05:23.202776'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-11-13T00:00:00+05:30'
rss_url: https://nsearchives.nseindia.com/content/circulars/CD77414.pdf
severity: low
source: nse
stocks: []
tags:
- mock-trading
- currency-derivatives
- neat
- nnf
- trading-operations
- disaster-recovery
title: Mock Trading on Saturday, November 14, 2026 - New Version of NEAT 3.5.5
---

## Summary

National Stock Exchange (NSE) is conducting a mock trading session in the Currency Derivatives Segment on Saturday, November 14, 2026, to test primary and disaster recovery (DR) sites, release NEAT version 3.5.5, and outline NNF encryption migration timelines.

## Key Points

- Mock trading scheduled for Saturday, November 14, 2026, in the Currency Derivatives Segment.
- Trading Session-1 from Primary Site runs from 09:00 hrs to 11:30 hrs; Session-2 from DR Site runs from 12:15 hrs to 14:30 hrs.
- NEAT version 3.5.5 available for download from November 13, 2026, 18:00 hrs at `/cdsftp/cdscommon/NEATCDS355`.
- NEAT version 3.5.4 scheduled for discontinuation on December 05, 2026 (mock).
- NNF members must migrate from the old interactive port (10885) to new encryption mechanisms by December 07, 2026.

## Regulatory Changes

Mandatory adoption of NEAT 3.5.5 and migration to new NNF encryption mechanisms supporting MD5 checksum removal and immediate acknowledgment for order-related messages.

## Compliance Requirements

Trading members must participate in mock trading sessions, upgrade software versions, manually clear outstanding orders in NNF systems, and update gateway router configurations and ports (10889, 10891, 10870, 10871).

## Important Dates

- NEAT 3.5.5 Availability: November 13, 2026, 18:00 hrs
- Mock Trading Date: November 14, 2026
- Version 3.5.4 Discontinuation: December 05, 2026
- Discontinuation of Existing Encryption/Old Port 10885: December 07, 2026

## Impact Assessment

Ensures system readiness, infrastructure resilience via DR switchover testing, and secure NNF connectivity compliance for trading members.