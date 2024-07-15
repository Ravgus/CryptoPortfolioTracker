# CryptoPortfolioTracker

A small price change tracker for your crypto portfolio.

Setup Instructions
---

1. Create `portfolio.json` File.

Before using the tracker, you need to create a `portfolio.json` file in the root directory of your project. This file should contain the details of your cryptocurrency portfolio. Here is an example:

```json
{
  "coins": [
    {
      "name": "ETH",
      "count": 2
    },
    {
      "name": "BTC",
      "count": 1
    }
  ]
}
```

2. Create `.env.local` File

Create a `.env.local` file in the root directory of your project to specify the credentials for your SMTP server and the API key for the CryptoCompare service. Here is an example:

```env
CRYPTO_COMPARE_API=your_crypto_compare_api_key

SMTP_USER_NAME=john@gmail.com
SMTP_PASSWORD=your_password
SMTP_PORT=587
NOTIFICATION_CHANGE_PERCENT=25
```

P.S. 

You can omit specifying of `NOTIFICATION_CHANGE_PERCENT` variable, which defines the percentage of changing your portfolio price in order to send a notification (by default it's 25%)

3. Run the Tracker

Make sure all dependencies are installed and run the tracker:

```bash
# Install dependencies
go get ./...

# Run the tracker
go run main.go
```

How It Works
---

The CryptoPortfolioTracker fetches the latest prices of the specified cryptocurrencies from the CryptoCompare API and compares them with previous values. If significant changes are detected, it sends email notifications using the specified SMTP server.