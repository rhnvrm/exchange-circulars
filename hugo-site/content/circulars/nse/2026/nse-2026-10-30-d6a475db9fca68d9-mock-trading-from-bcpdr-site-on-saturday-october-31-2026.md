---
category: market-operations
circular_id: d6a475db9fca68d9
date: '2026-10-30'
description: NSE is conducting a mock trading session from Primary and BCP/DR sites
  on Saturday, October 31, 2026, followed by a live re-login session.
draft: false
feed_type: ''
guid: https://nsearchives.nseindia.com/content/circulars/CMTR77196.pdf
impact: low
impact_ranking: low
importance_ranking: low
justification: Routine mock trading session from BCP/DR site to test exchange system
  resilience and member readiness.
pdf_url: https://nsearchives.nseindia.com/content/circulars/CMTR77196.pdf
processing:
  attempts: 1
  content_hash: a0a5855eb903f9c2
  processed_at: '2026-10-30T17:54:34.839671'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-10-30T00:00:00+05:30'
rss_url: https://nsearchives.nseindia.com/content/circulars/CMTR77196.pdf
severity: low
source: nse
stocks: []
tags:
- mock-trading
- capital-market
- disaster-recovery
- bcp
- special-trading
- live-activities-schedule
title: Mock trading from BCP/DR site on Saturday, October 31, 2026
---

## Summary

National Stock Exchange of India Limited (NSE) will conduct a mock trading session on Saturday, October 31, 2026, across both the Primary and Disaster Recovery (DR)/Business Continuity Plan (BCP) sites. The session will test non-graceful switchovers, order purging, multicast TBT sequence resets, and will conclude with a re-login session to ensure smooth operations for Monday, November 02, 2026.

## Key Points

- **Mock Date:** Saturday, October 31, 2026.
- **Session Structure:** Trading Session-1 from Primary Site followed by Trading Session-2 from DR Site (including pre-open, normal market, T+0 market, call auction illiquid session, auction market, and closing sessions).
- **Site Switchover:** The exchange will execute a non-graceful shutdown from the primary site to the DR site during the mock session.
- **Order Handling:** All outstanding orders will be purged prior to trading from the BCP site. Members using NNF software must clear session 1 outstanding orders.
- **Multicast TBT:** Sequence numbers will restart from '1' upon switching to the BCP site.
- **Re-login Window:** Live re-login is scheduled between 18:30 hrs and 19:00 hrs on October 31, 2026.

## Regulatory Changes

None.

## Compliance Requirements

- Members must ensure their NEAT Adapter settings and NNF software systems are ready for BCP/DR switchover and non-graceful shutdown procedures in accordance with circular NSE/MSD/48662.
- Clear outstanding orders from Session-1 before initiating trading from the BCP site.

## Important Dates

- **Mock Trading Date:** October 31, 2026
- **Live Re-login Start:** 18:30 hrs (October 31, 2026)
- **Live Re-login Close:** 19:00 hrs (October 31, 2026)
- **Normal Live Operations Resume:** November 02, 2026

## Impact Assessment

Operational readiness drill for trading members with no direct financial impact on market positions.