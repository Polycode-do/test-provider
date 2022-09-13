terraform {
  required_providers {
    polycode = {
      source  = "do-2021.fr/polycode/polycode"
      version = "0.3.2"
    }
  }
}

provider "polycode" {
  host     = "http://localhost:3000"
  username = "admin@gmail.com"
  password = "12345678"
}

resource "polycode_content" "test" {
  name        = "alooooo"
  description = "testqlskdlqksjd"
  type        = "exercise"
  reward      = 1
  container {
    orientation = "vertical"
    position    = 0
    markdown {
      position = 1
      content  = "1"
    }
    markdown {
      position = 2
      content  = "3"
    }
    markdown {
      position = 3
      content  = "3"
    }
    editor {
      position = 4
      language_settings {
        default_code = "5"
        language     = "PYTHON"
        version      = ""
      }
      hint = ["value"]
      validator {
        inputs    = ["value"]
        outputs   = ["value"]
        is_hidden = false
      }
    }
    markdown {
      position = 5
      content  = "alo"
    }
    markdown {
      position = 6
      content  = "alo"
    }
    container {
      position    = 7
      orientation = "vertical"
      markdown {
        position = 1
        content  = "hello"
      }
    }
  }
}
