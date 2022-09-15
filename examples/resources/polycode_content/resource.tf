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
