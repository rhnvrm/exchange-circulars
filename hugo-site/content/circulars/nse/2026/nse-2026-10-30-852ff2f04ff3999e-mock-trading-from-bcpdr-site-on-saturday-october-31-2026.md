---
category: market-operations
circular_id: 852ff2f04ff3999e
date: '2026-10-30'
description: NSE is conducting a mock trading session in the Currency Derivatives
  segment from Primary and BCP/DR sites on Saturday, October 31, 2026.
draft: false
feed_type: ''
guid: https://nsearchives.nseindia.com/content/circulars/CMPL77214.zip
impact: low
impact_ranking: low
importance_ranking: medium
justification: Routine mock trading session for member system testing and disaster
  recovery switchover preparedness.
pdf_url: https://nsearchives.nseindia.com/content/circulars/CMPL77214.zip
processing:
  attempts: 1
  content_hash: 2b1ca7eac95d9618
  processed_at: '2026-10-30T17:51:49.942534'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-10-30T00:00:00+05:30'
rss_url: https://nsearchives.nseindia.com/content/circulars/CMPL77214.zip
severity: low
source: nse
stocks: []
tags:
- mock-trading
- currency-derivatives
- bcp
- dr-site
- market-operations
title: Mock Trading from BCP/DR site on Saturday, October 31, 2026
---

## Summary

National Stock Exchange of India Limited (NSE) has announced a mock trading session in the Currency Derivatives segment from the Primary and Disaster Recovery (BCP) sites on Saturday, October 31, 2026. The mock session will test system resilience, non-graceful shutdown recovery, and include a re-login session from the primary site.

## Key Points

- **Session Schedule:** Mock trading takes place on Saturday, October 31, 2026, covering Primary site (Trading Session-1: 09:00 to 12:00 hrs), BCP site (Trading Session-2: 12:45 to 17:00 hrs, trade modification end at 17:30 hrs), and a Re-Log in Session from the Primary site (18:30 to 19:00 hrs).
- **Connectivity:** No changes required in NEAT Adapter settings; live session settings of Friday, October 30, 2026, will be retained.
- **Order Management:** All outstanding orders will be purged before transitioning to the BCP site. NNF software users must manually clear session 1 orders.
- **Operational Details:** Non-graceful shutdown from the primary site will occur during the mock session. Multicast TBT sequence numbers will restart from '1' upon switching to the BCP site.
- **Compliance & Rules:** No financial obligations or fund pay-ins/pay-outs will result from trades executed during mock sessions. Only valid and approved UCC/PAN are permitted.

## Regulatory Changes

No direct regulatory policy changes; conducted pursuant to standard exchange mock trading guidelines and SEBI requirements.

## Compliance Requirements

Members must ensure systems are prepared for non-graceful shutdown scenarios, clear outstanding orders between sessions, and verify valid UCC/PAN details.

## Important Dates

- **Mock Trading Date:** October 31, 2026
- **Circular Date:** October 30, 2026

## Impact Assessment

Minimal operational impact on live trading days. Allows trading members to test connectivity, failover handling, and disaster recovery preparedness without financial risk.