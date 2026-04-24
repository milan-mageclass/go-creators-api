# go-creators-api

A small Go client for the Amazon Creators API.

Amazon is moving everyone from PAAPI to Creators API, and this package is here to make that transition a bit less annoying.

This project is inspired by [`utekaravinash/gopaapi5`](https://github.com/utekaravinash/gopaapi5). The idea was to keep the Go side familiar and easy to work with, while adapting it to the Creators API auth flow and payload format.

Amazon's official migration guide is here:

https://affiliate-program.amazon.com/creatorsapi/docs/en-us/migrating-to-creatorsapi-from-paapi

## What it does right now

- `GetItems`
- `GetVariations`
- in-memory bearer token caching
- automatic token refresh before expiry
- typed resource constants
- typed response models

## What is different from PAAPI

There are two main differences that matter when moving from PAAPI:

1. Authentication

PAAPI signs every request with AWS Signature V4.

Creators API uses a bearer token instead. This library fetches it, keeps it in memory, and refreshes it automatically when needed.

2. Request and response naming

PAAPI mostly uses PascalCase.

Creators API uses lower camel case for request fields and resource names. For example:

- PAAPI: `ItemIds`, `PartnerTag`
- Creators API: `itemIds`, `partnerTag`

## Install

```bash
go get github.com/milan-mageclass/go-creators-api
```

## Basic setup

If you just want to get moving, the examples are the best place to start:

- [`_examples/get_items`](_examples/get_items)
- [`_examples/get_variations`](_examples/get_variations)

## Credentials

You need:

- `CredentialID`
- `CredentialSecret`
- `CredentialVersion`
- `PartnerTag`
- `Marketplace`

Typical environment variables:

```bash
export CREATORS_API_CLIENT_ID="..."
export CREATORS_API_CLIENT_SECRET="..."
export CREATORS_API_CLIENT_VERSION="2.1"
export CREATORS_API_PARTNER_TAG="yourtag-20"
```

## Notes about Offers

Creators API does not use the old `Offers` shape in the same way PAAPI did.

This package currently models `OffersV2`, which is what you should use.

`RentalOffers` are not supported here because they are not part of the Creators API flow this package targets.

## Status

This package is still pretty young, and some response models will likely keep getting adjusted as more real-world payloads show up.

That said, it is already usable today for `GetItems` and `GetVariations`, and the goal is to keep it small, predictable, and pleasant to use while the rest of the API support is added.
