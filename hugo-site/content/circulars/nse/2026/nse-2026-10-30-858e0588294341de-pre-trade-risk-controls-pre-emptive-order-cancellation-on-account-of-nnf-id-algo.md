---
category: trading
circular_id: 858e0588294341de
date: '2026-10-30'
description: NSE updates pre-trade risk control validations between NNF ID and Algo
  ID for commodity derivatives trading.
draft: false
feed_type: ''
guid: https://nsearchives.nseindia.com/content/circulars/COM77215.pdf
impact: medium
impact_ranking: medium
importance_ranking: medium
justification: Introduces additional NNF ID and Algo ID validations resulting in order
  rejections for non-compliance in commodity derivatives.
pdf_url: https://nsearchives.nseindia.com/content/circulars/COM77215.pdf
processing:
  attempts: 1
  content_hash: 1725cac5cd1d8162
  processed_at: '2026-10-30T22:02:09.625225'
  processor_version: '2.0'
  stage: completed
  status: published
published_date: '2026-10-30T00:00:00+05:30'
rss_url: https://nsearchives.nseindia.com/content/circulars/COM77215.pdf
severity: medium
source: nse
stocks: []
tags:
- risk-management
- algo-trading
- trading-controls
- commodity-derivatives
- nse
title: 'Pre-Trade Risk Controls: Pre-Emptive Order Cancellation on Account of NNF
  ID-Algo ID Validations - Update'
---

## Summary

National Stock Exchange of India Limited (NSE) has issued an update regarding pre-trade risk controls concerning pre-emptive order cancellations on account of NNF ID and Algo ID validations in the commodity derivatives segment, modifying consolidated circular NSE/COM/73930 dated April 28, 2026.

## Key Points

- Additional validations have been introduced between NNF ID and Algo ID.
- Orders will be rejected and error messages sent to NNF users if specific mismatch criteria involving 12th/13th digits of NNF ID and Algo ID are met.
- Effective date for these validations is November 23, 2026.

## Regulatory Changes

- If the 13th digit in the NNF ID field is "0, 2, 4, or 5" and the first 12 digits are "111111111111" or "333333333333", the order will be rejected.
- If the 13th digit in the NNF ID field is "1, 3, 5, 6, 7, or 8" and the first 12 digits are "444444444444", the order will be rejected.
- If the first 12 digits of the NNF ID have values other than "444444444444" and the Algo ID field contains "99999", the order will be rejected.

## Compliance Requirements

- Members must ensure their trading systems and order messages comply with the updated NNF ID and Algo ID validation rules.
- Refer to the NNF API protocol (COM version 3.6) for details on NNF error codes and messages.

## Important Dates

- Circular Date: October 30, 2026
- Effective Date: November 23, 2026

## Impact Assessment

Trading members must update their algorithmic trading systems and order routing logic to prevent order rejections and ensure seamless trade execution in the commodity derivatives segment starting November 23, 2026.