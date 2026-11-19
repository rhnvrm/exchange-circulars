---
category: market-operations
circular_id: 6b9b90da948a4e17
date: '2026-11-19'
description: NSE announces a mock trading session for the Capital Market segment on
  Saturday, November 21, 2026, along with the release of NEAT+ version 7.9.1 and version
  migration timelines.
draft: false
feed_type: ''
guid: https://nsearchives.nseindia.com/content/circulars/CMTR77516.pdf
impact: medium
impact_ranking: medium
importance_ranking: medium
justification: Mandatory software upgrade to NEAT+ 7.9.1, mock trading participation,
  and NNF encryption port migrations for trading members.
pdf_url: https://nsearchives.nseindia.com/content/circulars/CMTR77516.pdf
processing:
  attempts: 1
  content_hash: 639359cd86ca77b1
  processed_at: '2026-11-19T11:12:26.717214'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-11-19T00:00:00+05:30'
rss_url: https://nsearchives.nseindia.com/content/circulars/CMTR77516.pdf
severity: medium
source: nse
stocks: []
tags:
- mock-trading
- neat-plus
- capital-market
- software-upgrade
- trading-operations
- nnf
title: Mock trading on Saturday, November 21, 2026 - New version of NEAT+ 7.9.1
---

## Summary

National Stock Exchange of India Limited has announced a mock trading session in the Capital Market segment on Saturday, November 21, 2026, utilizing the newly released NEAT+ version 7.9.1. The exchange will discontinue the existing NEATPlus version 7.9.0 effective from the December 05, 2026 mock session. Members are also advised regarding pre-trade risk controls and upcoming NNF encryption mechanism changes.

## Key Points

- Mock trading session scheduled for Saturday, November 21, 2026, across primary and DR sites.
- New NEAT+ version 7.9.1 available for download from November 20, 2026 (19:00 hrs) via extranet path `/cmftp/common/NEATPlus791`.
- Existing version 7.9.0 discontinuation scheduled for December 05, 2026 mock session.
- Members requested to test pre-trade risk controls (pre-emptive order cancellation on account of NNF ID-Algo ID validations) during the mock session.
- Discontinuation of the existing encryption mechanism with old interactive port (10827) effective December 07, 2026, with migration advised to new encryption mechanisms.

## Regulatory Changes

- Pre-trade risk controls regarding pre-emptive order cancellation (Circular NSE/CM/77217) become effective in LIVE from November 23, 2026.
- Discontinuation of the old interactive port encryption mechanism (port 10827) starting December 07, 2026, requiring migration to updated encryption mechanisms with MD5 checksum removal or immediate order acknowledgment.

## Compliance Requirements

- Trading members must participate or test systems during the mock trading session on November 21, 2026.
- Upgrade trading terminals to NEAT+ version 7.9.1 prior to the discontinuation of version 7.9.0 on December 05, 2026.
- Share feedback and confirmation on pre-trade risk controls testing to `msm@nse.co.in`.
- Migrate gateway router ports and encryption configurations as per specified subnet ranges and new ports (e.g., port 10831).

## Important Dates

- **November 20, 2026 (19:00 hrs):** Availability of NEAT+ version 7.9.1
- **November 21, 2026:** Mock Trading Session
- **November 23, 2026:** Pre-trade risk controls effective in LIVE
- **December 05, 2026:** Discontinuation date for NEAT+ version 7.9.0 (during mock session)
- **December 07, 2026:** Discontinuation of old interactive port encryption mechanism

## Impact Assessment

Operational impact is medium, requiring all trading members to execute software upgrades, test pre-trade risk controls, and reconfigure network encryption ports to ensure uninterrupted trading connectivity.