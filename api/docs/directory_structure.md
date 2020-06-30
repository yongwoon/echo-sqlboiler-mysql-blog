# Directory structure

- api
  - OpenAPI/Swagger specs, JSON schema files, protocol definition files.
- config
  - Configuration file templates or default configs.
- config/initializer
  - System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.
- controller
- db
  - migrations, fixtures
- deployments
  - IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh). Note that in some repos (especially apps deployed with kubernetes) this directory is called
- lib
- middleware
  - middlewares
  - ex. auth
- model
- route
  - Define routing
- script
  - Scripts to perform various build, install, analysis, etc operations.
  - ex. kubernetes helm, terraform
- test
  - Additional external test apps and test data. Feel free to structure the /test directory anyway you want.