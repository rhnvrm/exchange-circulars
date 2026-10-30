---
category: market-operations
circular_id: 57e916dffda88849
date: '2026-10-30'
description: NSE is conducting a mock trading session for the Commodity Derivatives
  segment from Primary and BCP/DR sites on Saturday, October 31, 2026.
draft: false
feed_type: ''
guid: https://nsearchives.nseindia.com/content/circulars/COM77198.pdf
impact: low
impact_ranking: low
importance_ranking: low
justification: Routine mock trading session for system readiness and BCP switchover
  testing.
pdf_url: https://nsearchives.nseindia.com/content/circulars/COM77198.pdf
processing:
  attempts: 1
  content_hash: 733dbce92734c9c8
  processed_at: '2026-10-30T17:53:57.881116'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-10-30T00:00:00+05:30'
rss_url: https://nsearchives.nseindia.com/content/circulars/COM77198.pdf
severity: low
source: nse
stocks: []
tags:
- mock-trading
- market-operations
- commodity-derivatives
- disaster-recovery
title: Mock trading from BCP/DR site on Saturday, October 31, 2026
---

## Summary

National Stock Exchange of India Limited (NSE) is conducting a mock trading session in the Commodity Derivatives segment from its Primary and BCP/DR sites on Saturday, October 31, 2026. The mock session will be followed by a live re-login session from the Primary site to prevent login issues on Monday, November 02, 2026.

## Key Points

- Primary Site Mock Session (Session 1): Normal Market open at 09:00 hrs and close at 12:00 hrs.
- BCP Site Mock Session (Session 2): Normal Market open at 12:45 hrs and close at 17:00 hrs; Position Limit/Collateral value setup and trade modification end time is 17:15 hrs.
- Re-Login Session: Live Re-login start time at 18:30 hrs and close time at 19:00 hrs from Primary Site.
- Operations will shift from primary site to DR site via a non-graceful shutdown during the mock session.
- All outstanding orders will be purged before trading starts from the BCP site.
- Multicast TBT sequence numbers will restart from '1' upon switching to the BCP site.
- Trades during mock sessions carry no financial obligations or fund pay-in/pay-out.

## Regulatory Changes

No regulatory changes; conducted as part of routine Exchange disaster recovery and business continuity planning.

## Compliance Requirements

- Members must ensure NEAT Adapter settings remain as of Friday, October 30, 2026.
- Members using NNF software must manually clear session 1 outstanding orders before trading commences from the BCP site.
- Take necessary actions in case of non-graceful shutdown as outlined in circular NSE/MSD/48662.

## Important Dates

- Mock Trading Date: October 31, 2026
- Live Operations Resume: November 02, 2026

## Impact Assessment

Ensures member systems and trading infrastructure are robust and capable of handling disaster recovery site switchovers smoothly without disrupting regular market hours.