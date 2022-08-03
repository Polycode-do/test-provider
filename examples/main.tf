terraform {
  required_providers {
    polycode = {
      source  = "do-2021.fr/polycode/polycode"
      version = "0.1.0"
    }
  }
  required_version = ">= 1.1.0"
}

provider "polycode" {
  host     = "http://localhost:3000"
  username = "admin@gmail.com"
  password = "12345678"
}
