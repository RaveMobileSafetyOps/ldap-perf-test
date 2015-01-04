# LDAP Performance Tester

[![Build Status](https://secure.travis-ci.org/jbuchbinder/ldap-perf-test.png)](http://travis-ci.org/jbuchbinder/ldap-perf-test)

[![Gobuild Download](http://gobuild.io/badge/github.com/jbuchbinder/ldap-perf-test/downloads.svg)](http://gobuild.io/github.com/jbuchbinder/ldap-perf-test)

## Building

```
go get github.com/jbuchbinder/goldap
go build
```

## Usage

```
Usage of ./ldap-perf-test:
  -delay=5000: Delay randomization
  -laxssl=false: Use lax SSL restrictions
  -ldapbase="": LDAP base DN
  -ldapbind="": LDAP bind DN
  -ldaphost="127.0.0.1": LDAP host
  -ldappass="": LDAP password
  -ldapport=389: LDAP port
  -ldapsearch="": LDAP search string
  -ldapssl=false: LDAP SSL enabled
  -threads=1: Test threads
```

