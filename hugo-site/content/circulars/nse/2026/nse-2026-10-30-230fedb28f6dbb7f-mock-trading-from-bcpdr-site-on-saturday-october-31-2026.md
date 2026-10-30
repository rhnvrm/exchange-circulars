---
category: market-operations
circular_id: 230fedb28f6dbb7f
date: '2026-10-30'
description: NSE is conducting a mock trading session from its Primary and BCP/DR
  sites for the Currency Derivatives segment on Saturday, October 31, 2026.
draft: false
feed_type: ''
guid: https://nsearchives.nseindia.com/content/circulars/CD77197.pdf
impact: low
impact_ranking: low
importance_ranking: low
justification: Routine mock trading session to test BCP/DR site readiness and ensure
  seamless Monday operations.
pdf_url: https://nsearchives.nseindia.com/content/circulars/CD77197.pdf
processing:
  attempts: 1
  content_hash: 9ea0ecb529957f0b
  processed_at: '2026-10-30T17:54:16.240773'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-10-30T00:00:00+05:30'
rss_url: https://nsearchives.nseindia.com/content/circulars/CD77197.pdf
severity: low
source: nse
stocks: []
tags:
- mock-trading
- currency-derivatives
- bcp
- disaster-recovery
- neat-adapter
title: Mock trading from BCP/DR site on Saturday, October 31, 2026
---

## Summary

National Stock Exchange of India Limited (NSE) has announced a mock trading session for the Currency Derivatives segment from the Primary and BCP/DR site on Saturday, October 31, 2026. The session includes primary site trading, BCP site trading with a non-graceful shutdown switchover, and a subsequent Re-Log in session from the Primary site.

## Key Points

- Mock trading covers both Primary and BCP/DR sites for Currency Derivatives.
- Schedule involves Session-1 (Primary Site: 09:00 - 12:00 hrs) and Session-2 (BCP Site: 12:45 - 17:00 hrs, trade modification end at 17:30 hrs).
- Live Re-log in session from Primary site takes place between 18:30 - 19:00 hrs to prevent login issues on Monday, November 02, 2026.
- Switchover from primary to DR site will be a non-graceful shutdown.
- All outstanding orders will be purged before trading starts from the BCP site.
- Multicast TBT sequence number will restart from '1' at the BCP site.

## Regulatory Changes

No direct regulatory changes; conducted as part of mandatory periodic business continuity and disaster recovery testing.

## Compliance Requirements

- Members must ensure NEAT Adapter version 1.0.29 and NEAT 3.5.4 are utilized as per guidelines.
- Members using NNF software must manually clear outstanding orders of Session 1 before BCP site trading begins.
- Only valid and compliant UCC/PAN uploaded and approved before cutoff will be allowed.
- Actions in case of non-graceful shutdown should follow circular NSE/MSD/48662 dated June 18, 2021.

## Important Dates

- Mock Trading Date: October 31, 2026
- Normal Live Trading resumption: November 02, 2026

## Impact Assessment

Minimal operational impact. Allows trading members to test disaster recovery switchover readiness, verify connectivity parameters, and ensure smooth system operations.