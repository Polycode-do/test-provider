terraform {
  required_providers {
    polycode = {
      source  = "do-2021.fr/polycode/polycode"
      version = "0.2.0"
    }
  }
  required_version = ">= 1.1.0"
}

provider "polycode" {
  host     = "http://localhost:3000"
  username = "admin@gmail.com"
  password = "12345678"
}

data "polycode_user" "guest_user" {
  id = "2f59a4db-f922-4e1e-adf6-e7f3cb4b91f1"
}

output "output_user" {
  value = data.polycode_user.guest_user
}
