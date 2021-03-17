## Initial Setup

1. Setup cloud provider.
2. Configuration in config.yml
3. (Optional) Use ctftk to create cluster
  1. Use cloud API to provision
4. Apply cluster config
  1. Network policy
6. Get CA Certificates
  1. Requires DNS Support
  2. ACME wildcard for LB
5. Create HTTP(s) LoadBalancer
  1. Allocate IP
  2. Provision LB

## Challenge Deployment

1. Build Container(s)
2. Push container to registry
3. Launch deployment
4. Launch service
5. Push to scoreboard

## Challenge Update

Options:
  - Force rebuild
  - Only scoreboard
  - Only k8s

1. Build Container(s)
2. Push container to registry
3. Push update to deployment
4. Push update to service
