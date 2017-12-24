# godaddy-api

Interact with GoDaddy's API.  This skips over domain purchase, and registration.  This repo is solely meant
to allow an API to create or update (upsert), or delete existing records within a domain.

The API calls wrapped in this package thing are a part of the following [API document](https://developer.godaddy.com/doc).

I created this for use by anything and anyone, but needed something more reliable for use with the [external-dns](https://github.com/kubernetes-incubator/external-dns) code.

For a simple example of API interaction using a bash script please go to the bin directory and see the `example_get.sh` script that lets you list
your records for your domain.

## Development Setup

* `Install [Golang](https://golang.org/doc/install#install).`
  -  This was done using v1.9.2 for Darwin.

* Ensure you have a GoDaddy account in some form.  Follow [this document](https://developer.godaddy.com/getstarted).

* This repo skips over domain purchase, and registration.  That is done on your own.  Be sure you have a domain purchased, managed,
  and registered under _your_ GoDaddy account.

* Generate your GoDaddy API [keys](https://developer.godaddy.com/keys/).



