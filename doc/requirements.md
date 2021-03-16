## Design choices

* Only Kubernetes-Based
  * Just too hard to decouple concepts, k8s is platform independent
  * Direct k8s or using a cloud
* Initial support for
  * Gcloud
  * CTFScoreboard

## Features/TODO

- [ ] Provision challenges in Kubernetes
  - [ ] Deploy
  - [ ] Update
  - [ ] Status
- [ ] Support Challenge Types
  - [ ] Offline
  - [ ] Web (via HTTP/HTTPS LB)
  - [ ] TCP (via service)
- [ ] Challenge edge cases
  - [ ] Multiple flags/challenge
- [ ] Multiple cloud providers
  - [ ] GCloud (via cli??)
  - [ ] Kubernetes (via api)
- [ ] Build Docker Images
- [ ] Support Infrastructure Images/Containers (non-challenge)
    - [ ] Deploy non-challenge containers
    - [ ] Types
      - [ ] Sidecar
      - [ ] Container (i.e., deployment only)
      - [ ] Service (i.e., deployment + service)
- [ ] Deploy sidecars
  - [ ] Webbot
  - [ ] CloudSQL Proxy
- [ ] Provision Underlying Resources
  - [ ] IP Allocation
  - [ ] k8s cluster setup
  - [ ] DNS Entries
  - [ ] Network Policies
  - [ ] k8s ingress
- [ ] Local Run
  - [ ] Support running challenges locally via podman/docker
