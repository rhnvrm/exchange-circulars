---
category: trading
circular_id: 0ecb4c202dd3f08b
date: '2026-10-30'
description: NSE introduces additional validations between NNF ID and Algo ID for
  pre-trade risk controls, effective November 23, 2026.
draft: false
feed_type: ''
guid: https://nsearchives.nseindia.com/content/circulars/CMTR77217.pdf
impact: medium
impact_ranking: medium
importance_ranking: medium
justification: Modifies pre-trade risk control validations for NNF ID and Algo ID
  matching, requiring trading members to update order routing and risk management
  systems.
pdf_url: https://nsearchives.nseindia.com/content/circulars/CMTR77217.pdf
processing:
  attempts: 1
  content_hash: 13c99be8a3ede6e9
  processed_at: '2026-10-30T17:51:44.260380'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-10-30T00:00:00+05:30'
rss_url: https://nsearchives.nseindia.com/content/circulars/CMTR77217.pdf
severity: medium
source: nse
stocks: []
tags:
- pre-trade-risk
- algo-trading
- risk-management
- trading-controls
- nse
title: 'Pre-Trade Risk Controls: Pre-Emptive Order Cancellation on Account of NNF
  ID-Algo ID Validations Update'
---

## Summary

National Stock Exchange of India Limited (NSE) has issued an update regarding pre-trade risk controls and pre-emptive order cancellations based on NNF ID and Algo ID validations, modifying item no. 3.6 C of consolidated circular NSE/CMTR/73927.

## Key Points

- Additional mandatory validations introduced between NNF ID and Algo ID in order messages.
- Specific rejection rules defined based on the 13th digit and the first 12 digits of the NNF ID.
- Orders with NNF ID first 12 digits other than "444444444444" combined with Algo ID "99999" will be rejected.
- Non-compliant orders will be rejected by the Exchange and rejection messages sent to the NNF user.
- Members should consult NNF API protocol (CM version 6.7) for details on NNF error codes and error messages.

## Regulatory Changes

Partial modification to item no. 3.6 C of exchange consolidated circular NSE/CMTR/73927 dated April 28, 2026, establishing strict cross-validation criteria between NNF IDs and Algo IDs.

## Compliance Requirements

Trading members must update their trading terminals, algorithmic trading systems, and risk management software to comply with the new NNF ID and Algo ID validation rules to prevent automated order rejections.

## Important Dates

- Circular Date: October 30, 2026
- Effective Date: November 23, 2026

## Impact Assessment

Algorithmic trading desks, institutional desks, and trading members must review their order generation logic against the newly specified NNF/Algo ID combinations to ensure seamless execution and avoid operational disruptions.