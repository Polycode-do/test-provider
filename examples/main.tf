terraform {
  required_providers {
    polycode = {
      source  = "do-2021.fr/polycode/polycode"
      version = "v0.3.5-rc4"
    }
  }
}

provider "polycode" {
  host     = "http://localhost:3000"
  username = "admin@gmail.com"
  password = "12345678"
}

resource "polycode_item" "test_item" {
  cost = 10
  hint {
    text = "This is a hint"
  }
}

resource "polycode_content" "test_content" {
  name        = "Test content"
  description = "This is a test content"
  reward      = 100
  type        = "exercise"

  container {
    orientation = "vertical"
    position    = 0

    markdown {
      position = 1
      content  = "# This is a markdown component"
    }

    container {
      position    = 2
      orientation = "vertical"

      markdown {
        position = 1
        content  = "# This is a nested markdown component"
      }

      editor {
        position = 2
        hint     = [polycode_item.test_item.id]

        language_settings {
          default_code = "print('Hello world')"
          language     = "PYTHON"
        }

        validator {
          inputs  = []
          outputs = ["Hello world"]
        }
      }
    }
  }
}
