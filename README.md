# hr-pharma-hlf

Solution repository for the **Pharma Cold-Chain Custody** Hyperledger Fabric
chaincode challenge (NPCI / HackerRank, Hard).

Standard Fabric **test-network** plus a chaincode skeleton at
[`chaincode/pharma.go`](chaincode/pharma.go). Cloned into the candidate's
environment by the HackerRank Setup Script (via [`setup.sh`](setup.sh)).

## Candidate task
1. Implement the functions in `chaincode/pharma.go`, including a one-way breach
   flag and full custody provenance via GetHistoryForKey.
2. Deploy: `cd test-network && ./network.sh deployCC -ccn pharmacc -ccp ../chaincode -ccl go`
3. Create shp1 (Factory), transfer to Distributor then Pharmacy, record a breach, deliver.

---

Authored by **Dayal Mukati** — [dayalmukati.com](https://dayalmukati.com)
