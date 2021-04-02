# crypto_stats

## Dependency manager

GOMOD

## HOW TO BUILD AND RUN

    make build_dependency run

# CONFIGURABLE PARMAS

    ./bin/crypto "<comma separated symbols>"

# DEFAULT PARAMS

    'BTCUSD,ETHBTC'

# ENDPOINTS

    `http://localhost:8080/currency/<SUPPORTED_SYMBOL>`

NOTE: to get all the symbol details, replace `SUPPORTED_SYMBOL` with `all`

eg

## REQ:-> http://localhost:8080/currency/all

---

RESP

```json
[
  {
    "id": "BTCUSD",
    "fullName": "",
    "ask": "5127.99",
    "bid": "5123.77",
    "last": "5124.33",
    "open": "5262.15",
    "low": "5015.01",
    "high": "5536.48",
    "feeCurrency": "USD"
  },
  {
    "id": "ETHBTC",
    "fullName": "",
    "ask": "0.021892",
    "bid": "0.021889",
    "last": "0.021892",
    "open": "0.022054",
    "low": "0.021546",
    "high": "0.022260",
    "feeCurrency": "BTC"
  }
]
```

## REQ:-> http://localhost:8080/currency/ETHBTC

RESP

```json
{
  "id": "BTCUSD",
  "fullName": "",
  "ask": "5101.33",
  "bid": "",
  "last": "5100.88",
  "open": "5230.25",
  "low": "5015.01",
  "high": "5536.48",
  "feeCurrency": "USD"
}
```

## REQ:-> http://localhost:8080/currency/ABCD

RESP

```json
"Symbol ABCD not supported"
```
