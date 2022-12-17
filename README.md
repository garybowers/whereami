# Where Am I?

Where Am I? is a demonstration application written in Golang desiged to demonstrate the multi-cluster setups of GKE & Anthos.  

Where Am I? simply reads the GCP metadata and returns the current cluster name that the user is accessing, it prints this out on a http server on port 8080.

The main purpose is when you have Multi Cluster Ingress or a Global Loadbalancer in front to demonstrate the network path taken.

**Note:** This only works with GKE running on Google Cloud Platform - I don't use or like other cloud platforms so have no interest in porting it over.
