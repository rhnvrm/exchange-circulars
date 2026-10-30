---
category: market-operations
circular_id: 6e811e2fd71ec66f
date: '2026-10-30'
description: NSE updates revised dates for closing price broadcasts via trading APIs
  and discontinuation timelines for Extranet and NNF API Bhavcopy reports.
draft: false
feed_type: ''
guid: https://nsearchives.nseindia.com/content/circulars/MSD77208.pdf
impact: medium
impact_ranking: medium
importance_ranking: medium
justification: Provides revised timelines and schedule updates for EOD Bhavcopy dissemination
  across trading APIs, Extranet, and NNF API broadcasts affecting trading members.
pdf_url: https://nsearchives.nseindia.com/content/circulars/MSD77208.pdf
processing:
  attempts: 1
  content_hash: 06e7dd7d4868abe8
  processed_at: '2026-10-30T22:02:16.066809'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-10-30T00:00:00+05:30'
rss_url: https://nsearchives.nseindia.com/content/circulars/MSD77208.pdf
severity: medium
source: nse
stocks: []
tags:
- bhavcopy
- market-operations
- trading
- extranet
- segment-bhavcopy
title: Streamlining EOD information dissemination related to Bhavcopy - Update
---

## Summary

National Stock Exchange of India Limited (NSE) has issued an update regarding the streamlining of End of Day (EOD) information dissemination related to Bhavcopy. The circular specifies revised dates for closing price broadcasts through trading APIs for COM and FO segments, as well as revised discontinuation dates for Bhavcopy reports on Extranet and NNF API broadcasts.

## Key Points

- **Trading API Closing Price Broadcasts:** 
  - COM (Segment): Revised date is November 23, 2026 (transcode: BCAST_ONLY_MBP / 7208).
  - FO (Segment): Revised date is November 23, 2026 (transcode: BCAST_ENHNCD_ONLY_MBP / 17208).
  - CD (Segment): No change (November 16, 2026; transcode: BCAST_ONLY_MBP / 7208).
  - CM and EGR: Currently disseminated through BCAST_ONLY_MBP (7208).
- **Extranet Bhavcopy Discontinuation:** The discontinuation date for Bhavcopy reports on Extranet is revised from December 07, 2026 to January 09, 2027.
- **NNF API Broadcast Discontinuation:** Mode 2 Bhavcopy (.MS & .MD) disseminated via NNF API broadcast for CM, FO, CD, COM, EG, and Index Bhavcopy are subject to updated discontinuation schedules.

## Regulatory Changes

Adjustment of timelines and discontinuation schedules for legacy Bhavcopy dissemination modes (Extranet and NNF API broadcasts) and transition to streamlined API broadcast mechanisms across commodity, derivatives, currency, and capital market segments.

## Compliance Requirements

Trading members must review their data ingestion pipelines, back-office systems, and automated downloading routines to align with the revised broadcast transaction codes and updated discontinuation dates for Extranet and NNF API Bhavcopy files.

## Important Dates

- **November 23, 2026:** Revised effective date for closing price broadcast in COM and FO segments via trading APIs.
- **January 09, 2027:** Revised discontinuation date for Bhavcopy reports on Extranet.

## Impact Assessment

Trading members and data vendors relying on NSE EOD Bhavcopy reports must update their automated ingestion systems to accommodate the revised broadcast transaction codes and prepare for the eventual deprecation of Extranet and NNF API Bhavcopy delivery modes.